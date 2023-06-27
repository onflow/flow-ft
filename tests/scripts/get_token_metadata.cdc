import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"

pub fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account
        .getCapability(ExampleToken.VaultPublicPath)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

    return ftView.ftDisplay != nil && ftView.ftVaultData != nil
}
