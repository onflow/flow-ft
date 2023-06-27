import ExampleToken from "ExampleToken"
import MetadataViews from "MetadataViews"

pub fun main(address: Address): AnyStruct? {
    let account = getAccount(address)
    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.resolveView(Type<String>())
}
