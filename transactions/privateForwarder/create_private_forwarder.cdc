import "FungibleToken"
import "ExampleToken"
import "PrivateReceiverForwarder"
import "FungibleTokenMetadataViews"

// This transaction creates a new private receiver in an account that 
// doesn't already have a private receiver or a public token receiver
// but does already have a Vault

transaction {

    prepare(signer: auth(IssueStorageCapabilityController, PublishCapability, SaveValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Issue a Receiver Capability targetting the ExampleToken Vault
        let receiverCapability = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(
            vaultData.storagePath
        )
        // Create the Forwarder resource
        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)
        // Save the Forwarder resource to storage
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
