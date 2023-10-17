
import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import PrivateReceiverForwarder from "PrivateReceiverForwarder"

/// This transaction adds a Vault, a private receiver forwarder
/// a balance capability, and a public capability for the receiver

transaction {

    prepare(signer: auth(IssueStorageCapabilityController, PublishCapability, SaveValue) &Account) {
        if signer.capabilities.get<&PrivateReceiverForwarder.Forwarder>(PrivateReceiverForwarder.PrivateReceiverPublicPath) != nil {
            // private forwarder was already set up
            return
        }

        if signer.storage.check<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) == false {
            // Create a new ExampleToken Vault and put it in storage
            signer.storage.save(
                <-ExampleToken.createEmptyVault(),
                to: ExampleToken.VaultStoragePath
            )
        }

        // Create a public Vault Capability if needed
        if signer.capabilities.borrow<&{FungibleToken.Vault}>(ExampleToken.VaultPublicPath) == nil {
            let vaultCap = signer.capabilities.storage.issue<&{FungibleToken.Vault}>(
                    ExampleToken.VaultStoragePath
                )
            signer.capabilities.publish(vaultCap, at: ExampleToken.VaultPublicPath)
        }

        // Issue a Receiver Capability targetting the ExampleToken Vault
        let receiverCapability = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(
                ExampleToken.VaultStoragePath
            )

        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        signer.storage.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        // Issue a Capability to the Forwarder resource
        let forwarderCap = signer.capabilities.storage.issue<&PrivateReceiverForwarder.Forwarder>(
                PrivateReceiverForwarder.PrivateReceiverStoragePath
            )
        // Publish the Capability to the Forwarder resource
        signer.capabilities.publish(
            forwarderCap,
            at: PrivateReceiverForwarder.PrivateReceiverPublicPath
        )
    }
}
