import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import ExampleToken from "ExampleToken"

/// This transaction is a template for a transaction that could be used by anyone to add a new fungible token vault
/// capability to their switchboard resource
///
transaction {

    let exampleTokenVaultCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue, IssueStorageCapabilityController, PublishCapability, SaveValue, UnpublishCapability) &Account) {

        /* ExampleToken Vault configuration */
        //
        // Configure an ExampleToken Vault if needed
        if signer.storage.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) == nil {
            // Create a new ExampleToken Vault and save it in storage
            signer.storage.save(<-ExampleToken.createEmptyVault(), to: ExampleToken.VaultStoragePath)
            // Clear existing Capabilities at canonical paths
            signer.capabilities.unpublish(ExampleToken.VaultPublicPath)
            signer.capabilities.unpublish(ExampleToken.ReceiverPublicPath)
            // Issue Vault & Receiver Capabilities
            let vaultCap = signer.capabilities.storage.issue<&ExampleToken.Vault>(ExampleToken.VaultStoragePath)
            let receiverCap = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(ExampleToken.VaultStoragePath)
            // Publish Capabilities
            signer.capabilities.publish(vaultCap, at: ExampleToken.VaultPublicPath)
            signer.capabilities.publish(receiverCap, at: ExampleToken.ReceiverPublicPath)
        }
        
        // Get the example token vault capability from the signer's account
        self.exampleTokenVaultCapability = signer.capabilities.get<&{FungibleToken.Receiver}>(
                ExampleToken.ReceiverPublicPath
            ) ?? panic("Signer does not have a Example Token receiver capability")
        
        // Check if the receiver capability exists
        assert(
            self.exampleTokenVaultCapability.check(), 
            message: "Signer does not have a Example Token receiver capability"
        )
        
        /* Switchboard setup */
        //
        // Configure .Switchboard if needed
        if signer.storage.borrow<&FungibleTokenSwitchboard.Switchboard>(from: FungibleTokenSwitchboard.StoragePath) == nil {
            // Create a new Switchboard and save it in storage
            signer.storage.save(<-FungibleTokenSwitchboard.createSwitchboard(), to: FungibleTokenSwitchboard.StoragePath)
            // Clear existing Capabilities at canonical paths
            signer.capabilities.unpublish(FungibleTokenSwitchboard.ReceiverPublicPath)
            signer.capabilities.unpublish(FungibleTokenSwitchboard.PublicPath)
            // Issue Receiver & Switchboard Capabilities
            let receiverCap = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(
                    FungibleTokenSwitchboard.StoragePath
                )
            let switchboardPublicCap = signer.capabilities.storage.issue<&{FungibleTokenSwitchboard.SwitchboardPublic, FungibleToken.Receiver}>(
                    FungibleTokenSwitchboard.StoragePath
                )
            // Publish Capabilities
            signer.capabilities.publish(receiverCap, at: FungibleTokenSwitchboard.ReceiverPublicPath)
            signer.capabilities.publish(switchboardPublicCap, at: FungibleTokenSwitchboard.PublicPath)
        }
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.storage.borrow<&FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath
            ) ?? panic("Could not borrow reference to switchboard")
    
    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVault(capability: self.exampleTokenVaultCapability)
    
    }

}
