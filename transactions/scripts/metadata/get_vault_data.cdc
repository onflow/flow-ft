import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import ViewResolver from "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTVaultData {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{ViewResolver.Resolver}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let vaultData = FungibleTokenMetadataViews.getFTVaultData(vaultRef)
        ?? panic("Token does not implement FTVaultData view")

    return vaultData
}
