package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"

	"github.com/onflow/flow-ft/lib/go/templates"
)

// Steps:
//
// 1. All the token contracts deploy properly
// 2. Total supply is initialized to 1000.0
func TestTokenDeployment(t *testing.T) {
	b, adapter, accountKeys := newTestSetup(t)

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()

	env := templates.Environment{}

	_ = deployTokenContracts(b, adapter, t, []*flow.AccountKey{exampleTokenAccountKey}, &env)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		script := templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

// Steps:
//
//  1. Create an empty vault that doesn't change the total supply
//     (verify directly and through the metadata view)
func TestTokenSetupAccount(t *testing.T) {
	b, adapter, accountKeys := newTestSetup(t)

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()

	env := templates.Environment{}
	_ = deployTokenContracts(b, adapter, t, []*flow.AccountKey{exampleTokenAccountKey}, &env)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		joshAddress, _, _ := createAccountWithVault(b, adapter, t,
			accountKeys,
			env,
		)

		// Make sure the vault balance is zero
		script := templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assert.Equal(t, CadenceUFix64("0.0"), result)

		// Directly query the total supply to make sure it hasn't changed
		script = templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)

		// Query the total supply via the metadata view to make sure it is also correct
		script = templates.GenerateInspectSupplyViewScript(env)
		supply = executeScriptAndCheck(t, b, script, [][]byte{
			jsoncdc.MustEncode(cadence.Address(joshAddress)),
		})
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

// Steps:
//
// 1. Make sure extra tokens cannot be withdrawn
// 2. Test a regular transfer
// 3. Test a transfer to multiple accounts
func TestTokenExternalTransfers(t *testing.T) {
	b, adapter, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	env := templates.Environment{}
	exampleTokenAddr := deployTokenContracts(b, adapter, t, []*flow.AccountKey{exampleTokenAccountKey}, &env)

	joshAddress, _, joshSigner := createAccountWithVault(b, adapter, t,
		accountKeys,
		env,
	)

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		script := templates.GenerateTransferVaultScript(env)
		tx := createTxWithTemplateAndAuthorizer(b, script, exampleTokenAddr)

		_ = tx.AddArgument(CadenceUFix64("30000.0"))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			true,
		)

		// Assert that the vaults' balances are correct
		// Sender vault
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Receiver Vault
		script = templates.GenerateInspectVaultScript(env)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assert.Equal(t, CadenceUFix64("0.0"), result)
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		script := templates.GenerateTransferVaultScript(env)

		tx := createTxWithTemplateAndAuthorizer(b, script, exampleTokenAddr)

		_ = tx.AddArgument(CadenceUFix64("300.0"))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		// Sender vault
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assert.Equal(t, CadenceUFix64("700.0"), result)

		// Receiver Vault
		script = templates.GenerateInspectVaultScript(env)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assert.Equal(t, CadenceUFix64("300.0"), result)

		// Supply should not have changed
		script = templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer to multiple accounts ", func(t *testing.T) {

		recipient1Address := cadence.Address(joshAddress)
		recipient1Amount := CadenceUFix64("300.0")

		pair := cadence.KeyValuePair{Key: recipient1Address, Value: recipient1Amount}
		recipientPairs := make([]cadence.KeyValuePair, 1)
		recipientPairs[0] = pair

		script := templates.GenerateTransferManyAccountsScript(env)

		tx := flow.NewTransaction().
			SetScript(script).
			SetGasLimit(100).
			SetProposalKey(
				b.ServiceKey().Address,
				b.ServiceKey().Index,
				b.ServiceKey().SequenceNumber,
			).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewDictionary(recipientPairs))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		// Sender vault
		script = templates.GenerateInspectVaultScript(env)
		result, err := b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, CadenceUFix64("400.0"), balance)

		// Receiver Vault
		script = templates.GenerateInspectVaultScript(env)
		result, err = b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, CadenceUFix64("600.0"), balance)

		// Supply should not have changed
		script = templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer tokens with the generic transfer transaction", func(t *testing.T) {

		script := templates.GenerateTransferGenericVaultScript(env)

		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		_ = tx.AddArgument(CadenceUFix64("300.0"))
		_ = tx.AddArgument(cadence.NewAddress(exampleTokenAddr))

		storagePath := cadence.Path{Domain: common.PathDomainStorage, Identifier: "exampleTokenVault"}
		publicPath := cadence.Path{Domain: common.PathDomainPublic, Identifier: "exampleTokenReceiver"}

		_ = tx.AddArgument(storagePath)
		_ = tx.AddArgument(publicPath)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				serviceSigner,
				joshSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assertEqual(t, CadenceUFix64("700.0"), result)

		script = templates.GenerateInspectVaultScript(env)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assertEqual(t, CadenceUFix64("300.0"), result)

	})
}

// Steps:
//
// 1. Setup a forwarder in josh's account to forward to the token address
// 2. Test a transfer from the token account to josh, which will go back to the token account
func TestTokenForwarding(t *testing.T) {
	b, adapter, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	env := templates.Environment{}

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	exampleTokenAddr := deployTokenContracts(b, adapter, t, []*flow.AccountKey{exampleTokenAccountKey}, &env)

	joshAddress, _, joshSigner := createAccountWithVault(b, adapter, t,
		accountKeys,
		env,
	)

	t.Run("Should be able to transfer tokens through a forwarder from a vault", func(t *testing.T) {

		// Setup the forwarder in josh's account to forward to the token addr
		script := templates.GenerateCreateForwarderScript(env)

		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		_ = tx.AddArgument(cadence.NewAddress(exampleTokenAddr))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				serviceSigner,
				joshSigner,
			},
			false,
		)

		// Transfer tokens from the token account to josh
		// which will be forwarded back to the token account
		script = templates.GenerateTransferVaultScript(env)
		tx = createTxWithTemplateAndAuthorizer(b, script, exampleTokenAddr)

		_ = tx.AddArgument(CadenceUFix64("300.0"))
		_ = tx.AddArgument(cadence.NewAddress(joshAddress))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct (the same as before)
		// Token account (sender)
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assertEqual(t, CadenceUFix64("1000.0"), result)

		// Receiver account
		script = templates.GenerateInspectVaultScript(env)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assertEqual(t, CadenceUFix64("0.0"), result)
	})
}

// Steps:
//
// 1. Mint tokens with the ExampleToken Admin (verify that supply and balances are increased)
// 2. Burn tokens, which will decrease the supply and balances
func TestMintingAndBurning(t *testing.T) {
	b, adapter, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	env := templates.Environment{}

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	exampleTokenAddr := deployTokenContracts(b, adapter, t, []*flow.AccountKey{exampleTokenAccountKey}, &env)

	joshAddress, _, _ := createAccountWithVault(b, adapter, t,
		accountKeys,
		env,
	)

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		script := templates.GenerateMintTokensScript(env)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(CadenceUFix64("50.0"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		// token account should not have increased
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Assert that the vaults' balances are correct
		// Josh account should have increased by 50, the amount minted
		script = templates.GenerateInspectVaultScript(env)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("50.0"), result)

		script = templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1050.0"), supply)
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		script := templates.GenerateBurnTokensScript(env)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, exampleTokenAddr)

		_ = tx.AddArgument(CadenceUFix64("50.0"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				exampleTokenAddr,
			},
			[]crypto.Signer{
				serviceSigner,
				exampleTokenSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		// token account balance should have decreased by the burned amount
		script = templates.GenerateInspectVaultScript(env)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assert.Equal(t, CadenceUFix64("950.0"), result)

		// total supply also decreases by the burned amount
		script = templates.GenerateInspectSupplyScript(env)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}
