module github.com/onflow/flow-ft/lib/go/test

go 1.16

require (
	github.com/onflow/cadence v0.24.1
	github.com/onflow/flow-emulator v0.33.2
	github.com/onflow/flow-ft/lib/go/contracts v0.5.0
	github.com/onflow/flow-ft/lib/go/templates v0.0.0-00010101000000-000000000000
	github.com/onflow/flow-go-sdk v0.26.1
	github.com/stretchr/testify v1.7.1
)

replace github.com/onflow/flow-ft/lib/go/contracts => ../contracts

replace github.com/onflow/flow-ft/lib/go/templates => ../templates
