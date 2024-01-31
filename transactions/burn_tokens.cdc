import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import Burner from "Burner"

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
            ?? panic("Could not get vault data view for the contract")

        // Withdraw tokens from the signer's vault in storage
        let sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(
                from: vaultData.storagePath
            ) ?? panic("Could not borrow a reference to the signer's ExampleToken vault")
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
