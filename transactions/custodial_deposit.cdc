
// This transaction is a template for a transaction that could be
//used by anyone to deposit their tokens into a central contract that is
// controlled by a custodial entity like an exchange.
// They would specify a user ID in the transaction so that
// the deposit would emit an event that indicates which account
// should be credited.
//
// The custodial service would simply create this transaction from a
// template, filling in the necessary fields, and have the user
// sign the transaction to deposit their funds.
//
// The withdraw amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x01
import FlowToken from 0x02
import CustodialDeposit from 0x03

transaction {

    // Vault resource that holds the tokens that are being transferred
    var sentVault: @FlowToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        let storedVault = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)!

        // Withdraw 10 tokens from the signer's stored vault
        self.sentVault <- storedVault.withdraw(amount: 10.0)
    }

    execute {
        // Get the custodial service's public account object
        let recipient = getAccount(0x03)

        // Get the custodial service's public reference to the resource
        // that emits events when deposits happen

        let receiver = recipient
            .getCapability(/public/depositResourcePublic)!
            .borrow<&{CustodialDeposit.DepositPublic}>()!

        // Deposit the withdrawn tokens to the recipient's Receiver.
        // Include a tag for the event that is emitted
        //
        receiver.taggedDeposit(from: <-self.sentVault, tag: "1234")
    }
}
