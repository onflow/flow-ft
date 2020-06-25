module github.com/onflow/flow-ft/lib/go/test

go 1.13

require (
	github.com/dapperlabs/flow-emulator v0.4.0
	github.com/onflow/cadence v0.4.0
	github.com/onflow/flow-ft/contracts v0.1.3
	github.com/onflow/flow-ft/lib/go/contracts v0.1.3
	github.com/onflow/flow-ft/lib/go/templates v0.0.0-00010101000000-000000000000
	github.com/onflow/flow-go-sdk v0.4.1
	github.com/onflow/flow/protobuf/go/flow v0.1.5-0.20200611205353-548107cc9aca // indirect
	github.com/stretchr/testify v1.6.1
)

replace github.com/onflow/flow-ft/lib/go/contracts => ../contracts
replace github.com/onflow/flow-ft/lib/go/templates => ../templates
