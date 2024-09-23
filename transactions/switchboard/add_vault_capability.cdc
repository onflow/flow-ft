import "FungibleToken"
import "FungibleTokenSwitchboard"
import "ExampleToken"
import "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to add a new fungible token vault
/// capability to their switchboard resource
///
transaction {

    let exampleTokenVaultCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue, IssueStorageCapabilityController, PublishCapability, SaveValue, UnpublishCapability) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        /* ExampleToken Vault configuration */
        //
        // Configure an ExampleToken Vault if needed
        if signer.storage.borrow<&ExampleToken.Vault>(from: vaultData.storagePath) == nil {
            // Create a new ExampleToken Vault and save it in storage
            signer.storage.save(<-ExampleToken.createEmptyVault(vaultType: Type<@ExampleToken.Vault>()), to: vaultData.storagePath)
            // Clear existing Capabilities at canonical paths
            signer.capabilities.unpublish(vaultData.metadataPath)
            signer.capabilities.unpublish(vaultData.receiverPath)
            // Issue Vault & Receiver Capabilities
            let vaultCap = signer.capabilities.storage.issue<&{FungibleToken.Balance, FungibleToken.Vault}>(vaultData.storagePath)
            let receiverCap = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(vaultData.storagePath)
            // Publish Capabilities
            signer.capabilities.publish(vaultCap, at: vaultData.metadataPath)
            signer.capabilities.publish(receiverCap, at: vaultData.receiverPath)
        }
        
        // Get the example token vault capability from the signer's account
        self.exampleTokenVaultCapability = signer.capabilities.get<&{FungibleToken.Receiver}>(
                vaultData.receiverPath)
        
        // Check if the receiver capability exists
        assert(
            self.exampleTokenVaultCapability.check(), 
            message: "The signer does not store a ExampleToken Vault capability at the path "
                .concat(vaultData.receiverPath.toString())
                .concat(". The signer must initialize their account with this object first!")
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
        self.switchboardRef = signer.storage.borrow<auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath)
			?? panic("The signer does not store a FungibleToken Switchboard object at the path "
                .concat(FungibleTokenSwitchboard.StoragePath.toString())
                .concat(". The signer must initialize their account with this object first!"))
    
    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVault(capability: self.exampleTokenVaultCapability)
    
    }

}
