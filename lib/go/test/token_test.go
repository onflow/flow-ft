package test

import (
	"testing"

	emulator "github.com/dapperlabs/flow-emulator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-ft/lib/go/contracts"
	"github.com/onflow/flow-ft/lib/go/templates"
)

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

func TestTokenDeployment(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
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

		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value

		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
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
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.UFix64(0_00000000))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0))
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.UFix64(30000_00000000))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0))
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.UFix64(300_00000000))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(700_00000000))

		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(300_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})

	t.Run("Should be able to transfer tokens through a forwarder from a vault", func(t *testing.T) {

		tx := flow.NewTransaction().
			SetScript(templates.GenerateCreateForwarderScript(fungibleAddr, forwardingAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(joshAddress)

		_ = tx.AddArgument(cadence.NewAddress(exampleTokenAddr))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		tx = flow.NewTransaction().
			SetScript(templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.UFix64(300_00000000))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(700_00000000))

		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(300_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
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
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(exampleTokenAddr)

	_ = tx.AddArgument(cadence.UFix64(300_00000000))
	_ = tx.AddArgument(cadence.NewAddress(joshAddress))

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
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(600_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(900_00000000))
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
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(200_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(800_00000000))
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
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.UFix64(0_00000000))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.UFix64(101_00000000))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.UFix64(50_00000000))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(50_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1050_00000000))
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateBurnTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.UFix64(50_00000000))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken"), [][]byte{jsoncdc.MustEncode(cadence.Address(exampleTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(950_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})
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

		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, "UtilityCoin"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(0))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(templates.GenerateMintTokensScript(fungibleAddr, tokenAddr, "UtilityCoin")).
			SetGasLimit(100).
			SetProposalKey(b.ServiceKey().Address, b.ServiceKey().ID, b.ServiceKey().SequenceNumber).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(tokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.UFix64(50_00000000))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, tokenAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), tokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, "UtilityCoin"), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		// Assert that the vaults' balances are correct
		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, "UtilityCoin"), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(50_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1050_00000000))
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
		result, err := b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, tokenAddr, "UtilityCoin"), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		result, err = b.ExecuteScript(templates.GenerateInspectVaultScript(fungibleAddr, badTokenAddr, "BadCoin"), [][]byte{jsoncdc.MustEncode(cadence.Address(badTokenAddr))})
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, balance.(cadence.UFix64), cadence.UFix64(1000_00000000))

		supply := executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, tokenAddr, "UtilityCoin"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1050_00000000))

		supply = executeScriptAndCheck(t, b, templates.GenerateInspectSupplyScript(fungibleAddr, badTokenAddr, "BadCoin"))
		assert.Equal(t, supply.(cadence.UFix64), cadence.UFix64(1000_00000000))
	})
}
