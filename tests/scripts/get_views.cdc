import MetadataViews from "MetadataViews"
import ExampleToken from "ExampleToken"

pub fun main(address: Address): [String] {
    let account = getAccount(address)

    let vaultRef = account.getCapability(ExampleToken.VaultPublicPath)
        .borrow<&ExampleToken.Vault{MetadataViews.Resolver}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    let views = vaultRef.getViews()
    let viewIDs: [String] = []
    for view in views {
        viewIDs.append(view.identifier)
    }

    return viewIDs
}
