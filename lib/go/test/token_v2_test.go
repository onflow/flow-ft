package test

import (
	"testing"

	// sdktemplates "github.com/onflow/flow-go-sdk/templates"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"

	// "github.com/onflow/cadence"
	// jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	// "github.com/onflow/flow-go-sdk/crypto"
	// "github.com/onflow/flow-ft/lib/go/contracts"
	// "github.com/onflow/flow-ft/lib/go/templates"
)

func TestV2TokenDeployment(t *testing.T) {
	b, accountKeys := newTestSetup(t)

	exampleTokenAccountKey, _ := accountKeys.NewWithSigner()
	_ = DeployV2TokenContracts(b, t, []*flow.AccountKey{exampleTokenAccountKey})

	// t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
	// 	script := templates.GenerateInspectSupplyScript(fungibleAddr, exampleTokenAddr, "ExampleToken")
	// 	supply := executeScriptAndCheck(t, b, script, nil)
	// 	assert.Equal(t, CadenceUFix64("1000.0"), supply)
	// })
}
