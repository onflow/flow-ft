import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import PrivateReceiverForwarder from "PrivateReceiverForwarder"

/// This transaction is used to create a user's Flow account with a private forwarder

transaction {

    /// New Account that will hold the forwarder
    let newAccount: auth(Storage, Contracts, Keys, Inbox, Capabilities) &Account

    prepare(payer: auth(BorrowValue) &Account) {
        self.newAccount = Account(payer: payer)
    }

    execute {

        // Save a regular vault to the new account
        self.newAccount.storage.save(<-ExampleToken.createEmptyVault(), to: ExampleToken.VaultStoragePath)

        // Issue a Receiver Capability targetting the ExampleToken Vault
        let receiverCapability = self.newAccount.capabilities.storage.issue<&{FungibleToken.Receiver}>(
            ExampleToken.VaultStoragePath
        )

        // Use the private receiver to create a private forwarder
        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        // Save the private forwarder to account storage
        self.newAccount.storage.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        // Issue a Capability to the Forwarder resource
        let forwarderCap = self.newAccount.capabilities.storage.issue<&PrivateReceiverForwarder.Forwarder>(
                PrivateReceiverForwarder.PrivateReceiverStoragePath
            )
        // Publish the Capability to the Forwarder resource
        self.newAccount.capabilities.publish(
            forwarderCap,
            at: PrivateReceiverForwarder.PrivateReceiverPublicPath
        )

    }
}