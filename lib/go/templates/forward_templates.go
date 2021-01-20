package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"strings"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	createPrivateForwarderFilename       = "privateForwarder/create_private_forwarder.cdc"
	setupAccountPrivateForwarderFilename = "privateForwarder/setup_and_create_forwarder.cdc"
	transferPrivateManyAccountsFilename  = "privateForwarder/transfer_private_many_accounts.cdc"
)

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreatePrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createPrivateForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultForwardingAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

func GenerateSetupAccountPrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(setupAccountPrivateForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultForwardingAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

func GenerateTransferPrivateManyAccountsScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(transferPrivateManyAccountsFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultForwardingAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}
