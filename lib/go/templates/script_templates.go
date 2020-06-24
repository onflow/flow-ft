package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../scripts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../scripts

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
)

const (
	readDataFilename = "get_balance.cdc"
)

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(fungibleAddr, tokenAddr, userAddr flow.Address, tokenName string, expectedBalance float64) []byte {
	storageName := MakeFirstLowerCase(tokenName)

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s

		pub fun main() {
			let acct = getAccount(0x%[5]s)
			let vaultRef = acct.getCapability(/public/%[4]sBalance)!.borrow<&%[3]s.Vault{FungibleToken.Balance}>()
				?? panic("Could not borrow Balance reference to the Vault")
			assert(
                vaultRef.balance == UFix64(%[6]f),
                message: "incorrect balance!"
            )
		}
    `

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, storageName, userAddr, expectedBalance))
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, tokenAddr flow.Address, tokenName string, expectedSupply int) []byte {

	template := `
		import FungibleToken from 0x%[1]s 
		import %[3]s from 0x%[2]s

		pub fun main() {
			assert(
                %[3]s.totalSupply == UFix64(%[4]d),
                message: "incorrect totalSupply!"
            )
		}
	`

	return []byte(fmt.Sprintf(template, fungibleAddr, tokenAddr, tokenName, expectedSupply))
}
