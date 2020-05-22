package contracts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/onflow/flow-go-sdk"
)

const (
	RepoBaseURL                   = "https://raw.githubusercontent.com/onflow/flow-ft/master/contracts/"
	FungibleTokenContractFilename = "FungibleToken.cdc"
	FlowTokenContractFilename     = "FlowToken.cdc"
)

var (
	fungibleTokenCode []byte
	flowTokenTemplate []byte
)

// Load the contracts when the module is loaded.
func init() {
	fungibleTokenCode, _ = getContract(FungibleTokenContractFilename)
	flowTokenTemplate, _ = getContract(FlowTokenContractFilename)
}

func getContract(path string) ([]byte, error) {
	resp, err := http.Get(RepoBaseURL + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// FungibleTokenContract returns the FungibleToken contract interface.
func FungibleTokenContract() ([]byte, error) {
	if len(fungibleTokenCode) == 0 {
		return nil, fmt.Errorf("fungible token code not available")
	}
	return fungibleTokenCode, nil
}

// FlowTokenContract returns the FlowToken contract, importing the
// FungibleToken contract interface from the specified Flow address.
func FlowTokenContract(fungibleTokenAddress flow.Address) ([]byte, error) {
	// Replace hard-coded address of FungibleToken contract interface.
	if len(flowTokenTemplate) == 0 {
		return nil, fmt.Errorf("flow token code not available")
	}
	code := strings.ReplaceAll(
		string(flowTokenTemplate),
		"0x02",
		"0x"+fungibleTokenAddress.Hex(),
	)
	return []byte(code), nil
}
