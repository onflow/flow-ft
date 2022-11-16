package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	transferTokensFilename       = "transfer_tokens.cdc"
	genericTransferFilename      = "generic_transfer.cdc"
	transferManyAccountsFilename = "transfer_many_accounts.cdc"
	setupAccountFilename         = "setup_account.cdc"
	mintTokensFilename           = "mint_tokens.cdc"
	createForwarderFilename      = "create_forwarder.cdc"
	burnTokensFilename           = "burn_tokens.cdc"
)

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(fungibleAddr, tokenAddr, metadataViewsAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(setupAccountFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, metadataViewsAddr, tokenName)
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
func GenerateTransferVaultScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(transferTokensFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateTransferGenericVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault for any vault type
func GenerateTransferGenericVaultScript(fungibleAddr flow.Address) []byte {

	code := assets.MustAssetString(genericTransferFilename)

	return replaceAddresses(code, fungibleAddr, flow.EmptyAddress, flow.EmptyAddress, flow.EmptyAddress, "")
}

// GenerateTransferManyAccountsScript creates a script that transfers the same number of tokens
// to a list of accounts
func GenerateTransferManyAccountsScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(transferManyAccountsFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(mintTokensFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(burnTokensFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateTransferInvalidVaultScript creates a script that withdraws an tokens from an account
// and tries to deposit it into a vault of the wrong type. Should fail
func GenerateTransferInvalidVaultScript(fungibleAddr, tokenAddr, otherTokenAddr, receiverAddr flow.Address, tokenName, otherTokenName string, amount int) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	otherStorageName := MakeFirstLowerCase(otherTokenName)

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

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createForwarderFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, forwardingAddr, flow.EmptyAddress, tokenName)
}
