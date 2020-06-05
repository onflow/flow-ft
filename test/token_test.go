package test

import (
	"strings"
	"testing"

	emulator "github.com/dapperlabs/flow-emulator"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-ft/contracts"
)

func TestTokenDeployment(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestCreateToken(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken")).
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

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestExternalTransfers(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken")).
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
			SetScript(GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 1000))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0))
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 30000)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 1000))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0))
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 300)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 700))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 300))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func TestVaultDestroy(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken")).
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
		SetScript(GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 300)).
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
			SetScript(GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 100)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 600))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 900))
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 100)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 200))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 800))
	})

}

func TestMintingAndBurning(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken")).
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
			SetScript(GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 101)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 50)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, joshAddress, "ExampleToken", "exampleToken", 50))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1050))
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateBurnTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 50)).
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
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, exampleTokenAddr, "ExampleToken", "exampleToken", 950))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 1000))
	})
}

func DeployTokenContracts(b *emulator.Blockchain, t *testing.T, key []*flow.AccountKey) (flow.Address, flow.Address) {

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	exampleTokenCode := readFile("../src/contracts/ExampleToken.cdc")
	codeWithFTAddr := strings.ReplaceAll(string(exampleTokenCode), "02", fungibleAddr.String())

	tokenAddr, err := b.CreateAccount(key, []byte(codeWithFTAddr))
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	return fungibleAddr, tokenAddr
}

func TestCreateCustomToken(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := contracts.FungibleToken()
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	exampleTokenCode := contracts.CustomizableExampleToken(fungibleAddr.String(), "UtilityCoin")

	tokenAddr, err := b.CreateAccount([]*flow.AccountKey{exampleTokenAccountKey}, exampleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateCreateTokenScript(fungibleAddr, tokenAddr, "UtilityCoin", "utilityCoin")).
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

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, tokenAddr, joshAddress, "UtilityCoin", "utilityCoin", 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin", 1000))
	})
}
