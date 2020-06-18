package test

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
)

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
	  import FungibleToken from 0x%[1]s 
	  import %[3]s from 0x%[2]s

      transaction {

          prepare(acct: AuthAccount) {
              let vault <- %[3]s.createEmptyVault()
              acct.save(<-vault, to: /storage/%[4]sVault)

              acct.link<&{FungibleToken.Receiver}>(/public/%[4]sReceiver, target: /storage/%[4]sVault)
              acct.link<&%[3]s.Vault{FungibleToken.Balance}>(/public/%[4]sBalance, target: /storage/%[4]sVault)
          }
      }
    `
	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName))
}

// GenerateDestroyVaultScript creates a script that withdraws
// tokens from a vault and destroys the tokens
func GenerateDestroyVaultScript(fungibleAddr, tokenAddr flow.Address, tokenName string, withdrawAmount int) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s

		transaction {
		  prepare(acct: AuthAccount) {
			let vault <- acct.load<@%[3]s.Vault>(from: /storage/%[4]sVault)
				?? panic("Couldn't load Vault from storage")
			
			let withdrawVault <- vault.withdraw(amount: %[5]d.0)

			acct.save(<-vault, to: /storage/%[4]sVault) 

			destroy withdrawVault
		  }
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName, withdrawAmount))
}

// GenerateTransferVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault
func GenerateTransferVaultScript(fungibleAddr, tokenAddr flow.Address, receiverAddr flow.Address, tokenName string, amount int) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%s 
		import %s from 0x%s

		transaction {
			prepare(acct: AuthAccount) {
				let recipient = getAccount(0x%s)

				let providerRef = acct.borrow<&{FungibleToken.Provider}>(from: /storage/%sVault)
					?? panic("Could not borrow Provider reference to the Vault!")

				let receiverRef = recipient.getCapability(/public/%sReceiver)!.borrow<&{FungibleToken.Receiver}>()
					?? panic("Could not borrow receiver reference to the recipient's Vault")

				let tokens <- providerRef.withdraw(amount: %d.0)

				receiverRef.deposit(from: <-tokens)
			}
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenName, tokenAddr, receiverAddr, storageName, storageName, amount))
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(fungibleAddr, tokenAddr flow.Address, receiverAddr flow.Address, tokenName string, amount float64) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s
	
		transaction {
			let tokenAdmin: &%[3]s.Administrator
			let tokenReceiver: &{FungibleToken.Receiver}
	
			prepare(signer: AuthAccount) {
			  self.tokenAdmin = signer
				.borrow<&%[3]s.Administrator>(from: /storage/%[4]sAdmin) 
				?? panic("Signer is not the token admin")
	
			  self.tokenReceiver = getAccount(0x%[5]s)
				.getCapability(/public/%[4]sReceiver)!
				.borrow<&{FungibleToken.Receiver}>()
				?? panic("Unable to borrow receiver reference")
			}
	
			execute {
			  let minter <- self.tokenAdmin.createNewMinter(allowedAmount: 100.0)
			  let mintedVault <- minter.mintTokens(amount: %[6]f)
	
			  self.tokenReceiver.deposit(from: <-mintedVault)
	
			  destroy minter
			}
		  }
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName, receiverAddr, amount))
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(fungibleAddr, tokenAddr flow.Address, tokenName string, amount int) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
	import FungibleToken from 0x%[1]s 
	import %[3]s from 0x%[2]s
	
	transaction {
	
		// Vault resource that holds the tokens that are being burned
		let vault: @FungibleToken.Vault
	
		let admin: &%[3]s.Administrator
	
		prepare(signer: AuthAccount) {
	
			// Withdraw tokens from the admin vault in storage
			self.vault <- signer.borrow<&%[3]s.Vault>(from: /storage/%[4]sVault)!
				.withdraw(amount: UFix64(%[5]d))
	
			// Create a reference to the admin admin resource in storage
			self.admin = signer.borrow<&%[3]s.Administrator>(from: /storage/%[4]sAdmin)
				?? panic("Could not borrow a reference to the admin resource")
		}
	
		execute {
			let burner <- self.admin.createNewBurner()
			
			burner.burnTokens(from: <-self.vault)

            destroy burner
		}
	}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName, amount))
}

// GenerateTransferInvalidVaultScript creates a script that withdraws an tokens from an account
// and tries to deposit it into a vault of the wrong type. Should fail
func GenerateTransferInvalidVaultScript(fungibleAddr, tokenAddr, otherTokenAddr, receiverAddr flow.Address, tokenName, otherTokenName string, amount int) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	otherStorageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%s 
		import %s from 0x%s
		import %s from 0x%s

		transaction {
			prepare(acct: AuthAccount) {
				let recipient = getAccount(0x%s)

				let providerRef = acct.borrow<&{FungibleToken.Provider}>(from: /storage/%sVault)
					?? panic("Could not borrow Provider reference to the Vault!")

				let receiverRef = recipient.getCapability(/public/%sReceiver)!.borrow<&{FungibleToken.Receiver}>()
					?? panic("Could not borrow receiver reference to the recipient's Vault")

				let tokens <- providerRef.withdraw(amount: %d.0)

				receiverRef.deposit(from: <-tokens)
			}
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenName, tokenAddr, otherTokenName, otherTokenAddr, receiverAddr, storageName, otherStorageName, amount))
}

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(fungibleAddr, tokenAddr, userAddr flow.Address, tokenName string, expectedBalance float64) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s

		pub fun main() {
			let acct = getAccount(0x%[5]s)
			let vaultRef = acct.getCapability(/public/%[4]sBalance)!.borrow<&%[3]s.Vault{FungibleToken.Balance}>()
				?? panic("Could not borrow Balance reference to the Vault")
			assert(
                vaultRef.balance == UFix64(%[6]f),
                message: "incorrect balance!"
            )
		}
    `

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName, userAddr, expectedBalance))
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, tokenAddr flow.Address, tokenName string, expectedSupply int) []byte {

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s

		pub fun main() {
			assert(
                %[3]s.totalSupply == UFix64(%[4]d),
                message: "incorrect totalSupply!"
            )
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, expectedSupply))
}

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreateForwarderScript(fungibleAddr, forwardingAddr, receiverAddr flow.Address, tokenName string) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
	  	import FungibleToken from 0x%[1]s 
	  	import TokenForwarding from 0x%[2]s

      	transaction {

        	prepare(acct: AuthAccount) {
				let recipient = getAccount(0x%[4]s).getCapability(/public/%[3]sReceiver)!

		        let vault <- TokenForwarding.createNewForwarder(recipient: recipient)
              	acct.save(<-vault, to: /storage/%[3]sForwarder)

				if acct.getCapability(/public/%[3]sReceiver)!.borrow<&{FungibleToken.Receiver}>() != nil {
					acct.unlink(/public/%[3]sReceiver)
				}
				acct.link<&{FungibleToken.Receiver}>(/public/%[3]sReceiver, target: /storage/%[3]sForwarder)
          	}
      	}
    `
	return []byte(fmt.Sprintf(template, fungibleAddr, forwardingAddr, storageName, receiverAddr))
}
