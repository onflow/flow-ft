import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account
        .getCapability(ExampleToken.VaultPublicPath)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the vault resolver")

    let vaultData = FungibleTokenMetadataViews.getFTVaultData(vaultRef)
        ?? panic("Token does not implement FTVaultData view")

    let vault <- vaultData.createEmptyVault()

    let vaultIsEmpty = vault.balance == 0.0

    destroy vault

    return vaultIsEmpty
}
