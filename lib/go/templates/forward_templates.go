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
	defaultPrivateForwardAddr = "PRIVATEFORWARDINGADDRESS"
)

func GenerateDeployPrivateForwardingScript() []byte {
	code := assets.MustAssetString(deployPrivateForwardingFilanems)

	return []byte(code)
}

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreatePrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createPrivateForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

func GenerateSetupAccountPrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(setupAccountPrivateForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

func GenerateTransferPrivateManyAccountsScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(transferPrivateManyAccountsFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}

func GenerateCreateAccountPrivateForwarderScript(fungibleAddr, forwardingAddr, tokenAddr flow.Address, tokenName string) []byte {
	code := assets.MustAssetString(createAccountPrivateForwarderFilename)

	code = replaceAddresses(code, fungibleAddr.String(), tokenAddr.String(), tokenName)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultPrivateForwardAddr,
		"0x"+forwardingAddr.String(),
	)

	return []byte(code)
}
