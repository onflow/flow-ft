import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"
import "Burner"

/// This transaction is a template for a transaction that could be used by the admin account to burn tokens from their
/// stored Vault
///
/// The burning amount would be a parameter to the transaction
///
transaction(amount: UFix64) {

    /// The total supply of tokens before the burn
    let supplyBefore: UFix64

    /// Vault resource that holds the tokens that are being burned
    let burnVault: @ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {

        self.supplyBefore = ExampleToken.totalSupply

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Withdraw tokens from the signer's vault in storage
        let sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(
                from: vaultData.storagePath)
			?? panic("The signer does not store a ExampleToken Vault object at the path "
                .concat(vaultData.storagePath.toString())
                .concat("The signer must initialize their account with this object first!"))
                
        self.burnVault <- sourceVault.withdraw(amount: amount) as! @ExampleToken.Vault
    }

    execute {

        Burner.burn(<-self.burnVault)

    }

    post {
        ExampleToken.totalSupply == (self.supplyBefore - amount):
            "Before: ".concat(self.supplyBefore.toString())
            .concat(" | After: ".concat(ExampleToken.totalSupply.toString()))
            .concat(" | Expected: ".concat((self.supplyBefore - amount).toString()))
    }
}
