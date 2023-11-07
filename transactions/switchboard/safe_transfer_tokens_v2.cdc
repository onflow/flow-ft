import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import ExampleToken from "ExampleToken"

/// This transaction is a template for a transaction that could be used by anyone to send tokens to another account
/// through a switchboard using the deposit method but before depositing we will explicitly check whether receiving
/// capability is borrowable or not and if yes then it will deposit the vault to the receiver capability.
///
transaction(to: Address, amount: UFix64) {

    // The reference to the vault from the payer's account
    let vaultRef: auth(FungibleToken.Withdrawable) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdrawable) &ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
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
