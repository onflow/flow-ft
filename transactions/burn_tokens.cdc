// This transaction is a template for a transaction that
// could be used by the admin account to burn tokens
// from their stored Vault
//
// The burning amount would be a parameter to the transaction

import FungibleToken from 0x01
import FlowToken from 0x02

transaction {

    // Vault resource that holds the tokens that are being burned
    let vault: @FlowToken.Vault

    let mintAndBurn: &FlowToken.MintAndBurn

    prepare(signer: AuthAccount) {

        // Withdraw 10 tokens from the admin vault in storage
        self.vault <- signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)!
            .withdraw(amount: 10.0)

        // Create a reference to the admin MintAndBurn resource in storage
        self.mintAndBurn = signer.borrow<&FlowToken.MintAndBurn>(from: /storage/flowTokenMintAndBurn)
            ?? panic("Could not borrow a reference to the Burn resource")
    }

    execute {
        // burn the withdrawn tokens
        self.mintAndBurn.burnTokens(from: <-self.vault)
    }
}
 