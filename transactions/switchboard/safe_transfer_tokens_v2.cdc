import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to send tokens to another account
/// through a switchboard using the deposit method but before depositing we will explicitly check whether receiving
/// capability is borrowable or not and if yes then it will deposit the vault to the receiver capability.
///
transaction(to: Address, amount: UFix64) {

    // The reference to the vault from the payer's account
    let vaultRef: auth(FungibleToken.Withdraw) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())
            ?? panic("Could not get vault data view for the contract")

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
			?? panic("Could not borrow reference to the owner's Vault!")

    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        let sentVault <- self.vaultRef.withdraw(amount: amount)

        // Get a reference to the recipient's SwitchboardPublic
        let switchboardRef = recipient.capabilities.borrow<&{FungibleTokenSwitchboard.SwitchboardPublic}>(
                FungibleTokenSwitchboard.PublicPath
            ) ?? panic("Could not borrow receiver reference to switchboard!")    

        // Validate the receiving capability by using safeBorrowByType
        if let receivingRef = switchboardRef.safeBorrowByType(type: Type<@ExampleToken.Vault>()) {
            switchboardRef.deposit(from: <-sentVault)
        } else {
            // Return funds to signer's account if receiver is not configured to receive the funds
            self.vaultRef.deposit(from: <-sentVault)
        }
    }

}
