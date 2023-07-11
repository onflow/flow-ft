import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"

/// Gets the total supply of the vault's token directly from the vault

pub fun main(address: Address): UFix64 {
    let account = getAccount(address)

    let vaultRef = account
        .getCapability(ExampleToken.VaultPublicPath)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftSupply = vaultRef.resolveView(Type<FungibleTokenMetadataViews.TotalSupply>())!

    let supplyView = ftSupply as! FungibleTokenMetadataViews.TotalSupply

    return supplyView.supply
}
