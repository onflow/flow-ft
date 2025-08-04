import "FungibleToken"
import "MetadataViews"
import "FungibleTokenMetadataViews"
import "Burner"

/// This transaction is a template for a transaction that could be used
/// by any account to burn tokens from their stored Vault for any Fungible Token
///
/// @param ftTypeIdentifier: The type identifier name of the FT type to burn
/// Ex: "A.1654653399040a61.FlowToken.Vault"
/// @param amount: The amount of tokens to burn
///
transaction(ftTypeIdentifier: String, amount: UFix64) {

    /// Vault resource that holds the tokens that are being burned
    let burnVault: @{FungibleToken.Vault}

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = MetadataViews.resolveContractViewFromTypeIdentifier(
            resourceTypeIdentifier: ftTypeIdentifier,
            viewType: Type<FungibleTokenMetadataViews.FTVaultData>()
        ) as? FungibleTokenMetadataViews.FTVaultData
            ?? panic("Could not construct valid FT type and view from identifier \(ftTypeIdentifier)")

        // Withdraw tokens from the signer's vault in storage
        let sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &{FungibleToken.Vault}>(
                from: vaultData.storagePath)
			?? panic("The signer does not store a FungibleToken Vault object at the path \(vaultData.storagePath.toString())"
                .concat(". The signer must initialize their account with this object first!"))

        self.burnVault <- sourceVault.withdraw(amount: amount)
    }

    execute {

        Burner.burn(<-self.burnVault)

    }
}
