package test

import (
	"testing"

	emulator "github.com/dapperlabs/flow-emulator"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-ft/lib/go/contracts"
	"github.com/onflow/flow-ft/lib/go/templates"
)

func TestTokenDeployment(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestCreateToken(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestExternalTransfers(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, forwardingAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, joshAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
		false,
	)

	t.Run("Shouldn't be able to deposit an empty Vault", func(t *testing.T) {

		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			AddArgument(cadence.UFix64(0_00000000)).
			AddArgument(cadence.NewAddress(joshAddress)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 1000))

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0))
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			AddArgument(cadence.UFix64(30000_00000000)).
			AddArgument(cadence.NewAddress(joshAddress)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 1000))

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0))
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			AddArgument(cadence.UFix64(300_00000000)).
			AddArgument(cadence.NewAddress(joshAddress)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 700))

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 300))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})

	t.Run("Should be able to transfer tokens through a forwarder from a vault", func(t *testing.T) {

		tx := flow.NewTransaction().
			SetScript(templates.GenerateCreateForwarderScript(fungibleAddr, forwardingAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		tx = flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			AddArgument(cadence.UFix64(300_00000000)).
			AddArgument(cadence.NewAddress(joshAddress)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 700))

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 300))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestVaultDestroy(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, joshAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
		false,
	)

	tx = flow.NewTransaction().
		SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
		AddArgument(cadence.UFix64(300_00000000)).
		AddArgument(cadence.NewAddress(joshAddress)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(exampleTokenAddr)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
		[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
		false,
	)

	t.Run("Should subtract tokens from supply when they are destroyed", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 100)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 600))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 900))
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 100)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 200))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 800))
	})

}

func TestMintingAndBurning(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, joshAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
		false,
	)

	t.Run("Shouldn't be able to mint zero tokens", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 101)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 0))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 50)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", 50))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1050))
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateBurnTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 50)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", 950))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func DeployTokenContracts(b *emulator.Blockchain, t *testing.T, key []*flow.AccountKey) (flow.Address, flow.Address, flow.Address) {

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	exampleTokenCode := contracts.ExampleToken(fungibleAddr.String())

	tokenAddr, err := b.CreateAccount(key, []byte(exampleTokenCode))
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	forwardingCode := contracts.TokenForwarding(fungibleAddr.String())

	forwardingAddr, err := b.CreateAccount(key, []byte(forwardingCode))
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	return fungibleAddr, tokenAddr, forwardingAddr
}

func TestCreateCustomToken(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, tokenSigner := accountKeys.NewWithSigner()
	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	exampleTokenCode := contracts.CustomToken(fungibleAddr.String(), "UtilityCoin", "utilityCoin", "1000.0")

	tokenAddr, err := b.CreateAccount([]*flow.AccountKey{exampleTokenAccountKey}, exampleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	badTokenCode := contracts.CustomToken(fungibleAddr.String(), "BadCoin", "badCoin", "1000.0")
	badTokenAccountKey, _ := accountKeys.NewWithSigner()
	badTokenAddr, err := b.CreateAccount([]*flow.AccountKey{badTokenAccountKey}, badTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateCreateTokenScript(fungibleAddr, tokenAddr, "UtilityCoin")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, joshAddress, "UtilityCoin", 0))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin", 1000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, tokenAddr, joshAddress, "UtilityCoin", 50)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(tokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, tokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), tokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, tokenAddr, "UtilityCoin", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, joshAddress, "UtilityCoin", 50))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin", 1050))
	})

	t.Run("Shouldn't be able to transfer token from a vault to a differenly typed vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferInvalidVaultScript(fungibleAddr, tokenAddr, badTokenAddr, badTokenAddr, "UtilityCoin", "BadCoin", 20)).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(tokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, tokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), tokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, tokenAddr, "UtilityCoin", 1000))

		executeScriptAndCheck(t, b, templates.GenerateInspectVaultScript(fungibleAddr, badTokenAddr, badTokenAddr, "BadCoin", 1000))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin", 1050))

		executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, badTokenAddr, "BadCoin", 1000))
	})
}
