// This transaction is a template for a transaction that
// could be used by anyone to send tokens to another account
// through a switchboard, as long as they have set up their
// switchboard and have add the proper capability to it
//
// The withdraw amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

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
        self.sentVault <- self.vaultRef.withdraw(amount: amount)
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Switchboard Receiver
        let switchboardRef = recipient.getCapability(FungibleTokenSwitchboard.PublicPath)
            .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
			?? panic("Could not borrow receiver reference to switchboard!")

        // Deposit the withdrawn tokens in the recipient's switchboard receiver,
        // then deposit the returned vault in the signer's vault
        self.vaultRef.deposit(from: <- switchboardRef.safeDeposit(from: <-self.sentVault))
    }
}
