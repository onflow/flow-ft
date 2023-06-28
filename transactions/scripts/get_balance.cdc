// This script reads the balance field
// of an account's ExampleToken Balance

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

pub fun main(address: Address): UFix64 {
    let account = getAccount(address)
    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{FungibleToken.Balance}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.balance
}
