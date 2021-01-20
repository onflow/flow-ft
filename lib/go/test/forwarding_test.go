package test

import (
	"testing"

	"github.com/onflow/flow-emulator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-ft/lib/go/contracts"
	"github.com/onflow/flow-ft/lib/go/templates"
)

func TestPrivateForwarder(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	exampleTokenAccountKey, exampleTokenSigner := accountKeys.NewWithSigner()
	fungibleAddr, exampleTokenAddr, _ :=
		DeployTokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	forwardingCode := contracts.PrivateReceiverForwarder(fungibleAddr.String())

	forwardingAccountKey, forwardingSigner := accountKeys.NewWithSigner()
	forwardingAddr, err := b.CreateAccount(
		[]*flow.AccountKey{forwardingAccountKey},
		[]sdktemplates.Contract{
			{
				Name:   "PrivateReceiverForwarder",
				Source: string(forwardingCode),
			},
		},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	script := templates.GenerateCreateTokenScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
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
			b.ServiceKey().Signer(),
			joshSigner,
		},
		false,
	)

	t.Run("Should be able to transfer tokens through a forwarder from a vault", func(t *testing.T) {

		script := templates.GenerateCreateForwarderScript(
			fungibleAddr,
			forwardingAddr,
			exampleTokenAddr,
			"ExampleToken",
		)

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

		_ = tx.AddArgument(cadence.NewAddress(exampleTokenAddr))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
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
				b.ServiceKey().Signer(),
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
		assert.Equal(t, CadenceUFix64("700.0"), balance)

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
		assert.Equal(t, CadenceUFix64("300.0"), balance)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
		supply := executeScriptAndCheck(t, b, script)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}
