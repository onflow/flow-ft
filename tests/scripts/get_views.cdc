// This script checks the supported views from ExampleToken
// are the expected ones. This is merely used in testing.

import "MetadataViews"
import "ExampleToken"
import "FungibleTokenMetadataViews"

pub fun main(address: Address): [Type] {
    let account = getAccount(address)

    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.getViews()
}
