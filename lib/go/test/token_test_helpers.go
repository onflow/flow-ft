package test

import (
	"testing"

	"github.com/onflow/flow-emulator"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-ft/lib/go/contracts"
)

// Deploys the FungibleToken, ExampleToken, and TokenForwarding contracts
// to different accounts and returns their addresses
func DeployTokenContracts(
	b *emulator.Blockchain,
	t *testing.T,
	key []*flow.AccountKey,
) (
	fungibleAddr flow.Address,
	tokenAddr flow.Address,
	forwardingAddr flow.Address,
) {
	var err error

	// Deploy the FungibleToken contract
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err = b.CreateAccount(
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleToken",
				Source: string(fungibleTokenCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Deploy the ExampleToken contract
	exampleTokenCode := contracts.ExampleToken(fungibleAddr.String())
	tokenAddr, err = b.CreateAccount(
		key,
		[]sdktemplates.Contract{
			{
				Name:   "ExampleToken",
				Source: string(exampleTokenCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Deploy the TokenForwarding contract
	forwardingCode := contracts.TokenForwarding(fungibleAddr.String())
	forwardingAddr, err = b.CreateAccount(
		key,
		[]sdktemplates.Contract{
			{
				Name:   "TokenForwarding",
				Source: string(forwardingCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	return fungibleAddr, tokenAddr, forwardingAddr
}
