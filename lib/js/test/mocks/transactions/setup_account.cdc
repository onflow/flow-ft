// This transaction is a template for a transaction to allow 
// anyone to add a Vault resource to their account so that 
// they can use the TestToken

import FungibleToken from "../contracts/Token.cdc"
import TestToken from "../contracts/TestToken.cdc"
import MetadataViews from "../../../../../contracts/utility/MetadataViews.cdc"

transaction () {

    prepare(signer: AuthAccount) {

        // Return early if the account already stores a TestToken Vault
        if signer.borrow<&TestToken.Vault>(from: TestToken.VaultStoragePath) != nil {
            return
        }

        // Create a new TestToken Vault and put it in storage
        signer.save(
            <-TestToken.createEmptyVault(),
            to: TestToken.VaultStoragePath
        )

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&TestToken.Vault{Token.Receiver}>(
            TestToken.ReceiverPublicPath,
            target: TestToken.VaultStoragePath
        )

        // Create a public capability to the Vault that exposes the Balance and Resolver interfaces
        signer.link<&TestToken.Vault{Token.Balance, MetadataViews.Resolver}>(
            TestToken.VaultPublicPath,
            target: TestToken.VaultStoragePath
        )
    }
}
