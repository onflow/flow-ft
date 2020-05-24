module github.com/onflow/flow-ft/test

go 1.13

require (
	github.com/dapperlabs/flow-emulator v1.0.0-alpha.7.0.20200522225230-51e97f42cb03
	github.com/onflow/cadence v0.3.0-beta4.0.20200524043105-6b94cabe6a65
	github.com/onflow/flow-go-sdk v0.3.0-beta1
	github.com/stretchr/testify v1.5.1
)

replace github.com/dapperlabs/flow-emulator => ../../../go/src/github.com/dapperlabs/flow-emulator

replace github.com/onflow/flow-go-sdk => ../../../go/src/github.com/onflow/flow-go-sdk
