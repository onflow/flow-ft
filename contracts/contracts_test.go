package contracts_test

import (
	"testing"

	"github.com/onflow/flow-ft/contracts"
	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestFungibleTokenContract(t *testing.T) {
	contract, err := contracts.FungibleToken()
	assert.NotNil(t, contract)
	assert.NoError(t, err)
}

func TestFlowTokenContract(t *testing.T) {
	contract, err := contracts.FlowToken(flow.Address{0x3})
	assert.NotNil(t, contract)
	assert.NoError(t, err)
	assert.Contains(t, string(contract), "0x03")
}
