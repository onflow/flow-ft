// This transaction is a template for a transaction that could be used
// by an admin account to create a new MintAndBurn resource
// and store it in another account
//
// Both accounts need to sign the transaction because the transaction needs
// to be able to access both accounts' storage

import FungibleToken from 0x01
import FlowToken from 0x02

transaction {

    prepare(existingAdmin: AuthAccount, newAdmin: AuthAccount) {

        // Get a reference to the existing admin's MintAndBurn resource in storage
        let existingMintAndBurn = existingAdmin.borrow<&FlowToken.MintAndBurn>(from: /storage/flowTokenMintAndBurn)!

        // Use the existing admin's MintAndBurn resource to create a new one

        let newMintAndBurn <- existingMintAndBurn.createNewMinter(allowedAmount: 10.0)

        // Store the new MintAndBurn resource in the new admin's account's storage
        newAdmin.save(<-newMintAndBurn, to: /storage/flowTokenMintAndBurn)
    }
}
