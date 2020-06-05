package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../src/contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../src/contracts

import (
	"strings"

	"github.com/onflow/flow-ft/contracts/internal/assets"
)

const (
	fungibleTokenFilename       = "FungibleToken.cdc"
	flowTokenFilename           = "FlowToken.cdc"
	exampleTokenFilename        = "ExampleToken.cdc"
	defaultFungibleTokenAddress = "02"
	tokenForwardingFilename     = "TokenForwarding.cdc"
)

// FungibleToken returns the FungibleToken contract interface.
func FungibleToken() []byte {
	return assets.MustAsset(fungibleTokenFilename)
}

// FlowToken returns the FlowToken contract.
//
// The returned contract will import the FungibleToken contract from the specified address.
func FlowToken(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(flowTokenFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddr,
	)

	return []byte(code)
}

// ExampleToken returns the ExampleToken contract.
//
// The returned contract will import the FungibleToken contract from the specified address.
func ExampleToken(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(exampleTokenFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddr,
	)

	return []byte(code)
}

// CustomizableExampleToken returns the ExampleToken contract with a custom Name.
//
// The returned contract will import the FungibleToken contract from the specified address.
func CustomizableExampleToken(fungibleTokenAddr, tokenName string) []byte {
	code := assets.MustAssetString(exampleTokenFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddr,
	)

	code = strings.ReplaceAll(
		code,
		"ExampleToken",
		tokenName,
	)

	return []byte(code)
}

// TokenForwarding returns the TokenForwarding contract.
//
// The returned contract will import the FungibleToken contract from the specified address.
func TokenForwarding(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(tokenForwardingFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddr,
	)

	return []byte(code)
}

// CustomizableTokenForwarding returns the TokenForwarding contract for a custom token
//
// The returned contract will import the FungibleToken contract from the specified address.
func CustomizableTokenForwarding(fungibleTokenAddr, tokenName string) []byte {
	code := assets.MustAssetString(tokenForwardingFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultFungibleTokenAddress,
		"0x"+fungibleTokenAddr,
	)

	code = strings.ReplaceAll(
		code,
		"ExampleToken",
		tokenName,
	)

	return []byte(code)
}
