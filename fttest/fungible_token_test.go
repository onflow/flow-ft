package fttest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/examples"
	"github.com/onflow/flow-go-sdk/test"
)

const (
	fungibleTokenContractFile = "../contracts/FungibleToken.cdc"
	flowTokenContractFile     = "../contracts/FlowToken.cdc"
)

func TestTokenDeployment(t *testing.T) {
	b := NewEmulator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	flowAddr, err := b.CreateAccount(nil, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestCreateToken(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	flowAddr, err := b.CreateAccount(nil, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(joshAddress)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, joshAddress},
			[]crypto.Signer{b.RootKey().Signer(), joshSigner},
			false,
		)

		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestExternalTransfers(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
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
			SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 0)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 30000)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 300)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 700))

		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 300))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestVaultDestroy(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
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
		SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 300)).
		SetGasLimit(100).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(flowAddr)

	SignAndSubmit(
		t, b, tx,
		[]flow.Address{b.RootKey().Address, flowAddr},
		[]crypto.Signer{b.RootKey().Signer(), flowSigner},
		false,
	)

	t.Run("Should subtract tokens from supply when they are destroyed", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(fungibleAddr, flowAddr, 100)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 600))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 900))
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(fungibleAddr, flowAddr, 100)).
			SetGasLimit(100).
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
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 200))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 800))
	})

}

func TestMintingAndBurning(t *testing.T) {
	b := NewEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := examples.ReadFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := examples.ReadFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
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
			SetScript(GenerateMintTokensScript(fungibleAddr, flowAddr, joshAddress, 0)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, flowAddr, joshAddress, 101)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, flowAddr, joshAddress, 50)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 50))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1050))
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateBurnTokensScript(fungibleAddr, flowAddr, 50)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		SignAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		ExecuteScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 950))

		ExecuteScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}
