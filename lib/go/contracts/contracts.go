package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts/...

import (
	"regexp"
	"strings"

	"github.com/onflow/flow-ft/lib/go/contracts/internal/assets"

	_ "github.com/kevinburke/go-bindata"
)

var (
	placeholderFungibleToken   = regexp.MustCompile(`"FungibleToken"`)
	placeholderExampleToken    = regexp.MustCompile(`"ExampleToken"`)
	placeholderMetadataViews   = regexp.MustCompile(`"MetadataViews"`)
	placeholderFTMetadataViews = regexp.MustCompile(`"FungibleTokenMetadataViews"`)
	placeholderViewResolver    = regexp.MustCompile(`"ViewResolver"`)
	placeholderMultipleVaults  = regexp.MustCompile(`"MultipleVaults"`)
)

const (
	filenameFungibleToken    = "FungibleToken.cdc"
	filenameFungibleTokenV2  = "FungibleToken-v2.cdc"
	filenameExampleToken     = "ExampleToken.cdc"
	filenameExampleTokenV2   = "ExampleToken-v2.cdc"
	filenameTokenForwarding  = "utility/TokenForwarding.cdc"
	filenamePrivateForwarder = "utility/PrivateReceiverForwarder.cdc"
	filenameFTSwitchboard    = "FungibleTokenSwitchboard.cdc"
	filenameFTMetadataViews  = "FungibleTokenMetadataViews.cdc"
	filenameViewResolver     = "utility/ViewResolver.cdc"
	filenameMultipleVaults   = "MultipleVaults.cdc"
)

// FungibleToken returns the FungibleToken contract interface.
func FungibleToken() []byte {
	code := assets.MustAssetString(filenameFungibleToken)

	return []byte(code)
}

// FungibleTokenV2 returns the FungibleToken-v2 contract.
func FungibleTokenV2(resolverAddr string) []byte {
	code := assets.MustAssetString(filenameFungibleTokenV2)

	code = placeholderViewResolver.ReplaceAllString(code, "0x"+resolverAddr)

	return []byte(code)
}

// FungibleToken returns the FungibleToken contract interface.
func FungibleTokenMetadataViews(fungibleTokenAddr, metadataViewsAddr, viewResolverAddr string) []byte {
	code := assets.MustAssetString(filenameFTMetadataViews)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataViewsAddr)
	code = placeholderViewResolver.ReplaceAllString(code, "0x"+viewResolverAddr)

	return []byte(code)
}

// FungibleTokenSwitchboard returns the FungibleTokenSwitchboard contract.
func FungibleTokenSwitchboard(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenameFTSwitchboard)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}

// MultipleVaults returns the MultipleVaults contract.
func MultipleVaults(fungibleTokenAddr string) []byte {
	code := assets.MustAssetString(filenameMultipleVaults)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)

	return []byte(code)
}

// ExampleToken returns the ExampleToken contract.
//
// The returned contract will import the FungibleToken interface from the specified address.
func ExampleToken(fungibleTokenAddr, metadataViewsAddr, ftMetadataViewsAddr string) []byte {
	code := assets.MustAssetString(filenameExampleToken)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataViewsAddr)
	code = placeholderFTMetadataViews.ReplaceAllString(code, "0x"+ftMetadataViewsAddr)

	return []byte(code)
}

// ExampleTokenV2 returns the second version of the ExampleToken contract.
//
// The returned contract will import the FungibleToken interface from the specified address.
func ExampleTokenV2(fungibleTokenAddr, metadataViewsAddr, ftMetadataViewsAddr, viewResolverAddr, multipleVaultsAddr string) []byte {
	code := assets.MustAssetString(filenameExampleTokenV2)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataViewsAddr)
	code = placeholderFTMetadataViews.ReplaceAllString(code, "0x"+ftMetadataViewsAddr)
	code = placeholderViewResolver.ReplaceAllString(code, "0x"+viewResolverAddr)
	code = placeholderMultipleVaults.ReplaceAllString(code, "0x"+multipleVaultsAddr)

	return []byte(code)
}

// CustomToken returns the ExampleToken contract with a custom name.
//
// The returned contract will import the FungibleToken interface from the specified address.
func CustomToken(fungibleTokenAddr,
	metadataViewsAddr,
	ftMetadataViewsAddr,
	tokenName,
	storageName,
	initialBalance string) []byte {

	code := assets.MustAssetString(filenameExampleToken)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+fungibleTokenAddr)
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataViewsAddr)
	code = placeholderFTMetadataViews.ReplaceAllString(code, "0x"+ftMetadataViewsAddr)

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
