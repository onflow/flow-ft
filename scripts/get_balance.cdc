// This script reads the balance field of an account's ExampleToken Balance

import "FungibleToken"
import "ExampleToken"

pub fun main(account: Address): UFix64 {
    let acct = getAccount(account)
    let vaultRef = acct.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{FungibleToken.Balance}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.balance
}
