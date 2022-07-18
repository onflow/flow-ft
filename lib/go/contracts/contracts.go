package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts/...

import (
	"regexp"
	"strings"

	"github.com/onflow/flow-ft/lib/go/contracts/internal/assets"

	_ "github.com/kevinburke/go-bindata"
)

var (
	placeholderFungibleToken            = regexp.MustCompile(`"[^"\s].*/FungibleToken.cdc"`)
	placeholderFungibleTokenV2          = regexp.MustCompile(`"[^"\s].*/FungibleToken-v2.cdc"`)
	placeholderExampleToken             = regexp.MustCompile(`"[^"\s].*/ExampleToken.cdc"`)
	placeholderFungibleTokenV2Interface = regexp.MustCompile(`"[^"\s].*/FungibleToken-v2-ContractInterface.cdc"`)
)

const (
	filenameFungibleToken            = "FungibleToken.cdc"
	filenameFungibleTokenV2          = "FungibleToken-v2.cdc"
	filenameFungibleTokenV2Interface = "FungibleToken-v2-ContractInterface.cdc"
	filenameExampleToken             = "ExampleToken.cdc"
	filenameExampleTokenV2           = "ExampleToken-v2.cdc"
	filenameTokenForwarding          = "utilityContracts/TokenForwarding.cdc"
	filenamePrivateForwarder         = "utilityContracts/PrivateReceiverForwarder.cdc"
)

// FungibleToken returns the FungibleToken contract interface.
func FungibleToken() []byte {
	return assets.MustAsset(filenameFungibleToken)
}

// FungibleTokenV2 returns the FungibleToken-v2 contract.
func FungibleTokenV2() []byte {
	return assets.MustAsset(filenameFungibleTokenV2)
}

// FungibleTokenV2Interface returns the FungibleToken-v2 contract interface.
func FungibleTokenV2Interface(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenameFungibleTokenV2Interface)

	code = placeholderFungibleTokenV2.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}

// ExampleToken returns the ExampleToken contract.
//
// The returned contract will import the FungibleToken interface from the specified address.
func ExampleToken(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenameExampleToken)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}

// ExampleTokenV2 returns the second version of the ExampleToken contract.
//
// The returned contract will import the FungibleToken interface from the specified address.
func ExampleTokenV2(fungibleTokenAddr, fungibleTokenV2InterfaceAddr string) []byte {
	code := assets.MustAssetString(filenameExampleTokenV2)

	code = placeholderFungibleTokenV2.ReplaceAllString(code, "0x"+fungibleTokenAddr)
	code = placeholderFungibleTokenV2Interface.ReplaceAllString(code, "0x"+fungibleTokenV2InterfaceAddr)

	return []byte(code)
}

// CustomToken returns the ExampleToken contract with a custom name.
//
// The returned contract will import the FungibleToken interface from the specified address.
func CustomToken(fungibleTokenAddr, tokenName, storageName, initialBalance string) []byte {
	code := assets.MustAssetString(filenameExampleToken)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	code = strings.ReplaceAll(
		code,
		"ExampleToken",
		tokenName,
	)

	code = strings.ReplaceAll(
		code,
		"exampleToken",
		storageName,
	)

	code = strings.ReplaceAll(
		code,
		"1000.0",
		initialBalance,
	)

	return []byte(code)
}

// TokenForwarding returns the TokenForwarding contract.
//
// The returned contract will import the FungibleToken contract from the specified address.
func TokenForwarding(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenameTokenForwarding)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}

// CustomTokenForwarding returns the TokenForwarding contract for a custom token
//
// The returned contract will import the FungibleToken interface from the specified address.
func CustomTokenForwarding(fungibleTokenAddr, tokenName, storageName string) []byte {
	code := assets.MustAssetString(filenameTokenForwarding)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	code = strings.ReplaceAll(
		code,
		"ExampleToken",
		tokenName,
	)

	code = strings.ReplaceAll(
		code,
		"exampleToken",
		storageName,
	)

	return []byte(code)
}

func PrivateReceiverForwarder(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenamePrivateForwarder)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}
