package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var (
	defaultTokenName    = regexp.MustCompile(`ExampleToken`)
	defaultTokenStorage = regexp.MustCompile(`exampleToken`)
)

var (
	placeholderFungibleToken      = "\"FungibleToken\""
	placeholderExampleToken       = "\"ExampleToken\""
	placeholderForwarding         = "\"TokenForwarding\""
	placeholderPrivateForwardAddr = "\"PrivateReceiverForwarder\""
	placeholderSwitchboard        = "\"FungibleTokenSwitchboard\""
	placeholderMetadataViews      = "\"MetadataViews\""
	placeholderFTMetadataViews    = "\"FungibleTokenMetadataViews\""
	placeholderViewResolver       = "\"ViewResolver\""
	placeholderBurner             = "\"Burner\""
	exampleTokenImport            = "ExampleToken from "
	metadataViewsImport           = "MetadataViews from "
	ftMetadataViewsImport         = "FungibleTokenMetadataViews from "
	burnerImport                  = "Burner from "
	fungibleTokenImport           = "FungibleToken from "
	viewResolverImport            = "ViewResolver from "
	switchboardImport             = "FungibleTokenSwitchboard from "
	forwardingImport              = "TokenForwarding from "
	privateForwardingImport       = "PrivateReceiverForwarder from "
)

type Environment struct {
	Network                           string
	FungibleTokenAddress              string
	ExampleTokenAddress               string
	TokenForwardingAddress            string
	PrivateForwardingAddress          string
	MetadataViewsAddress              string
	FungibleTokenMetadataViewsAddress string
	ViewResolverAddress               string
	BurnerAddress                     string
	SwitchboardAddress                string
}

func withHexPrefix(address string) string {
	if address == "" {
		return ""
	}

	if address[0:2] == "0x" {
		return address
	}

	return fmt.Sprintf("0x%s", address)
}

func ReplaceAddresses(code string, env Environment) string {

	code = strings.ReplaceAll(
		code,
		placeholderFungibleToken,
		fungibleTokenImport+withHexPrefix(env.FungibleTokenAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderExampleToken,
		exampleTokenImport+withHexPrefix(env.ExampleTokenAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderForwarding,
		forwardingImport+withHexPrefix(env.TokenForwardingAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderPrivateForwardAddr,
		privateForwardingImport+withHexPrefix(env.PrivateForwardingAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderMetadataViews,
		metadataViewsImport+withHexPrefix(env.MetadataViewsAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderFTMetadataViews,
		ftMetadataViewsImport+withHexPrefix(env.FungibleTokenMetadataViewsAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderViewResolver,
		viewResolverImport+withHexPrefix(env.ViewResolverAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderBurner,
		burnerImport+withHexPrefix(env.BurnerAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderSwitchboard,
		switchboardImport+withHexPrefix(env.SwitchboardAddress),
	)

	// storageName := MakeFirstLowerCase(tokenName)
	// code = defaultTokenName.ReplaceAllString(code, tokenName)
	// code = defaultTokenStorage.ReplaceAllString(code, storageName)

	return code
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
