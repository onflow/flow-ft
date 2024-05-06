import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"
import "Burner"

/// This transaction is for testing burning an array of Vaults
///
transaction(amountPerIndex: UFix64, numIndicies: Int) {

    /// Vault resource that holds the tokens that are being burned
    let burnVaults: @[ExampleToken.Vault]

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not get vault data view for the contract")

        // Withdraw tokens from the signer's vault in storage
        let sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(
                from: vaultData.storagePath
            ) ?? panic("Could not borrow a reference to the signer's ExampleToken vault")
        
        self.burnVaults <- []
        var i = 0
        while i < numIndicies {
            let vault <- sourceVault.withdraw(amount: amountPerIndex) as! @ExampleToken.Vault
            self.burnVaults.append(<-vault)
            i = i + 1
        }
    }

    execute {

        Burner.burn(<-self.burnVaults)

    }
}
