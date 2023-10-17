// This script reads the balance field
// of an account's ExampleToken Balance

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

access(all) fun main(address: Address): UFix64 {
    return getAccount(address).capabilities.borrow<&{FungibleToken.Vault}>(
            ExampleToken.VaultPublicPath
        )?.getBalance()
        ?? panic("Could not borrow Balance reference to the Vault")
}
