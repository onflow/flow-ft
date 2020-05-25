package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"
)

const (
	fungibleTokenContractFile = "../src/contracts/FungibleToken.cdc"
	flowTokenContractFile     = "../src/contracts/FlowToken.cdc"
)

func TestTokenDeployment(t *testing.T) {
	b := newEmulator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)
	flowAddr, err := b.CreateAccount(nil, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestCreateToken(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)
	flowAddr, err := b.CreateAccount(nil, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, joshAddress},
			[]crypto.Signer{b.RootKey().Signer(), joshSigner},
			false,
		)

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestExternalTransfers(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
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

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))
	})

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 30000)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateTransferVaultScript(fungibleAddr, flowAddr, joshAddress, 300)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 700))

		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 300))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestVaultDestroy(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
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

	signAndSubmit(
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

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 600))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 900))
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateDestroyVaultScript(fungibleAddr, flowAddr, 100)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, joshAddress},
			[]crypto.Signer{b.RootKey().Signer(), joshSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 200))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 800))
	})

}

func TestMintingAndBurning(t *testing.T) {
	b := newEmulator()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	fungibleAddr, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)
	flowAccountKey, flowSigner := accountKeys.NewWithSigner()
	flowAddr, err := b.CreateAccount([]*flow.AccountKey{flowAccountKey}, flowTokenCode)
	assert.NoError(t, err)
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// then deploy the tokens to an account
	tx := flow.NewTransaction().
		SetScript(GenerateCreateTokenScript(fungibleAddr, flowAddr)).
		SetGasLimit(100).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
		SetPayer(b.RootKey().Address).
		AddAuthorizer(joshAddress)

	signAndSubmit(
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

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})

	t.Run("Shouldn't be able to mint more than the allowed amount", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, flowAddr, joshAddress, 101)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			true,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 0))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateMintTokensScript(fungibleAddr, flowAddr, joshAddress, 50)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 1000))

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, joshAddress, 50))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1050))
	})

	t.Run("Should burn tokens and update balance and total supply", func(t *testing.T) {
		tx := flow.NewTransaction().
			SetScript(GenerateBurnTokensScript(fungibleAddr, flowAddr, 50)).
			SetGasLimit(100).
			SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber).
			SetPayer(b.RootKey().Address).
			AddAuthorizer(flowAddr)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.RootKey().Address, flowAddr},
			[]crypto.Signer{b.RootKey().Signer(), flowSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		executeScriptAndCheck(t, b, GenerateInspectVaultScript(fungibleAddr, flowAddr, flowAddr, 950))

		executeScriptAndCheck(t, b, GenerateInspectSupplyScript(fungibleAddr, flowAddr, 1000))
	})
}

func TestTokenAdministrator(t *testing.T) {
	b := newEmulator()

	// Should be able to deploy a contract as a new account with no keys.
	fungibleTokenCode := readFile(fungibleTokenContractFile)
	_, err := b.CreateAccount(nil, fungibleTokenCode)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	flowTokenCode := readFile(flowTokenContractFile)

	setupScript := fmt.Sprintf(`
      transaction {
        prepare(signer: AuthAccount) {
          let flowTokenAcct = AuthAccount(payer: signer)

          let admin = signer
          flowTokenAcct.setCode("%s".decodeHex(), admin)
        }
      }
    `, hex.EncodeToString(flowTokenCode))

	setupTx := flow.NewTransaction().
		SetScript([]byte(setupScript)).
		AddAuthorizer(b.RootKey().Address).
		SetPayer(b.RootKey().Address).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber)

	err = setupTx.SignEnvelope(b.RootKey().Address, b.RootKey().ID, b.RootKey().Signer())
	require.NoError(t, err)

	err = b.AddTransaction(*setupTx)
	require.NoError(t, err)

	result, err := b.ExecuteNextTransaction()
	require.NoError(t, err)

	if !assert.Nil(t, result.Error) {
		t.Fatal(result.Error)
	}

	_, err = b.CommitBlock()
	require.NoError(t, err)

	mintScript := `
      import FungibleToken from 0x02
      import FlowToken from 0x06

      transaction {
        let tokenAdmin: &FlowToken.Administrator
        let tokenReceiver: &FlowToken.Vault{FungibleToken.Receiver}

        prepare(signer: AuthAccount) {
          self.tokenAdmin = signer.
            borrow<&FlowToken.Administrator>(from: /storage/flowTokenAdmin) 
            ?? panic("Signer is not the token admin")

          self.tokenReceiver = signer
            .getCapability(/public/flowTokenReceiver)!
            .borrow<&FlowToken.Vault{FungibleToken.Receiver}>()
            ?? panic("Unable to borrow receiver reference for recipient")
        }

        execute {
          let minter <- self.tokenAdmin.createNewMinter(allowedAmount: 100.0)
          let mintedVault <- minter.mintTokens(amount: 100.0)

          self.tokenReceiver.deposit(from: <-mintedVault)

          log("Minted 100 tokens and deposited to admin account")

          destroy minter
        }
      }
    `

	mintTx := flow.NewTransaction().
		SetScript([]byte(mintScript)).
		AddAuthorizer(b.RootKey().Address).
		SetPayer(b.RootKey().Address).
		SetProposalKey(b.RootKey().Address, b.RootKey().ID, b.RootKey().SequenceNumber)

	err = mintTx.SignEnvelope(b.RootKey().Address, b.RootKey().ID, b.RootKey().Signer())
	require.NoError(t, err)

	err = b.AddTransaction(*mintTx)
	require.NoError(t, err)

	result, err = b.ExecuteNextTransaction()
	require.NoError(t, err)

	if !assert.Nil(t, result.Error) {
		t.Fatal(result.Error)
	}

	_, err = b.CommitBlock()
	require.NoError(t, err)
}
