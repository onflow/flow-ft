/// This script reads the balance of a vault at whatever path
/// is passed in as an argument

import "FungibleToken"

access(all) fun main(address: Address, path: PublicPath): UFix64 {
    return getAccount(address).capabilities.borrow<&{FungibleToken.Balance}>(
            path
        )?.balance
        ?? panic("Could not borrow a balance reference to the FungibleToken Vault in account "
                .concat(address.toString()).concat(" at path ").concat(path.toString())
                .concat(". Make sure you are querying an address that has ")
                .concat("a FungibleToken Vault set up properly at the specified path."))
}
