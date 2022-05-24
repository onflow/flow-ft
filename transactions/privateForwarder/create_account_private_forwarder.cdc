import FungibleToken from "../../contracts/FungibleToken.cdc"
import ExampleToken from "../../contracts/ExampleToken.cdc"
import PrivateReceiverForwarder from "../../contracts/PrivateReceiverForwarder.cdc"

/// This transaction is used to create a user's Flow account with a private forwarder

transaction {

    /// New Account that will hold the forwarder
    let newAccount: AuthAccount

    prepare(payer: AuthAccount) {
        self.newAccount = AuthAccount(payer: payer)
    }

    execute {

        // Save a regular vault to the new account
        self.newAccount.save(<-ExampleToken.createEmptyVault(),
            to: ExampleToken.VaultStoragePath
        )

        // Create a private receiver
        let receiverCapability = self.newAccount.link<&{FungibleToken.Receiver}>(
            /private/exampleTokenReceiver,
            target: ExampleToken.VaultStoragePath
        )!

        // Use the private receiver to create a private forwarder
        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        // Save the private forwarder to account storage
        self.newAccount.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        // Link the forwarder to a private path
        self.newAccount.link<&PrivateReceiverForwarder.Forwarder>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )

    }
}