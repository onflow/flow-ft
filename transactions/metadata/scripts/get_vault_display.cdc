import "ExampleToken"
import "FungibleTokenMetadataViews"
import "FungibleToken"
import "ViewResolver"

access(all) fun main(address: Address): FungibleTokenMetadataViews.FTDisplay {
    let account = getAccount(address)

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not resolve FTVaultData view. The ExampleToken"
            .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction."))

    let vaultRef = account.capabilities.borrow<&ExampleToken.Vault>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the ExampleToken Vault in account "
                .concat(address.toString()).concat(" at path ").concat(vaultData.metadataPath.toString())
                .concat(". Make sure you are querying an address that has an ExampleToken Vault set up properly."))

    let ftDisplay = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
        ?? panic("Token does not implement FTDisplay view")

    return ftDisplay
}
