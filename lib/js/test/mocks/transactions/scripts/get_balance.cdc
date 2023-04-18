// This script reads the balance field of an account's ExampleToken Balance

import "Token"
import "TestToken"

pub fun main(account: Address): UFix64 {
    let acct = getAccount(account)
    let vaultRef = acct.getCapability<&{Token.Balance}>(TestToken.VaultPublicPath).borrow()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.balance
}