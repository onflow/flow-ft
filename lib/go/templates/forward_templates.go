package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"github.com/onflow/flow-ft/lib/go/templates/internal/assets"
)

const (
	deployPrivateForwardingFilanems = "privateForwarder/deploy_forwarder_contract.cdc"

	createPrivateForwarderFilename        = "privateForwarder/create_private_forwarder.cdc"
	setupAccountPrivateForwarderFilename  = "privateForwarder/setup_and_create_forwarder.cdc"
	transferPrivateManyAccountsFilename   = "privateForwarder/transfer_private_many_accounts.cdc"
	createAccountPrivateForwarderFilename = "privateForwarder/create_account_private_forwarder.cdc"
)

func GenerateDeployPrivateForwardingScript() []byte {
	code := assets.MustAssetString(deployPrivateForwardingFilanems)

	return []byte(code)
}

// GenerateCreateForwarderScript creates a script that instantiates
// a new forwarder instance in an account
func GenerateCreatePrivateForwarderScript(env Environment) []byte {
	code := assets.MustAssetString(createPrivateForwarderFilename)

	return []byte(ReplaceAddresses(code, env))
}

func GenerateSetupAccountPrivateForwarderScript(env Environment) []byte {
	code := assets.MustAssetString(setupAccountPrivateForwarderFilename)

	return []byte(ReplaceAddresses(code, env))
}

func GenerateTransferPrivateManyAccountsScript(env Environment) []byte {
	code := assets.MustAssetString(transferPrivateManyAccountsFilename)

	return []byte(ReplaceAddresses(code, env))
}

func GenerateCreateAccountPrivateForwarderScript(env Environment) []byte {
	code := assets.MustAssetString(createAccountPrivateForwarderFilename)

	return []byte(ReplaceAddresses(code, env))
}
