import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"
import "Burner"

/// This transaction is a template for a transaction that could be used by an account
/// to load a vault from storage and burn the whole vault
///
/// It is meant for testing purposes to burn an optional vault with the burner contract
///
transaction {

    /// The total supply of tokens before the burn
    let supplyBefore: UFix64

    /// Vault resource that holds the tokens that are being burned
    let burnVault: @AnyResource

    prepare(signer: auth(Storage) &Account) {

        self.supplyBefore = ExampleToken.totalSupply

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not get vault data view for the contract")

        // Withdraw tokens from the signer's vault in storage
        self.burnVault <- signer.storage.load<@AnyResource>(
                from: vaultData.storagePath
        )
    }

    execute {

        Burner.burn(<-self.burnVault)

    }
}
