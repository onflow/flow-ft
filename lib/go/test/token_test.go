package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"

	"github.com/onflow/flow-ft/lib/go/templates"
)

func TestTokenDeployment(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _, _ := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		script := templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

func TestCreateToken(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _, metadataViewsAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		script := templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, metadataViewsAddr, "ExampleToken")
		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

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

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("0.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

func TestExternalTransfers(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, forwardingAddr, metadataViewsAddr :=
		DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	script := templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, metadataViewsAddr, "ExampleToken")
	tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

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

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		script := templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("0.0"), result)
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		script := templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")

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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("700.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("300.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer to multiple accounts ", func(t *testing.T) {

		recipient1Address := cadence.Address(joshAddress)
		recipient1Amount := CadenceUFix64("300.0")

		pair := cadence.KeyValuePair{Key: recipient1Address, Value: recipient1Amount}
		recipientPairs := make([]cadence.KeyValuePair, 1)
		recipientPairs[0] = pair

		script := templates.GenerateTransferManyAccountsScript(fungibleAddr, exampleTokenAddr, "ExampleToken")

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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer tokens through a forwarder from a vault", func(t *testing.T) {

		script := templates.GenerateCreateForwarderScript(
			fungibleAddr,
			forwardingAddr,
			exampleTokenAddr,
			"ExampleToken",
		)

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

		script = templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assertEqual(t, CadenceUFix64("400.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assertEqual(t, CadenceUFix64("600.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assertEqual(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer tokens with the generic transfer transaction", func(t *testing.T) {

		script := templates.GenerateTransferGenericVaultScript(
			fungibleAddr,
		)

		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		_ = tx.AddArgument(CadenceUFix64("300.0"))
		_ = tx.AddArgument(cadence.NewAddress(exampleTokenAddr))

		storagePath := cadence.Path{Domain: "storage", Identifier: "exampleTokenVault"}
		publicPath := cadence.Path{Domain: "public", Identifier: "exampleTokenReceiver"}

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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)
		assertEqual(t, CadenceUFix64("700.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assertEqual(t, CadenceUFix64("300.0"), result)

	})
}

func TestVaultDestroy(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _, metadataViewsAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	script := templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, metadataViewsAddr, "ExampleToken")
	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(
			b.ServiceKey().Address,
			b.ServiceKey().Index,
			b.ServiceKey().SequenceNumber,
		).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(joshAddress)

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

	script = templates.GenerateTransferVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
	tx = flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(
			b.ServiceKey().Address,
			b.ServiceKey().Index,
			b.ServiceKey().SequenceNumber,
		).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(exampleTokenAddr)

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

	t.Run("Should subtract tokens from supply when they are destroyed", func(t *testing.T) {
		script := templates.GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 100)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, exampleTokenAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, exampleTokenAddr},
			[]crypto.Signer{serviceSigner, exampleTokenSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("600.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("900.0"), supply)
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		script := templates.GenerateDestroyVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken", 100)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, joshAddress)

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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("200.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("800.0"), supply)
	})

}

func TestMintingAndBurning(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	serviceSigner, _ := b.ServiceKey().Signer()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _, metadataViewsAddr := DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	script := templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, metadataViewsAddr, "ExampleToken")
	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(
			b.ServiceKey().Address,
			b.ServiceKey().Index,
			b.ServiceKey().SequenceNumber,
		).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(joshAddress)

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

	t.Run("Shouldn't be able to mint zero tokens", func(t *testing.T) {
		script := templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		tx := createTxWithTemplateAndAuthorizer(
			b, script, exampleTokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(CadenceUFix64("0.0"))

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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("0.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		script := templates.GenerateMintTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("50.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1050.0"), supply)
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		script := templates.GenerateBurnTokensScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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
		script = templates.GenerateInspectVaultScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(exampleTokenAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("950.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}
