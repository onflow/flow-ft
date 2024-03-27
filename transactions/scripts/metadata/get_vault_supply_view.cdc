import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"
import "ViewResolver"

/// Gets the total supply of the vault's token directly from the vault

access(all) fun main(address: Address): UFix64 {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")

    let vaultRef = account.capabilities.borrow<&ExampleToken.Vault>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftSupply = vaultRef.resolveView(Type<FungibleTokenMetadataViews.TotalSupply>())!

    let supplyView = ftSupply as! FungibleTokenMetadataViews.TotalSupply

    return supplyView.supply
}
