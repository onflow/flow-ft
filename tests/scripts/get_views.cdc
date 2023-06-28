// This script checks the supported views from ExampleToken
// are the expected ones. This is merely used in testing,
// since we cannot return on-chain types to the test
// files yet.

import MetadataViews from "MetadataViews"
import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    let views = vaultRef.getViews()
    let expected: [Type] = [
        Type<FungibleTokenMetadataViews.FTView>(),
        Type<FungibleTokenMetadataViews.FTDisplay>(),
        Type<FungibleTokenMetadataViews.FTVaultData>()
    ]

    assert(expected == views)

    return true
}
