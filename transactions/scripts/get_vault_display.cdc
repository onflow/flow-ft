import ExampleToken from "../../contracts/ExampleToken.cdc"
import FungibleTokenMetadataViews from "../../contracts/FungibleTokenMetadataViews.cdc"
import MetadataViews from "../../contracts/utilityContracts/MetadataViews.cdc"

pub fun main(address: Address): FungibleTokenMetadataViews.FTDisplay{
  let account = getAccount(address)

  let vaultRef = account
    .getCapability(ExampleToken.ResolverPublicPath)
    .borrow<&{MetadataViews.Resolver}>()
    ?? panic("Could not borrow a reference to the vault resolver")

  let ftDisplay = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
    ?? panic("Token does not implement FTDisplay view")

  return ftDisplay

}