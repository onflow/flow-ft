// This script reads the balance field
// of an account's ExampleToken Balance

import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"

access(all) fun main(address: Address): UFix64 {
    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not resolve `FTVaultData` view. The ExampleToken contract needs to implement the `FTVaultData` Metadata view in order to execute this script.")

    return getAccount(address).capabilities.borrow<&{FungibleToken.Balance}>(
            vaultData.metadataPath
        )?.balance
        ?? panic("Could not borrow a balance reference to the `FungibleToken.Vault` in account \(address) at path \(vaultData.metadataPath). Make sure you are querying an address that has a `FungibleToken.Vault` set up properly at the specified path.")
}
