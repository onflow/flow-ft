package test

import (
	"context"
	"testing"

	"github.com/onflow/flow-ft/lib/go/templates"

	"github.com/onflow/flow-emulator/adapters"
	"github.com/onflow/flow-emulator/emulator"
	"github.com/onflow/flow-go-sdk/crypto"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk"

	nftcontracts "github.com/onflow/flow-nft/lib/go/contracts"

	"github.com/onflow/flow-ft/lib/go/contracts"
)

// Deploys the FungibleToken, ExampleToken, and TokenForwarding contracts
// to different accounts and returns their addresses
func deployTokenContracts(
	b emulator.Emulator,
	adapter *adapters.SDKAdapter,
	t *testing.T,
	key []*flow.AccountKey,
	env *templates.Environment,
) (
	tokenAddr flow.Address,
) {
	var err error

	// Deploy the ViewResolver contract
	resolverAddress, err := adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "ViewResolver",
				Source: string(nftcontracts.Resolver()),
			},
		},
	)
	assert.NoError(t, err)
	env.ViewResolverAddress = resolverAddress.Hex()

	// Deploy the NonFungibleToken contract
	nonFungibleTokenCode := nftcontracts.NonFungibleToken(resolverAddress)
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
	fungibleTokenCode := contracts.FungibleToken(resolverAddress.String())
	fungibleAddr, err := adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleToken",
				Source: string(fungibleTokenCode),
			},
		},
	)
	assert.NoError(t, err)
	env.FungibleTokenAddress = fungibleAddr.Hex()

	// Deploy the MetadataViews contract
	metadataViewsCode := nftcontracts.MetadataViews(fungibleAddr, nftAddress, resolverAddress)
	metadataViewsAddr, err := adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "MetadataViews",
				Source: string(metadataViewsCode),
			},
		},
	)
	assert.NoError(t, err)
	env.MetadataViewsAddress = metadataViewsAddr.Hex()

	// Deploy the FungibleTokenMetadataViews contract
	fungibleTokenMetadataViewsCode := contracts.FungibleTokenMetadataViews(fungibleAddr.String(), metadataViewsAddr.String(), resolverAddress.String())
	fungibleMetadataViewsAddr, err := adapter.CreateAccount(context.Background(),
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleTokenMetadataViews",
				Source: string(fungibleTokenMetadataViewsCode),
			},
		},
	)
	assert.NoError(t, err)
	env.FungibleTokenMetadataViewsAddress = fungibleMetadataViewsAddr.Hex()

	// Deploy the ExampleToken contract
	exampleTokenCode := contracts.ExampleToken(fungibleAddr.String(), metadataViewsAddr.String(), fungibleMetadataViewsAddr.String(), resolverAddress.String(), "")
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
	env.ExampleTokenAddress = tokenAddr.Hex()

	// Deploy the TokenForwarding contract
	forwardingCode := contracts.TokenForwarding(fungibleAddr.String())
	forwardingAddr, err := adapter.CreateAccount(context.Background(),
		key,
		[]sdktemplates.Contract{
			{
				Name:   "TokenForwarding",
				Source: string(forwardingCode),
			},
		},
	)
	assert.NoError(t, err)
	env.TokenForwardingAddress = forwardingAddr.Hex()

	// Deploy the FungibleTokenSwitchboard contract
	switchboardCode := contracts.FungibleTokenSwitchboard(fungibleAddr.String())
	switchboardAddr, err := adapter.CreateAccount(context.Background(),
		key,
		[]sdktemplates.Contract{
			{
				Name:   "FungibleTokenSwitchboard",
				Source: string(switchboardCode),
			},
		},
	)
	assert.NoError(t, err)
	env.SwitchboardAddress = switchboardAddr.Hex()

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	return tokenAddr
}

func createAccountWithVault(
	b emulator.Emulator,
	adapter *adapters.SDKAdapter,
	t *testing.T,
	keys *test.AccountKeys,
	env templates.Environment,
) (
	flow.Address, *flow.AccountKey, crypto.Signer,
) {

	newAccountKey, newSigner := keys.NewWithSigner()
	newAddress, _ := adapter.CreateAccount(context.Background(), []*flow.AccountKey{newAccountKey}, nil)

	serviceSigner, _ := b.ServiceKey().Signer()

	// Setup new account with an empty vault
	script := templates.GenerateCreateTokenScript(env)
	tx := createTxWithTemplateAndAuthorizer(b, script, newAddress)
	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			newAddress,
		},
		[]crypto.Signer{
			serviceSigner,
			newSigner,
		},
		false,
	)
	return newAddress, newAccountKey, newSigner
}
