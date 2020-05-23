package fttest

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
)

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(fungibleAddr, flowAddr flow.Address) []byte {
	template := `
	  import FungibleToken from 0x%s 
	  import FlowToken from 0x%s

      transaction {

          prepare(acct: AuthAccount) {
              let vault <- FlowToken.createEmptyVault()
              acct.save(<-vault, to: /storage/flowTokenVault)

              acct.link<&FlowToken.Vault{FungibleToken.Receiver}>(/public/flowTokenReceiver, target: /storage/flowTokenVault)
              acct.link<&FlowToken.Vault{FungibleToken.Balance}>(/public/flowTokenBalance, target: /storage/flowTokenVault)
          }
      }
    `
	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr))
}

// GenerateDestroyVaultScript creates a script that withdraws
// tokens from a vault and destroys the tokens
func GenerateDestroyVaultScript(fungibleAddr, flowAddr flow.Address, withdrawAmount int) []byte {
	template := `
		import FungibleToken from 0x%s 
		import FlowToken from 0x%s

		transaction {
		  prepare(acct: AuthAccount) {
			let vault <- acct.load<@FlowToken.Vault>(from: /storage/flowTokenVault)
				?? panic("Couldn't load Vault from storage")
			
			let withdrawVault <- vault.withdraw(amount: %d.0)

			acct.save(<-vault, to: /storage/flowTokenVault) 

			destroy withdrawVault
		  }
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, withdrawAmount))
}

// GenerateTransferVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault
func GenerateTransferVaultScript(fungibleAddr, flowAddr flow.Address, receiverAddr flow.Address, amount int) []byte {
	template := `
		import FungibleToken from 0x%s 
		import FlowToken from 0x%s

		transaction {
			prepare(acct: AuthAccount) {
				let recipient = getAccount(0x%s)

				let providerRef = acct.borrow<&FlowToken.Vault{FungibleToken.Provider}>(from: /storage/flowTokenVault)
					?? panic("Could not borrow Provider reference to the Vault!")

				let receiverRef = recipient.getCapability(/public/flowTokenReceiver)!.borrow<&FlowToken.Vault{FungibleToken.Receiver}>()
					?? panic("Could not borrow receiver reference to the recipient's Vault")

				let tokens <- providerRef.withdraw(amount: %d.0)

				receiverRef.deposit(from: <-tokens)
			}
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, receiverAddr, amount))
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(fungibleAddr, flowAddr flow.Address, receiverAddr flow.Address, amount int) []byte {
	template := `
		import FungibleToken from 0x%s 
		import FlowToken from 0x%s
	
		transaction {
	
			// Vault resource that holds the tokens that are being minted
			var vault: @FungibleToken.Vault
		
			prepare(signer: AuthAccount) {
		
				// Get a reference to the signer's MintAndBurn resource in storage
				let mintAndBurn = signer.borrow<&FlowToken.MintAndBurn>(from: /storage/flowTokenMintAndBurn)
					?? panic("Couldn't borrow MintAndBurn reference from storage")
		
				// Mint 10 new tokens
				self.vault <- mintAndBurn.mintTokens(amount: %d.0)
			}
		
			execute {
				// Get the recipient's public account object
				let recipient = getAccount(0x%s)
		
				// Get a reference to the recipient's Receiver
				let receiver = recipient.getCapability(/public/flowTokenReceiver)!
					.borrow<&FlowToken.Vault{FungibleToken.Receiver}>()
					?? panic("Couldn't borrow receiver reference to recipient's vault")
		
				// Deposit the newly minted token in the recipient's Receiver
				receiver.deposit(from: <-self.vault)
			}
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, amount, receiverAddr))
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(fungibleAddr, flowAddr flow.Address, amount int) []byte {
	template := `
	import FungibleToken from 0x%s 
	import FlowToken from 0x%s
	
	transaction {
	
		// Vault resource that holds the tokens that are being burned
		let vault: @FungibleToken.Vault
	
		let mintAndBurn: &FlowToken.MintAndBurn
	
		prepare(signer: AuthAccount) {
	
			// Withdraw 10 tokens from the admin vault in storage
			self.vault <- signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)!
				.withdraw(amount: UFix64(%d))
	
			// Create a reference to the admin MintAndBurn resource in storage
			self.mintAndBurn = signer.borrow<&FlowToken.MintAndBurn>(from: /storage/flowTokenMintAndBurn)
				?? panic("Could not borrow a reference to the Burn resource")
		}
	
		execute {
			// burn the withdrawn tokens
			self.mintAndBurn.burnTokens(from: <-self.vault)
		}
	}
	
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, amount))
}

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(fungibleAddr, flowAddr, userAddr flow.Address, expectedBalance float64) []byte {
	template := `
		import FungibleToken from 0x%s 
		import FlowToken from 0x%s

		pub fun main() {
			let acct = getAccount(0x%s)
			let vaultRef = acct.getCapability(/public/flowTokenBalance)!.borrow<&FlowToken.Vault{FungibleToken.Balance}>()
				?? panic("Could not borrow Balance reference to the Vault")
			assert(
                vaultRef.balance == UFix64(%f),
                message: "incorrect balance!"
            )
		}
    `

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, userAddr, expectedBalance))
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, flowAddr flow.Address, expectedSupply int) []byte {
	template := `
		import FungibleToken from 0x%s 
		import FlowToken from 0x%s

		pub fun main() {
			assert(
                FlowToken.totalSupply == UFix64(%d),
                message: "incorrect totalSupply!"
            )
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, flowAddr, expectedSupply))
}
