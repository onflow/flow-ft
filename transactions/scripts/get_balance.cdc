// This script reads the balance field
// of an account's ExampleToken Balance

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

access(all) fun main(address: Address): UFix64 {
    let account = getAccount(address)
    let vaultRef = account.getCapability<&{FungibleToken.Balance}>(ExampleToken.VaultPublicPath)
        .borrow()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.getBalance()
}
