package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/examples"
	"github.com/onflow/flow-go-sdk/test"

	. "github.com/dapperlabs/flow-emulator/examples"
)

const (
	fungibleTokenContractFile = "./contracts/FungibleToken.cdc"
	flowTokenContractFile     = "./contracts/FlowToken.cdc"
)

func TestTokenDeployment(t *testing.T) {
	b := NewEmulator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	_, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	contractAddr, err := b.CreateAccount(nil, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {

		result, err := b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})
}

func TestCreateToken(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// First, deploy the contract
	tokenCode := ReadFile(fungibleTokenContractFile)
	contractAddr, err := b.CreateAccount(nil, tokenCode)
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateCreateTokenScript(contractAddr)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(joshAddress)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, joshAddress},
			[]crypto.Signer{b.RootKey().Signer(), joshSigner},
			false,
		)

		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})
}

func TestExternalTransfers(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// First, deploy the token contract
	tokenCode := ReadFile(fungibleTokenContractFile)
	contractAccountKey, contractSigner := accountKeys.NewWithSigner()
	contractAddr, err := b.CreateAccount([]*flow.AccountKey{contractAccountKey}, tokenCode)
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(contractAddr)).
		SetGasLimit(20).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	SignAndSubmit(
		t, b, tx,
		[]flow.Address{b.RootKey().Address, joshAddress},
		[]crypto.Signer{b.RootKey().Signer(), joshSigner},
		false,
	)

	t.Run("Shouldn't be able to deposit an empty Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(contractAddr, joshAddress, 0)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(contractAddr, joshAddress, 30000)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(contractAddr, joshAddress, 300)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 700))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 300))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})
}

func TestVaultDestroy(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// First, deploy the token contract
	tokenCode := ReadFile(fungibleTokenContractFile)
	contractAccountKey, contractSigner := accountKeys.NewWithSigner()
	contractAddr, err := b.CreateAccount([]*flow.AccountKey{contractAccountKey}, tokenCode)
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(contractAddr)).
		SetGasLimit(20).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	SignAndSubmit(
		t, b, tx,
		[]flow.Address{b.RootKey().Address, joshAddress},
		[]crypto.Signer{b.RootKey().Signer(), joshSigner},
		false,
	)

	tx = flow.NewTransaction().
		SetScript(GenerateTransferVaultScript(contractAddr, joshAddress, 300)).
		SetGasLimit(20).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(contractAddr)

	SignAndSubmit(
		t, b, tx,
		[]flow.Address{b.RootKey().Address, contractAddr},
		[]crypto.Signer{b.RootKey().Signer(), contractSigner},
		false,
	)

	t.Run("Should subtract tokens from supply when they are destroyed", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(contractAddr, 100)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 600))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 900))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(contractAddr, 100)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(joshAddress)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, joshAddress},
			[]crypto.Signer{b.RootKey().Signer(), joshSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 200))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 800))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

}

func TestMintingAndBurning(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// First, deploy the token contract
	tokenCode := ReadFile(fungibleTokenContractFile)
	contractAccountKey, contractSigner := accountKeys.NewWithSigner()
	contractAddr, err := b.CreateAccount([]*flow.AccountKey{contractAccountKey}, tokenCode)
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(contractAddr)).
		SetGasLimit(20).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	SignAndSubmit(
		t, b, tx,
		[]flow.Address{b.RootKey().Address, joshAddress},
		[]crypto.Signer{b.RootKey().Signer(), joshSigner},
		false,
	)

	t.Run("Shouldn't be able to mint zero tokens", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(contractAddr, joshAddress, 0)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(contractAddr, joshAddress, 101)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(contractAddr, joshAddress, 101)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 0))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(contractAddr, joshAddress, 50)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(GenerateInspectVaultScript(contractAddr, joshAddress, 50))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1050))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateBurnTokensScript(contractAddr, 50)).
			SetGasLimit(20).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(contractAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, contractAddr},
			[]crypto.Signer{b.RootKey().Signer(), contractSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(GenerateInspectVaultScript(contractAddr, contractAddr, 950))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}

		result, err = b.ExecuteScript(GenerateInspectSupplyScript(contractAddr, 1000))
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	})

}
