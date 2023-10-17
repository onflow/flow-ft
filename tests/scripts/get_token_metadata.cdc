// This script checks the FTView view from ExampleToken
// is the expected one. This is merely used in testing,
// since we cannot return on-chain types to the test
// files yet.

import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{MetadataViews.Resolver}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

    return ftView.ftDisplay != nil && ftView.ftVaultData != nil
}
