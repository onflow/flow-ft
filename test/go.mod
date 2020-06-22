module github.com/onflow/flow-ft/test

go 1.13

require (
	github.com/dapperlabs/flow-emulator v0.4.0
	github.com/onflow/cadence v0.4.0
	github.com/onflow/flow-ft/contracts v0.0.0-20200525235630-0e8024a483ce
	github.com/onflow/flow-go-sdk v0.4.1
	github.com/onflow/flow/protobuf/go/flow v0.1.5-0.20200611205353-548107cc9aca // indirect
	github.com/stretchr/testify v1.6.1
)

replace github.com/onflow/flow-ft/contracts => ../contracts
