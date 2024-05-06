/// This script reads the balance of a vault at whatever path
/// is passed in as an argument

import "FungibleToken"

access(all) fun main(address: Address, path: PublicPath): UFix64 {
    return getAccount(address).capabilities.borrow<&{FungibleToken.Balance}>(
            path
        )?.balance
        ?? panic("Could not borrow Balance reference to the Vault")
}
