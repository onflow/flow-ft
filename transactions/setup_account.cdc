// This transaction is a template for a transaction to allow
// anyone to add a Vault resource to their account so that
// they can use the exampleToken

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import ViewResolver from "ViewResolver"

transaction () {

    prepare(signer: auth(BorrowValue, IssueStorageCapabilityController, PublishCapability, SaveValue) &Account) {

        // Return early if the account already stores a ExampleToken Vault
        if signer.storage.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) != nil {
            return
        }

        let vault <- ExampleToken.createEmptyVault()

        // Create a new ExampleToken Vault and put it in storage
        signer.storage.save(<-vault, to: ExampleToken.VaultStoragePath)

        // Create a public capability to the Vault that exposes the Vault interfaces
        let vaultCap = signer.capabilities.storage.issue<&{FungibleToken.Vault}>(
            ExampleToken.VaultStoragePath
        )
        signer.capabilities.publish(vaultCap, at: ExampleToken.VaultPublicPath)

        // Create a public Capability to the Vault's Receiver functionality
        let receiverCap = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(
            ExampleToken.VaultStoragePath
        )
        signer.capabilities.publish(receiverCap, at: ExampleToken.ReceiverPublicPath)
    }
}
