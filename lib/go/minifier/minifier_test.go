package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const expectedOutput = `import FungibleToken from 0xFUNGIBLETOKENADDRESS
import ExampleToken from 0xTOKENADDRESS
transaction(amount: UFix64, to: Address) {
let sentVault: @FungibleToken.Vault
prepare(signer: AuthAccount) {
let vaultRef = signer.borrow<&ExampleToken.Vault>(from: /storage/exampleTokenVault)
?? panic("Could not borrow reference to the owner's Vault!")
self.sentVault <- vaultRef.withdraw(amount: amount)
}
execute {
let recipient = getAccount(to)
let receiverRef = recipient.getCapability(/public/exampleTokenReceiver)!.borrow<&{FungibleToken.Receiver}>()
?? panic("Could not borrow receiver reference to the recipient's Vault")
receiverRef.deposit(from: <-self.sentVault)
}
}
`
// Test to test the minifier function
func TestMinify(t *testing.T) {
	inputFileName := "../../../transactions/transfer_tokens.cdc"
	outputFile, err := ioutil.TempFile("", "minified*.cdc")
	require.NoError(t, err)
	outputFileName := outputFile.Name()
	defer os.Remove(outputFileName)
	err = minify(inputFileName, outputFileName)
	require.NoError(t, err)
	actualOutput, err := ioutil.ReadFile(outputFileName)
	require.NoError(t, err)
	assert.Equal(t, expectedOutput, string(actualOutput))
}
