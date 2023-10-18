import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import ExampleToken from "ExampleToken"

// This transaction is a template for a transaction that could be used by anyone 
// to send tokens to another account through a switchboard using the safeDeposit
// method. This method will not panic if the switchboard does not have the
// capability to store the desired FT, returning the deposited vault instead.
// The withdraw amount and the account from getAccount would be the parameters 
// to the transaction.
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

        // Get a reference to the recipient's Switchboard Receiver
        let switchboardRef = recipient.getCapability(FungibleTokenSwitchboard.PublicPath)
            .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
			?? panic("Could not borrow receiver reference to switchboard!")    

        // Deposit the funds on the switchboard, if the deposit does not complete the method will return the funds
        // instead of panicking, so we have to recover those funds
        if let notDepositedVault <- switchboardRef.safeDeposit(from: <- sentVault.withdraw(amount: amount)){
            self.vaultRef.deposit(from: <-notDepositedVault)
        }

        destroy self.sentVault

    }

}
