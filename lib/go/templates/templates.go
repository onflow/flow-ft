package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/onflow/flow-go-sdk"
)

var (
	defaultTokenName    = regexp.MustCompile(`ExampleToken`)
	defaultTokenStorage = regexp.MustCompile(`exampleToken`)
)

var (
	placeholderFungibleToken = regexp.MustCompile(`"[^"\s].*/FungibleToken.cdc"`)
	placeholderExampleToken  = regexp.MustCompile(`"[^"\s].*/ExampleToken.cdc"`)
	placeholderForwarding    = regexp.MustCompile(`"[^"\s].*/TokenForwarding.cdc"`)
	placeholderMetadataViews = regexp.MustCompile(`"[^"\s].*/MetadataViews.cdc"`)
)

func replaceAddresses(code string, ftAddress, tokenAddress, forwardingAddress, metadataViewsAddress flow.Address, tokenName string) []byte {
	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())
	code = placeholderExampleToken.ReplaceAllString(code, "0x"+tokenAddress.String())
	code = placeholderForwarding.ReplaceAllString(code, "0x"+forwardingAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataViewsAddress.String())

	storageName := MakeFirstLowerCase(tokenName)
	code = defaultTokenName.ReplaceAllString(code, tokenName)
	code = defaultTokenStorage.ReplaceAllString(code, storageName)

	return []byte(code)
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
