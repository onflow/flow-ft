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
		withHexPrefix(env.FungibleTokenAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderExampleToken,
		withHexPrefix(env.ExampleTokenAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderForwarding,
		withHexPrefix(env.TokenForwardingAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderPrivateForwardAddr,
		withHexPrefix(env.PrivateForwardingAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderMetadataViews,
		withHexPrefix(env.MetadataViewsAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderFTMetadataViews,
		withHexPrefix(env.FungibleTokenMetadataViewsAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderViewResolver,
		withHexPrefix(env.ViewResolverAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderSwitchboard,
		withHexPrefix(env.SwitchboardAddress),
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
