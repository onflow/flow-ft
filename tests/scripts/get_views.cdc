// This script checks the supported views from ExampleToken
// are the expected ones. This is merely used in testing.

import "MetadataViews"
import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"

access(all) fun main(address: Address): [Type] {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())
    
    let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(vaultData.metadataPath)
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.getViews()
}
