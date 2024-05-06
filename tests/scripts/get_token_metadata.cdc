// This script checks the FTView view from ExampleToken
// is the expected one. This is merely used in testing.

import "ExampleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"
import "ViewResolver"

access(all) fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get the vault data view for ExampleToken")

    let vaultRef = account.capabilities.borrow<&{ViewResolver.Resolver}>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

    // FungibleTokenMetadataViews.FTVaultData cannot be returned as
    // a script result, because of the createEmptyVault() function.
    // So we perform the assertions here.
    return ftView.ftDisplay != nil && ftView.ftVaultData != nil
}
