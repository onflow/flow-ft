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
func FungibleToken() []byte {
	return MustAsset(FungibleTokenContractFilename)
}

// FlowToken returns the FlowToken contract. importing the
//
// The returned contract will import the FungibleToken contract from the specified address.
func FlowToken(fungibleTokenAddress flow.Address) []byte {
	code := MustAssetString(FlowTokenContractFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddress.Hex(),
	)

	return []byte(code)
}
