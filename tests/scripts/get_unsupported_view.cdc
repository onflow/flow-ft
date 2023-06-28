// This script checks the resolveView from ExampleToken
// returns nil for unsupported view. This is merely used
// in testing, since we cannot return on-chain types to
// the test files yet.

import ExampleToken from "ExampleToken"
import MetadataViews from "MetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)
    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    assert(nil == vaultRef.resolveView(Type<String>()))

    return true
}
