// This script checks the resolveView from ExampleToken
// returns nil for unsupported view. This is merely used
// in testing.

import "ExampleToken"
import "MetadataViews"
import "FungibleToken"

access(all) fun main(address: Address, type: Type): AnyStruct? {
    let account = getAccount(address)
    let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.resolveView(type)
}
