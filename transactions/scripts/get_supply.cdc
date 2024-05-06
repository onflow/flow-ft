// This script reads the total supply field
// of the ExampleToken smart contract

import "ExampleToken"

access(all) fun main(): UFix64 {
    return ExampleToken.totalSupply
}
