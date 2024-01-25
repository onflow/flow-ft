
import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import PrivateReceiverForwarder from "PrivateReceiverForwarder"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

/// This transaction adds a Vault, a private receiver forwarder
/// a balance capability, and a public capability for the receiver

transaction {

    prepare(signer: auth(IssueStorageCapabilityController, PublishCapability, SaveValue) &Account) {
        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())
            ?? panic("Could not get vault data view for the contract")

        if signer.capabilities.get<&PrivateReceiverForwarder.Forwarder>(PrivateReceiverForwarder.PrivateReceiverPublicPath) != nil {
            // private forwarder was already set up
            return
        }

        if signer.storage.check<&ExampleToken.Vault>(from: vaultData.storagePath) == false {
            // Create a new ExampleToken Vault and put it in storage
            signer.storage.save(
                <-ExampleToken.createEmptyVault(vaultType: Type<ExampleToken.Vault>()),
                to: vaultData.storagePath
            )
        }

        // Create a public Vault Capability if needed
        if signer.capabilities.borrow<&{FungibleToken.Vault}>(vaultData.metadataPath) == nil {
            let vaultCap = signer.capabilities.storage.issue<&ExampleToken.Vault>(
                    vaultData.storagePath
                )
            signer.capabilities.publish(vaultCap, at: vaultData.metadataPath)
        }

        // Issue a Receiver Capability targetting the ExampleToken Vault
        let receiverCapability = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(
                vaultData.storagePath
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
