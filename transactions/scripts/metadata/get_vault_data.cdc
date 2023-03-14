import ExampleToken from "../../../contracts/ExampleToken.cdc"
import FungibleTokenMetadataViews from "../../../contracts/FungibleTokenMetadataViews.cdc"
import MetadataViews from "../../../contracts/utility/MetadataViews.cdc"

pub fun main(address: Address): FungibleTokenMetadataViews.FTVaultData{
  let account = getAccount(address)

  let vaultRef = account
    .getCapability(ExampleToken.VaultPublicPath)
    .borrow<&{MetadataViews.Resolver}>()
    ?? panic("Could not borrow a reference to the vault resolver")

  let vaultData = FungibleTokenMetadataViews.getFTVaultData(vaultRef)
    ?? panic("Token does not implement FTVaultData view")

  return vaultData

}