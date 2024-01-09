// This script checks the FTDisplay view from ExampleToken
// is the expected one. This is merely used in testing.

import "ExampleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"
import "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTDisplay {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{ViewResolver.Resolver}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftDisplay = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
        ?? panic("Token does not implement FTDisplay view")

    return ftDisplay
}
