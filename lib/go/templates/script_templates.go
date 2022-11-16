package templates

import (
	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
	"github.com/onflow/flow-go-sdk"
)

const (
	scriptsPath         = "scripts/"
	readBalanceFilename = "get_balance.cdc"
	readSupplyFilename  = "get_supply.cdc"
)

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(scriptsPath + readBalanceFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, tokenAddr flow.Address, tokenName string) []byte {

	code := assets.MustAssetString(scriptsPath + readSupplyFilename)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}
