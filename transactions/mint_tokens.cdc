
// This transaction is a template for a transaction that
// could be used by the admin account to mint new tokens
// and deposit them in another account
//
// The minting amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x02
import FlowToken from 0x03

transaction {

    // Vault resource that holds the tokens that are being minted
    var vault: @FlowToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's MintAndBurn resource in storage
        let mintAndBurn = signer.borrow<&FlowToken.MintAndBurn>(from: /storage/flowTokenMintAndBurn)
            ?? panic("Couldn't borrow MintAndBurn reference from storage")

        // Mint 10 new tokens
        self.vault <- mintAndBurn.mintTokens(amount: 10.0)
    }

    execute {
        // Get the recipient's public account object
        let recipient = getAccount(0x03)

        // Get a reference to the recipient's Receiver
        let receiver = recipient.getCapability(/public/flowTokenReceiver)!
            .borrow<&FlowToken.Vault{FungibleToken.Receiver}>()
            ?? panic("Couldn't borrow receiver reference from the recipient")

        // Deposit the newly minted token in the recipient's Receiver
        receiver.deposit(from: <-self.vault)
    }
}
 