module github.com/onflow/flow-ft/lib/go/test

go 1.13

require (
	github.com/onflow/cadence v0.12.3
	github.com/onflow/flow-emulator v0.14.0-beta3
	github.com/onflow/flow-ft/lib/go/contracts v0.4.0
	github.com/onflow/flow-ft/lib/go/templates v0.0.0-00010101000000-000000000000
	github.com/onflow/flow-go-sdk v0.14.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/onflow/flow-ft/lib/go/contracts => ../contracts

replace github.com/onflow/flow-ft/lib/go/templates => ../templates
