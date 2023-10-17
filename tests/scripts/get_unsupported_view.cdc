// This script checks the resolveView from ExampleToken
// returns nil for unsupported view. This is merely used
// in testing, since we cannot return on-chain types to
// the test files yet.

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

pub fun main(address: Address): Bool {
    let account = getAccount(address)
    let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow Balance reference to the Vault")

    assert(nil == vaultRef.resolveView(Type<String>()))

    return true
}
