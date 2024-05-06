import "FungibleToken"
import "ExampleToken"
import "PrivateReceiverForwarder"
import "FungibleTokenMetadataViews"

/// This transaction is used to create a user's Flow account with a private forwarder

transaction {

    /// New Account that will hold the forwarder
    let newAccount: auth(Storage, Contracts, Keys, Inbox, Capabilities) &Account

    prepare(signer: auth(BorrowValue) &Account) {
        self.newAccount = Account(payer: signer)
    }

    execute {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not get vault data view for the contract")

        // Save a regular vault to the new account
        self.newAccount.storage.save(<-ExampleToken.createEmptyVault(vaultType: Type<@ExampleToken.Vault>()), to: vaultData.storagePath)

        // Issue a Receiver Capability targetting the ExampleToken Vault
        let receiverCapability = self.newAccount.capabilities.storage.issue<&{FungibleToken.Receiver}>(
            vaultData.storagePath
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