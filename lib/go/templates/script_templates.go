package templates

import (
	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	scriptsPath            = "scripts/"
	readBalanceFilename    = "get_balance.cdc"
	readSupplyFilename     = "get_supply.cdc"
	readSupplyViewFilename = "metadata/get_vault_supply_view.cdc"
)

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readBalanceFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(env Environment) []byte {

	code := assets.MustAssetString(scriptsPath + readSupplyFilename)

	return []byte(ReplaceAddresses(code, env))
}

// GenerateInspectSupplyViewScript creates a script that reads
// the total supply of tokens in existence through a metadata view
func GenerateInspectSupplyViewScript(env Environment) []byte {

	code := assets.MustAssetString(scriptsPath + readSupplyViewFilename)

	return []byte(ReplaceAddresses(code, env))
}
