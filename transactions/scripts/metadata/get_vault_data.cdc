import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import FungibleToken from "FungibleToken"
import ViewResolver from "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTVaultData {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")

    let vaultRef = account.capabilities.borrow<&{FungibleToken.Balance}>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let vaultData = FungibleTokenMetadataViews.getFTVaultData(vaultRef)
        ?? panic("Token does not implement FTVaultData view")

    return vaultData
}
