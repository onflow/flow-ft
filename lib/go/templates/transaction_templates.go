package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	defaultTokenName         = "ExampleToken"
	defaultTokenStorage      = "exampleToken"
	defaultFungibleTokenAddr = "FUNGIBLETOKENADDRESS"
	defaultTokenAddr         = "TOKENADDRESS"
	defaultForwardingAddr    = "FORWARDINGADDRESS"

	transferTokensFilename  = "transfer_tokens.cdc"
	setupAccountFilename    = "setup_account.cdc"
	mintTokensFilename      = "mint_tokens.cdc"
	createForwarderFilename = "create_forwarder.cdc"
	burnTokensFilename      = "burn_tokens.cdc"
)

func replaceAddresses(code string, fungibleAddr, tokenAddr, tokenName string) string {
	storageName := MakeFirstLowerCase(tokenName)

	replacer := strings.NewReplacer("0x"+defaultFungibleTokenAddr, "0x"+fungibleAddr,
		"0x"+defaultTokenAddr, "0x"+tokenAddr,
		defaultTokenName, tokenName,
		defaultTokenStorage, storageName)

	code = replacer.Replace(code)

	return code
}

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(setupAccountFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	return []byte(code)
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

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	return []byte(code)
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(mintTokensFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	return []byte(code)
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(burnTokensFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	return []byte(code)
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

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultForwardingAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

// MakeFirstLowerCase makes the first letter in a string lowercase
func MakeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
