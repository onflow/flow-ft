
import FungibleToken from "../../contracts/FungibleToken.cdc"
import ExampleToken from "../../contracts/ExampleToken.cdc"
import PrivateReceiverForwarder from "../../contracts/PrivateReceiverForwarder.cdc"

/// This transaction adds a Vault, a private receiver forwarder
/// a balance capability, and a public capability for the receiver

transaction {

    prepare(signer: AuthAccount) {
        if signer.getCapability<&PrivateReceiverForwarder.Forwarder>(PrivateReceiverForwarder.PrivateReceiverPublicPath).check() {
            // private forwarder was already set up
            return
        }

        if signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) == nil {
            // Create a new ExampleToken Vault and put it in storage
            signer.save(
                <-ExampleToken.createEmptyVault(),
                to: ExampleToken.VaultStoragePath
            )
        }

        signer.link<&{FungibleToken.Receiver}>(
            /private/exampleTokenReceiver,
            target: ExampleToken.VaultStoragePath
        )

        let receiverCapability = signer.getCapability<&{FungibleToken.Receiver}>(/private/exampleTokenReceiver)

        // Create a public capability to the Vault that only exposes
        // the balance field through the Balance interface
        signer.link<&ExampleToken.Vault{FungibleToken.Balance}>(
            ExampleToken.VaultPublicPath,
            target: ExampleToken.VaultStoragePath
        )

        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        signer.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        signer.link<&PrivateReceiverForwarder.Forwarder>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )
    }
}
