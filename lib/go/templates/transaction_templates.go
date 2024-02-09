package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	_ "github.com/kevinburke/go-bindata"

	"strings"

	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	transferTokensFilename       = "transfer_tokens.cdc"
	genericTransferFilename      = "generic_transfer.cdc"
	transferManyAccountsFilename = "transfer_many_accounts.cdc"
	setupAccountFilename         = "setup_account.cdc"
	setupGenericVaultFilename    = "metadata/setup_account_from_address.cdc"
	mintTokensFilename           = "mint_tokens.cdc"
	createForwarderFilename      = "create_forwarder.cdc"
	burnTokensFilename           = "burn_tokens.cdc"
)

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(env Environment) []byte {

	code := assets.MustAssetString(setupAccountFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateSetupGenericVaultScript creates a script that instantiates
// a new Vault instance and stores it in storage. It can be used
// to create any token as long as you know the address of the contract
// and the name of the contract
func GenerateSetupAccountFromAddressScript(fungibleTokenAddr, fungibleTokenMVAddr string) []byte {

	code := assets.MustAssetString(setupGenericVaultFilename)

	code = strings.ReplaceAll(
		code,
		placeholderFungibleToken,
		withHexPrefix(fungibleTokenAddr),
	)

	code = strings.ReplaceAll(
		code,
		placeholderFTMetadataViews,
		withHexPrefix(fungibleTokenMVAddr),
	)

	return []byte(code)
}

// GenerateTransferVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault
func GenerateTransferVaultScript(env Environment) []byte {

	code := assets.MustAssetString(transferTokensFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateTransferGenericVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault for any vault type
func GenerateTransferGenericVaultScript(fungibleTokenAddr string) []byte {

	code := assets.MustAssetString(genericTransferFilename)

	code = strings.ReplaceAll(
		code,
		placeholderFungibleToken,
		withHexPrefix(fungibleTokenAddr),
	)

	return []byte(code)
}

// GenerateTransferManyAccountsScript creates a script that transfers the same number of tokens
// to a list of accounts
func GenerateTransferManyAccountsScript(env Environment) []byte {

	code := assets.MustAssetString(transferManyAccountsFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(env Environment) []byte {

	code := assets.MustAssetString(mintTokensFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(env Environment) []byte {
	code := assets.MustAssetString(burnTokensFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreateForwarderScript(env Environment) []byte {
	code := assets.MustAssetString(createForwarderFilename)

	return []byte(ReplaceAddresses(code, env))
}
