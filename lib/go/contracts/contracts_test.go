package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-ft/lib/go/contracts"
)

const addrA = "0A"

func TestFungibleTokenContract(t *testing.T) {
	contract := contracts.FungibleToken()
	assert.NotNil(t, contract)
}

func TestExampleTokenContract(t *testing.T) {
	contract := contracts.ExampleToken(addrA, addrA, addrA)
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}

func TestCustomExampleTokenContract(t *testing.T) {
	contract := contracts.CustomToken(addrA, addrA, addrA, "UtilityCoin", "utilityCoin", "100.0")
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}

func TestTokenForwardingContract(t *testing.T) {
	contract := contracts.TokenForwarding(addrA)
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}

func TestCustomTokenForwardingContract(t *testing.T) {
	contract := contracts.CustomTokenForwarding(addrA, "UtilityCoin", "utilityCoin")
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}
