package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"strings"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	deployPrivateForwardingFilanems = "privateForwarder/deploy_forwarder_contract.cdc"

	createPrivateForwarderFilename        = "privateForwarder/create_private_forwarder.cdc"
	setupAccountPrivateForwarderFilename  = "privateForwarder/setup_and_create_forwarder.cdc"
	transferPrivateManyAccountsFilename   = "privateForwarder/transfer_private_many_accounts.cdc"
	createAccountPrivateForwarderFilename = "privateForwarder/create_account_private_forwarder.cdc"
)

const (
	defaultPrivateForwardAddr = "\"../../contracts/PrivateReceiverForwarder.cdc\""
)

func GenerateDeployPrivateForwardingScript() []byte {
	code := assets.MustAssetString(deployPrivateForwardingFilanems)

	return []byte(code)
}

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreatePrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createPrivateForwarderFilename)

	code = strings.ReplaceAll(
		code,
		defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

func GenerateSetupAccountPrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(setupAccountPrivateForwarderFilename)

	code = strings.ReplaceAll(
		code,
		defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

func GenerateTransferPrivateManyAccountsScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(transferPrivateManyAccountsFilename)

	code = strings.ReplaceAll(
		code,
		defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}

func GenerateCreateAccountPrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createAccountPrivateForwarderFilename)

	code = strings.ReplaceAll(
		code,
		defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return replaceAddresses(code, fungibleAddr, tokenAddr, flow.EmptyAddress, flow.EmptyAddress, tokenName)
}
