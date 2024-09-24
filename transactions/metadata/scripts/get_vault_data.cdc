import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"
import "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTVaultData {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not resolve FTVaultData view. The ExampleToken"
            .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction."))

    return vaultData
}
