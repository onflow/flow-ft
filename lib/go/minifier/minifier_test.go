package main

import (
	"fmt"
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
	inputFile := "../../../transactions/transfer_tokens.cdc"
	outputFile, err := ioutil.TempFile("/tmp", "minified*.cdc")
	require.NoError(t, err)
	fmt.Println(outputFile.Name())
	defer os.Remove(outputFile.Name())
	err = minify(inputFile, outputFile.Name())
	require.NoError(t, err)
	output, err := ioutil.ReadFile(outputFile.Name())
	require.NoError(t, err)
	assert.Equal(t, expectedOutput, string(output))
}
