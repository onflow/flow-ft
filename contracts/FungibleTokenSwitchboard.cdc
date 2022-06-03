import FungibleToken from "./FungibleToken.cdc"

/// FungibleTokenSwitchboard
///
/// The contract that allows an account to receive payments in multiple fungible
/// tokens using a single {FungibleToken.Receiver}
///
pub contract FungibleTokenSwitchboard {
    
    // Storage and Public Paths
    pub let StoragePath: StoragePath
    pub let PublicPath: PublicPath
    pub let ReceiverPublicPath: PublicPath
    
    /// SwitchboardInitialized
    ///
    /// The event that is emited when a new Switchboard resource is created
    ///
    pub event SwitchboardInitialized(switchboardResourceID: UInt64)

    /// VaultCapacityAdded
    ///
    /// The event that is emitted when a new vault capacity is added to a
    /// switchboard resource
    ///
    pub event VaultCapabilityAdded(type: Type)

    /// VaultCapacityRemoved
    ///
    /// The event that is emitted vault capacity is added removed from a 
    /// switchboard resource
    ///
    pub event VaultCapabilityRemoved(type: Type)
    
    /// SwitchboardPublic
    ///
    /// The interface that enforces the method to allow anyone to check on the
    /// available capabilities of a switchboard resource and also exposes the 
    /// deposit method to deposit funds on it
    ///
    pub resource interface SwitchboardPublic {
        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>}
        pub fun deposit(from: @FungibleToken.Vault)
    }
    
    /// Switchboard
    /// The resource that stores the multiple fungible token receiver capabilities,
    /// allowing the owner to add and remove then and anyone to deposit any
    /// fungible token among the available capabilities
    ///
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic {
        
        /// fungibleTokenReceiverCapabilities
        /// Dictoionary holding the fungible token receiver capabilities, 
        /// indexed by the fungible token vault type
        ///
        pub var fungibleTokenReceiverCapabilities: {Type: Capability<&{FungibleToken.Receiver}>}

        /// addVaultCapability adds a new fungible token receiver capability
        ///                    to the switchboard resource
        ///            
        /// 
        /// Parameters: capability: The capability to expose a certain fungible
        /// token vault deposit function through {FungibleToken.Receiver} that
        /// will be added to the switchboard
        ///
        pub fun addVaultCapability(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed by the capability we want
            // store inside the switchboard
            let vaultRef = capability.borrow() 
                ?? panic ("Cannot borrow reference to vault from capability")
            // Use the vault reference type as key for storing the capability
            self.fungibleTokenReceiverCapabilities[vaultRef.getType()] = capability
            // Emit the event that indicates that a new capability has been added
            emit VaultCapabilityAdded(type: vaultRef.getType())
        }

        /// removeVaultCapability removes a fungible token receiver capability 
        ///                       from the switchboard resource
        /// 
        /// Parameters: capability: The capability to a fungible token vault 
        ///                         to be removed from the switchboard
        ///
        pub fun removeVaultCapability(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed by the capability we want
            // store inside the switchboard            
            let vaultRef = capability.borrow() 
                ?? panic ("Cannot borrow reference to vault from capability")
            // Use the vault reference to find the capability to remove
            self.fungibleTokenReceiverCapabilities.remove(key: vaultRef.getType())
            // Emit the event that indicates that a new capability has been removed
            emit VaultCapabilityRemoved(type: vaultRef.getType())            
        }
        
        /// deposit Takes a fungible token vault and routes it to the proper
        ///         fungible token receiver capability for depositing it
        /// 
        /// Parameters: from: The deposited fungible token vault resource
        ///        
        pub fun deposit(from: @FungibleToken.Vault) {
            let depositedVaultCapability = self.fungibleTokenReceiverCapabilities[from.getType()] ?? 
                panic ("The deposited vault is not available on this switchboard")
            let vaultRef = depositedVaultCapability.borrow() ?? 
                panic ("Can not borrow a reference to the the vault")
            vaultRef.deposit(from: <- from)
        }

        /// getVaultCabilities function for get to know which tokens a certain
        /// switchboard resource is prepared to receive
        ///
        /// Returns: The dictionary of stored {FungibleToken.Receiver} 
        /// capabilities
        ///
        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>} {
            return self.fungibleTokenReceiverCapabilities
        }

        init() {
            // Initialize the capabilities dictionary
            self.fungibleTokenReceiverCapabilities = {}
            // Emit the event that indicates that a new switchboard resource
            // has been created
            emit SwitchboardInitialized(switchboardResourceID: self.uuid)
        }
    }

    /// createSwitchboard
    ///
    /// Function that allows to create a new blank switchboard. A user must call
    /// this function and store the returned resource in their storage
    ///
    pub fun createSwitchboard(): @Switchboard {
        return <- create Switchboard()
    }

    init() {
        self.StoragePath = StoragePath(identifier: "fungibleTokenSwitchboard")!
        self.PublicPath = PublicPath(identifier: "fungibleTokenSwitchboardPublic")!
        self.ReceiverPublicPath = PublicPath(identifier: "GenericFTReceiver")!
    }

}
 