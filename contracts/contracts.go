package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../src/contracts -o assets.go -pkg contracts -nometadata -nomemcopy ../src/contracts

import (
	"strings"

	"github.com/onflow/flow-go-sdk"
)

const (
	FungibleTokenContractFilename = "FungibleToken.cdc"
	FlowTokenContractFilename     = "FlowToken.cdc"
	defaultFungibleTokenAddress   = "02"
)

// FungibleToken returns the FungibleToken contract interface.
func FungibleToken() ([]byte, error) {
	return Asset(FungibleTokenContractFilename)
}

// FlowToken returns the FlowToken contract, importing the
// FungibleToken contract interface from the specified Flow address.
func FlowToken(fungibleTokenAddress flow.Address) ([]byte, error) {
	tpl, err := AssetString(FlowTokenContractFilename)
	if err != nil {
		return nil, err
	}

	code := strings.ReplaceAll(
		tpl,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddress.Hex(),
	)

	return []byte(code), nil
}
