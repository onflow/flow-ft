package templates

import (
	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
	"github.com/onflow/flow-go-sdk"
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
func GenerateInspectVaultScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(scriptsPath + readBalanceFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(scriptsPath + readSupplyFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateInspectSupplyViewScript creates a script that reads
// the total supply of tokens in existence through a metadata view
func GenerateInspectSupplyViewScript(fungibleAddr, tokenAddr, metadataViewsAddr, ftMetadataViewsAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(scriptsPath + readSupplyViewFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, metadataViewsAddr, ftMetadataViewsAddr, tokenName)
}
