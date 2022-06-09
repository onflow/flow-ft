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
        pub fun getVaultTypes(): [Type]
        pub fun deposit(from: @FungibleToken.Vault)
        pub fun safeDeposit(from: @FungibleToken.Vault): @FungibleToken.Vault
    }
    
    /// Switchboard
    /// The resource that stores the multiple fungible token receiver capabilities,
    /// allowing the owner to add and remove them and anyone to deposit any
    /// fungible token among the available capabilities
    ///
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic {
        
        /// receiverCapabilities
        /// Dictionary holding the fungible token receiver capabilities, 
        /// indexed by the fungible token vault type
        ///
        access(contract) var receiverCapabilities: {Type: Capability<&{FungibleToken.Receiver}>}

        /// addNewVault adds a new fungible token receiver capability
        ///                    to the switchboard resource
        ///            
        /// Parameters: capability: The capability to expose a certain fungible
        /// token vault deposit function through {FungibleToken.Receiver} that
        /// will be added to the switchboard
        ///
        pub fun addNewVault(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed by the capability we want
            // to store inside the switchboard
            let vaultRef = capability.borrow() 
                ?? panic ("Cannot borrow reference to vault from capability")
            // Use the vault reference type as key for storing the capability
            self.receiverCapabilities[vaultRef.getType()] = capability
            // Emit the event that indicates that a new capability has been added
            emit VaultCapabilityAdded(type: vaultRef.getType())
        }

        /// addNewVaultsByPath adds a number of new fungible token receiver 
        ///                    capabilities by using the paths where they are
        ///                    stored
        ///
        /// Parameters: paths: The paths where the public capabilities are stored
        ///             address: The address of the owner of the capabilities
        ///
        pub fun addNewVaultsByPath(paths: [PublicPath], address: Address) {
            // Get the account where the public capabilities are stored
            let owner = getAccount(address)
            // For each path, get the saved capability and store it 
            // into the switchboard's receiver capabilities dictionary 
            for path in paths {
                let capability = owner.getCapability<&{FungibleToken.Receiver}>(path)
                // Borrow a reference to the vault pointed by the capability we want
                // to store inside the switchboard
                let vaultRef = capability.borrow() 
                    ?? panic ("Cannot borrow reference to vault from capability")
                // Use the vault reference type as key for storing the capability
                self.receiverCapabilities[vaultRef.getType()] = capability
                // Emit the event that indicates that a new capability has been added
                emit VaultCapabilityAdded(type: vaultRef.getType())
            }
        }

        /// removeVault removes a fungible token receiver capability 
        ///                       from the switchboard resource
        /// 
        /// Parameters: capability: The capability to a fungible token vault 
        ///                         to be removed from the switchboard
        ///
        pub fun removeVault(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed by the capability we want
            // store inside the switchboard            
            let vaultRef = capability.borrow() 
                ?? panic ("Cannot borrow reference to vault from capability")
            // Use the vault reference to find the capability to remove
            self.receiverCapabilities.remove(key: vaultRef.getType())
            // Emit the event that indicates that a new capability has been removed
            emit VaultCapabilityRemoved(type: vaultRef.getType())            
        }
        
        /// deposit Takes a fungible token vault and routes it to the proper
        ///         fungible token receiver capability for depositing it
        /// 
        /// Parameters: from: The deposited fungible token vault resource
        ///        
        pub fun deposit(from: @FungibleToken.Vault) {
            let depositedVaultCapability = self.receiverCapabilities[from.getType()] ?? 
                panic ("The deposited vault is not available on this switchboard")
            let vaultRef = depositedVaultCapability.borrow() ?? 
                panic ("Can not borrow a reference to the the vault")
            vaultRef.deposit(from: <- from)
        }

        /// safeDeposit Takes a fungible token vault and tries to route it to the
        ///             proper fungible token receiver capability for depositing
        ///             the funds, avoiding panicking if the vault is not available
        ///             
        /// Parameters: vaultType: The type of the ft vault that wants to be 
        ///                        deposited
        ///
        /// Returns: The deposited fungible token vault resource, without the
        ///          funds if the deposit was succesful, or still containing the
        ///          funds if the reference to the needed vault was not found
        ///
        pub fun safeDeposit(from: @FungibleToken.Vault): @FungibleToken.Vault {
            // Try to get the proper vault capability from the switchboard
            let depositedVaultCapability = self.receiverCapabilities[from.getType()]
            // If the desired vault is present on the switchboard...
            if  depositedVaultCapability != nil {
                // We try to borrow a reference to the vault from the capability
                let vaultRef = depositedVaultCapability!.borrow()
                // Finally if we can borrow a reference to the vault...
                if vaultRef != nil {
                    // We deposit the funds on said vault
                    vaultRef!.deposit(from: <- from.withdraw(amount: from.balance) )
                }
            }
            // Either way we return the deposited vault avoiding panicking the tx
            // if the vault was not found on the switchboard or if the reference
            // to it could not be borrowed
            return <- from
        }

        /// getVaultTypes function for get to know which tokens a certain
        /// switchboard resource is prepared to receive
        ///
        /// Returns: The keys from the dictionary of stored {FungibleToken.Receiver} 
        /// capabilities
        ///
        pub fun getVaultTypes(): [Type] {
            return self.receiverCapabilities.keys
        }

        init() {
            // Initialize the capabilities dictionary
            self.receiverCapabilities = {}
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
        self.StoragePath = /storage/fungibleTokenSwitchboard
        self.PublicPath = /public/fungibleTokenSwitchboardPublic
        self.ReceiverPublicPath = /public/GenericFTReceiver
    }

}
 