module github.com/onflow/flow-ft/lib/go/test

go 1.13

require (
	github.com/dapperlabs/flow-emulator v0.10.0
	github.com/onflow/cadence v0.8.2
	github.com/onflow/flow v0.1.4-0.20200601215056-34a11def1d6b // indirect
	github.com/onflow/flow-ft/contracts v0.1.3
	github.com/onflow/flow-ft/lib/go/contracts v0.1.3
	github.com/onflow/flow-ft/lib/go/templates v0.0.0-00010101000000-000000000000
	github.com/onflow/flow-go-sdk v0.9.0
	github.com/stretchr/testify v1.6.1
	github.com/vektra/mockery v1.1.2 // indirect
	github.com/zenazn/goji v0.9.0 // indirect
)

replace github.com/onflow/flow-ft/lib/go/contracts => ../contracts

replace github.com/onflow/flow-ft/lib/go/templates => ../templates
