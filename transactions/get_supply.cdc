// This script reads the total supply field
// of the FlowToken smart contract

import FlowToken from 0x02

pub fun main(): UFix64 {

    let supply = FlowToken.totalSupply

    log(supply)

    return supply
}