// This script checks the FTDisplay view from ExampleToken
// is the expected one. This is merely used in testing.

import "ExampleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"
import "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTDisplay {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")
    
    let vaultRef = account.capabilities.borrow<&ExampleToken.Vault>(vaultData.metadataPath)
        ?? panic("Could not borrow Balance reference to the Vault")

    let ftDisplay = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
        ?? panic("Token does not implement FTDisplay view")

    return ftDisplay
}
