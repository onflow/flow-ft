package test

import (
	"context"
	"testing"

	"github.com/onflow/flow-emulator/adapters"
	"github.com/onflow/flow-emulator/emulator"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk"

	nftcontracts "github.com/onflow/flow-nft/lib/go/contracts"

	"github.com/onflow/flow-ft/lib/go/contracts"
)

// Deploys the FungibleToken, ExampleToken, and TokenForwarding contracts
// to different accounts and returns their addresses
func DeployTokenContracts(
	b emulator.Emulator,
	adapter *adapters.SDKAdapter,
	t *testing.T,
	key []*flow.AccountKey,
) (
	fungibleAddr flow.Address,
	tokenAddr flow.Address,
	forwardingAddr flow.Address,
	metadataViewsAddr flow.Address,
	fungibleMetadataViewsAddr flow.Address,
) {
	var err error

	// Deploy the NonFungibleToken contract
	nonFungibleTokenCode := nftcontracts.NonFungibleToken()
	nftAddress, err := adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "NonFungibleToken",
				Source: string(nonFungibleTokenCode),
			},
		},
	)
	assert.NoError(t, err)

	// Deploy the FungibleToken contract
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err = adapter.CreateAccount(context.Background(),
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

	// Deploy the MetadataViews contract
	metadataViewsCode := nftcontracts.MetadataViews(fungibleAddr, nftAddress)
	metadataViewsAddr, err = adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "MetadataViews",
				Source: string(metadataViewsCode),
			},
		},
	)
	assert.NoError(t, err)

	// Deploy the FungibleTokenMetadataViews contract
	fungibleTokenMetadataViewsCode := contracts.FungibleTokenMetadataViews(fungibleAddr.String(), metadataViewsAddr.String())
	fungibleMetadataViewsAddr, err = adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleTokenMetadataViews",
				Source: string(fungibleTokenMetadataViewsCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Deploy the ExampleToken contract
	exampleTokenCode := contracts.ExampleToken(fungibleAddr.String(), metadataViewsAddr.String(), fungibleMetadataViewsAddr.String())
	tokenAddr, err = adapter.CreateAccount(context.Background(),
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
	forwardingAddr, err = adapter.CreateAccount(context.Background(),
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

	// Deploy the FungibleTokenSwitchboard contract
	switchboardCode := contracts.FungibleTokenSwitchboard(fungibleAddr.String())
	_, err = adapter.CreateAccount(context.Background(),
		key,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleTokenSwitchboard",
				Source: string(switchboardCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	return fungibleAddr, tokenAddr, forwardingAddr, metadataViewsAddr, fungibleMetadataViewsAddr
}
