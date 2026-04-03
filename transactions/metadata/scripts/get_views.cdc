// This script checks the supported views from ExampleToken
// are the expected ones. This is merely used in testing.

import "MetadataViews"
import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"

access(all) fun main(address: Address): [Type] {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not resolve `FTVaultData` view. The ExampleToken contract needs to implement the `FTVaultData` Metadata view in order to execute this script.")

    let vaultRef = account.capabilities.borrow<&ExampleToken.Vault>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the `ExampleToken.Vault` in account \(address) at path \(vaultData.metadataPath). Make sure you are querying an address that has an `ExampleToken.Vault` set up properly.")

    return vaultRef.getViews()
}
