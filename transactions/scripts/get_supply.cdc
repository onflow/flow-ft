// This script reads the total supply field
// of the ExampleToken smart contract

import "ExampleToken"

pub fun main(): UFix64 {

    let supply = ExampleToken.totalSupply

    log(supply)

    return supply
}