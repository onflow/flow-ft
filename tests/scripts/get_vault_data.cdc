// This script checks the FTVaultData view from ExampleToken
// is the expected one. This is merely used in testing,
// since we cannot return on-chain types to the test
// files yet.

import ExampleToken from "ExampleToken"
import FungibleToken from "FungibleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"

access(all) fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let vaultData = FungibleTokenMetadataViews.getFTVaultData(vaultRef)
        ?? panic("Token does not implement FTVaultData view")

    assert(ExampleToken.VaultStoragePath == vaultData.storagePath)
    assert(ExampleToken.ReceiverPublicPath == vaultData.receiverPath)
    assert(ExampleToken.VaultPublicPath == vaultData.metadataPath)
    assert(/private/exampleTokenVault == vaultData.providerPath)
    assert(Type<&{FungibleToken.Receiver}>() == vaultData.receiverLinkedType)
    assert(Type<&ExampleToken.Vault>() == vaultData.providerLinkedType)
    assert(Type<&ExampleToken.Vault>() == vaultData.metadataLinkedType)
    let vault <- vaultData.createEmptyVault()
    let vaultIsEmpty = vault.getBalance() == 0.0
    assert(vaultIsEmpty)

    destroy vault

    return true
}
