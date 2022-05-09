import FungibleToken from "../../contracts/FungibleToken.cdc"
import ExampleToken from "../../contracts/ExampleToken.cdc"
import PrivateReceiverForwarder from "../../contracts/PrivateReceiverForwarder.cdc"

// This transaction is used to create a user's Flow account with a private forwarder

transaction {
    prepare(payer: AuthAccount) {
        let acct = AuthAccount(payer: payer)

        acct.save(<-ExampleToken.createEmptyVault(),
            to: /storage/exampleTokenVault
        )

        // Create a private receiver
        let receiverCapability = acct.link<&{FungibleToken.Receiver}>(
            /private/exampleTokenReceiver,
            target: /storage/exampleTokenVault
        )!

        // Use the private receiver to create a private forwarder
        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        acct.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        acct.link<&PrivateReceiverForwarder.Forwarder>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )
    }
}