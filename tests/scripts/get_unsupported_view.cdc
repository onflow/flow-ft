// This script checks the resolveView from ExampleToken
// returns nil for unsupported view. This is merely used
// in testing.

import "ExampleToken"
import "MetadataViews"

pub fun main(address: Address, type: Type): AnyStruct? {
    let account = getAccount(address)
    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.resolveView(type)
}
