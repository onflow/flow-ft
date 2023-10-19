// This script checks the FTView view from ExampleToken
// is the expected one. This is merely used in testing.

import "ExampleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account
        .getCapability(ExampleToken.VaultPublicPath)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

    // FungibleTokenMetadataViews.FTVaultData cannot be returned as
    // a script result, because of the createEmptyVault() function.
    // So we perform the assertions here.
    return ftView.ftDisplay != nil && ftView.ftVaultData != nil
}
