// This script reads the balance field
// of an account's ExampleToken Balance

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

access(all) fun main(address: Address): UFix64 {
    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())

    return getAccount(address).capabilities.borrow<&{FungibleToken.Vault}>(
            vaultData.metadataPath
        )?.balance
        ?? panic("Could not borrow Balance reference to the Vault")
}
