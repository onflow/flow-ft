package contracts_test

import (
	"testing"

	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-ft/contracts"
)

var addrA = flow.HexToAddress("0A")

func TestFungibleTokenContract(t *testing.T) {
	contract := contracts.FungibleToken()
	assert.NotNil(t, contract)
}

func TestFlowTokenContract(t *testing.T) {
	contract := contracts.FlowToken(addrA.Hex())
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}

func TestExampleTokenContract(t *testing.T) {
	contract := contracts.ExampleToken(addrA.Hex())
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}
