/// This script reads the balance of a vault at whatever path
/// is passed in as an argument

import "FungibleToken"

access(all) fun main(address: Address, path: PublicPath): UFix64 {
    return getAccount(address).capabilities.borrow<&{FungibleToken.Balance}>(
            path
        )?.balance
        ?? panic("Could not borrow a balance reference to the FungibleToken Vault in account \(address) at path \(path). Make sure you are querying an address that has a FungibleToken Vault set up properly at the specified path.")
}
