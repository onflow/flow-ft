import "FungibleToken"
import "FungibleTokenSwitchboard"
import "ExampleToken"
import "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to send tokens to another account
/// through a switchboard using the deposit method but before depositing we will explicitly check whether receiving
/// capability is borrowable or not and if yes then it will deposit the vault to the receiver capability.
///
transaction(to: Address, amount: UFix64) {

    // The reference to the vault from the payer's account
    let vaultRef: auth(FungibleToken.Withdraw) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
			?? panic("The signer does not store a ExampleToken Vault object at the path "
                .concat(vaultData.storagePath.toString())
                .concat(". The signer must initialize their account with this object first!"))

    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        let sentVault <- self.vaultRef.withdraw(amount: amount)

        // Get a reference to the recipient's SwitchboardPublic
        let switchboardRef = recipient.capabilities.borrow<&{FungibleTokenSwitchboard.SwitchboardPublic}>(
                FungibleTokenSwitchboard.PublicPath)
			?? panic("The signer does not store a FungibleToken Switchboard capability at the path "
                .concat(FungibleTokenSwitchboard.PublicPath.toString())
                .concat(". The signer must initialize their account with this object first!"))  

        // Validate the receiving capability by using safeBorrowByType
        if let receivingRef = switchboardRef.safeBorrowByType(type: Type<@ExampleToken.Vault>()) {
            switchboardRef.deposit(from: <-sentVault)
        } else {
            // Return funds to signer's account if receiver is not configured to receive the funds
            self.vaultRef.deposit(from: <-sentVault)
        }
    }

}
