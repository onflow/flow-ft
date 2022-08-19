import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that could be used by anyone 
// to send tokens to another account through a switchboard using the safeDeposit
// method. This method will not panic if the switchboard does not have the
// capability to store the desired FT, returning the deposited vault instead.
// The withdraw amount and the account from getAccount would be the parameters 
// to the transaction.
transaction(to: Address, amount: UFix64) {
    
    // The reference to the vault from the payer's account
    let vaultRef: &ExampleToken.Vault
    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        self.sentVault <-self.vaultRef.withdraw(amount: amount)
    
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Switchboard Receiver
        let switchboardRef = recipient.getCapability(FungibleTokenSwitchboard.PublicPath)
            .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
			?? panic("Could not borrow receiver reference to switchboard!")    
        
        // Deposit the funds on the switchboard, if the deposit does not complete
        // the method will return the funds instead of panicking, so we have to
        // recover those funds
        if let notDepositedVault <-switchboardRef.safeDeposit(from: <- self.sentVault.withdraw(amount: amount)){
            self.vaultRef.deposit(from: <-notDepositedVault)
        }

        destroy self.sentVault
    
    }

}
