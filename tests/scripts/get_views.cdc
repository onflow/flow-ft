// This script checks the supported views from ExampleToken
// are the expected ones. This is merely used in testing.

import "MetadataViews"
import "ExampleToken"
import "FungibleTokenMetadataViews"

pub fun main(address: Address): [Type] {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow reference to the Vault")

    return vaultRef.getViews()
}
