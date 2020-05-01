
// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the FlowTokens

import FungibleToken from 0x02
import FlowToken from 0x03

transaction {

    prepare(signer: AuthAccount) {

        // Create a new FlowToken Vault and put it in storage
        signer.save(<-FlowToken.createEmptyVault(), to: /storage/flowTokenVault)

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&{FungibleToken.Receiver}>(
            /public/flowTokenReceiver,
            target: /storage/flowTokenVault
        )

        // Create a public capability to the Vault that only exposes
        // the balance field through the Balance interface
        signer.link<&{FungibleToken.Balance}>(
            /public/flowTokenBalance,
            target: /storage/flowTokenVault
        )
    }
}
 