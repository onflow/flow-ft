// This transaction is a template for a transaction that
// could be used by anyone to send tokens to another account
// that has been set up to receive tokens.
//
// The withdraw amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x02
import FlowToken from 0x03

transaction {

    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        let storedVault = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
            ?? panic("Unable to borrow a reference to the sender's Vault")

        // Withdraw 10 tokens from the signer's stored vault
        self.sentVault <- storedVault.withdraw(amount: 10.0)
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(0x04)

        // Get a reference to the recipient's Receiver
        let receiver = recipient
            .getCapability(/public/flowTokenReceiver)!
            .borrow<&{FungibleToken.Receiver}>() 
            ?? panic("Unable to borrow receiver reference for recipient")

        // Deposit the withdrawn tokens in the recipient's receiver
        receiver.deposit(from: <-self.sentVault)
    }
}
 