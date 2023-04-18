import "ExampleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"

pub fun main(address: Address): FungibleTokenMetadataViews.FTView{
  let account = getAccount(address)

  let vaultRef = account
    .getCapability(ExampleToken.VaultPublicPath)
    .borrow<&{MetadataViews.Resolver}>()
    ?? panic("Could not borrow a reference to the vault resolver")

  let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)

  return ftView

}