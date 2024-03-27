// This script checks the FTVaultData view from ExampleToken
// is the expected one. This is merely used in testing.

import "ExampleToken"
import "FungibleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"

access(all) fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")

    // FungibleTokenMetadataViews.FTVaultData cannot be returned as
    // a script result, because of the createEmptyVault() function.
    // So we perform the assertions here.
    assert(Type<&ExampleToken.Vault>() == vaultData.receiverLinkedType)
    assert(Type<&ExampleToken.Vault>() == vaultData.metadataLinkedType)
    let vault <- vaultData.createEmptyVault()
    let vaultIsEmpty = vault.balance == 0.0
    assert(vaultIsEmpty)

    destroy vault

    return true
}
