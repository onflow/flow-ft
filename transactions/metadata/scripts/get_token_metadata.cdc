import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"
import "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTView {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")

    let vaultRef = account.capabilities.borrow<&ExampleToken.Vault>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

    return ftView
}
