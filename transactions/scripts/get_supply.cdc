// This script reads the total supply field
// of the ExampleToken smart contract

import ExampleToken from "ExampleToken"

pub fun main(): UFix64 {
    return ExampleToken.totalSupply
}
