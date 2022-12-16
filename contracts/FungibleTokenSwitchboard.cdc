import FungibleToken from "./FungibleToken.cdc"

/// The contract that allows an account to receive payments in multiple fungible
/// tokens using a single `{FungibleToken.Receiver}` capability.
/// This capability should ideally be stored at the 
/// `FungibleTokenSwitchboard.ReceiverPublicPath = /public/GenericFTReceiver`
/// but it can be stored anywhere.
/// 
pub contract FungibleTokenSwitchboard {
  
    // Storage and Public Paths
    pub let StoragePath: StoragePath
    pub let PublicPath: PublicPath
    pub let ReceiverPublicPath: PublicPath

    /// The event that is emitted when a new vault capability is added to a
    /// switchboard resource.
    /// 
    pub event VaultCapabilityAdded(type: Type, switchboardOwner: Address?, 
                                    capabilityOwner: Address?)

    /// The event that is emitted when a vault capability is removed from a 
    /// switchboard resource.
    /// 
    pub event VaultCapabilityRemoved(type: Type,  switchboardOwner: Address?, 
                                        capabilityOwner: Address?)

    /// The event that is emitted when a deposit can not be completed.
    /// 
    pub event NotCompletedDeposit(type: Type, amount: UFix64, 
                                    switchboardOwner: Address?)

    /// The interface that enforces the method to allow anyone to check on the
    /// available capabilities of a switchboard resource and also exposes the 
    /// deposit methods to deposit funds on it.
    /// 
    pub resource interface SwitchboardPublic {
        pub fun getVaultTypes(): [Type]
        pub fun deposit(from: @FungibleToken.Vault)
        pub fun safeDeposit(from: @FungibleToken.Vault): @FungibleToken.Vault?
    }

    /// The resource that stores the multiple fungible token receiver 
    /// capabilities, allowing the owner to add and remove them and anyone to 
    /// deposit any fungible token among the available types.
    /// 
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic {
       
        /// Dictionary holding the fungible token receiver capabilities, 
        /// indexed by the fungible token vault type.
        /// 
        access(contract) var receiverCapabilities: {Type: Capability<&{FungibleToken.Receiver}>}

        /// Adds a new fungible token receiver capability to the switchboard 
        /// resource.
        /// 
        /// @param capability: The capability to expose a certain fungible
        /// token vault deposit function through `{FungibleToken.Receiver}` that
        /// will be added to the switchboard.
        /// 
        pub fun addNewVault(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed to by the capability we 
            // want to store inside the switchboard
            let vaultRef = capability.borrow() 
                ?? panic ("Cannot borrow reference to vault from capability")
            // Check if there is a previous capability for this token, if not
            if (self.receiverCapabilities[vaultRef.getType()] == nil) {
                // use the vault reference type as key for storing the 
                // capability and then
                self.receiverCapabilities[vaultRef.getType()] = capability
                // emit the event that indicates that a new capability has been 
                // added
                emit VaultCapabilityAdded(type: vaultRef.getType(),
                                    switchboardOwner: self.owner?.address, 
                                    capabilityOwner: capability.address)
            } else {
                // If there was already a capability for that token, panic
                panic("There is already a vault in the Switchboard for this token")
            }
        }

        /// Adds a number of new fungible token receiver capabilities by using
        /// the paths where they are stored.
        ///                    
        /// @param paths: The paths where the public capabilities are stored.
        /// @param address: The address of the owner of the capabilities.
        /// 
        pub fun addNewVaultsByPath(paths: [PublicPath], address: Address) {
            // Get the account where the public capabilities are stored
            let owner = getAccount(address)
            // For each path, get the saved capability and store it 
            // into the switchboard's receiver capabilities dictionary
            for path in paths {
                let capability = owner.getCapability<&{FungibleToken.Receiver}>(path)
                // Borrow a reference to the vault pointed to by the capability  
                // we want to store inside the switchboard
                // If the vault was borrowed successfully...
                if let vaultRef = capability.borrow() {
                    // ...and if there is no previous capability added for that 
                    // token
                    if (self.receiverCapabilities[vaultRef!.getType()] == nil) {
                        // Use the vault reference type as key for storing the 
                        // capability
                        self.receiverCapabilities[vaultRef!.getType()] = capability
                        // and emit the event that indicates that a new   
                        // capability has been added
                        emit VaultCapabilityAdded(type: vaultRef.getType(), 
                            switchboardOwner: self.owner?.address, capabilityOwner: address)
                    }
                }
            }
        }

        /// Adds a new fungible token receiver capability to the switchboard 
        /// resource specifying which `Type`of `@FungibleToken.Vault` can be 
        /// deposited to it. Use it to include in your switchboard "wrapper"
        /// receivers such as a `@TokenForwarding.Forwarder`. It can also be
        /// used to overwrite the type attached to a certain capability without 
        /// having to remove that capability first.
        ///
        /// @param capability: The capability to expose a certain fungible
        /// token vault deposit function through `{FungibleToken.Receiver}` that
        /// will be added to the switchboard.
        ///
        /// @param type: The type of fungible token that can be deposited to that
        /// capability, rather than the `Type` from the reference borrowed from
        /// said capability
        /// 
        pub fun addNewVaultWrapper(capability: Capability<&{FungibleToken.Receiver}>, type: Type) {
            // Check if the capability is working
            assert(capability.check(), message: "The passed capability is not valid")
            // Use the type parameter as key for the capability
            self.receiverCapabilities[type] = capability
            // emit the event that indicates that a new capability has been 
            // added
            emit VaultCapabilityAdded(type: type,
                                    switchboardOwner: self.owner?.address, 
                                    capabilityOwner: capability.address)
        }
  

        /// Removes a fungible token receiver capability from the switchboard
        /// resource.
        /// 
        /// @param capability: The capability to a fungible token vault to be
        /// removed from the switchboard.
        /// 
        pub fun removeVault(capability: Capability<&{FungibleToken.Receiver}>) {
            // Borrow a reference to the vault pointed to by the capability we 
            // want to remove from the switchboard
            let vaultRef = capability.borrow()
                ?? panic ("Cannot borrow reference to vault from capability")
            // Use the vault reference to find the capability to remove
            self.receiverCapabilities.remove(key: vaultRef.getType())
            // Emit the event that indicates that a new capability has been 
            // removed
            emit VaultCapabilityRemoved(type: vaultRef.getType(),
                                    switchboardOwner: self.owner?.address, 
                                    capabilityOwner: capability.address)       
        }
        
        /// Takes a fungible token vault and routes it to the proper fungible 
        /// token receiver capability for depositing it.
        /// 
        /// @param from: The deposited fungible token vault resource.
        /// 
        pub fun deposit(from: @FungibleToken.Vault) {
            // Get the capability from the ones stored at the switchboard
            let depositedVaultCapability = self
                .receiverCapabilities[from.getType()] 
                ?? panic ("The deposited vault is not available on this switchboard")
            // Borrow the reference to the desired vault
            let vaultRef = depositedVaultCapability.borrow() 
                ?? panic ("Can not borrow a reference to the the vault")
            vaultRef.deposit(from: <-from)
        }

        /// Takes a fungible token vault and tries to route it to the proper
        /// fungible token receiver capability for depositing the funds, 
        /// avoiding panicking if the vault is not available.
        ///             
        /// @param vaultType: The type of the ft vault that wants to be 
        /// deposited.
        /// 
        /// @return The deposited fungible token vault resource, without the
        /// funds if the deposit was successful, or still containing the funds
        /// if the reference to the needed vault was not found.
        /// 
        pub fun safeDeposit(from: @FungibleToken.Vault): @FungibleToken.Vault? {
            // Try to get the proper vault capability from the switchboard
            // If the desired vault is present on the switchboard...
            if let depositedVaultCapability = self
                                        .receiverCapabilities[from.getType()] {
                // We try to borrow a reference to the vault from the capability
                // If we can borrow a reference to the vault...
                if let vaultRef =  depositedVaultCapability.borrow() {
                    // We deposit the funds on said vault
                    vaultRef.deposit(from: <-from
                                                .withdraw(amount: from.balance))
                }
            }
            // if deposit failed for some reason
            if from.balance > 0.0 {
                emit NotCompletedDeposit(type: from.getType(), 
                                        amount: from.balance, 
                                        switchboardOwner: self.owner?.address)              
                return <-from
            }
            destroy from 
            return nil
        }

        /// A getter function to know which tokens a certain switchboard 
        /// resource is prepared to receive.
        ///
        /// @return The keys from the dictionary of stored 
        /// `{FungibleToken.Receiver}` capabilities that can be effectively 
        /// borrowed.
        ///
        pub fun getVaultTypes(): [Type] {
            let effectiveTypes: [Type] = []
            for vaultType in self.receiverCapabilities.keys {
                if self.receiverCapabilities[vaultType]!.check() {
                    effectiveTypes.append(vaultType)
                }
            }
            return effectiveTypes
        }

        init() {
            // Initialize the capabilities dictionary
            self.receiverCapabilities = {}
        }

    }

    /// Function that allows to create a new blank switchboard. A user must call
    /// this function and store the returned resource in their storage.
    ///
    pub fun createSwitchboard(): @Switchboard {
        return <-create Switchboard()
    }

    init() {
        self.StoragePath = /storage/fungibleTokenSwitchboard
        self.PublicPath = /public/fungibleTokenSwitchboardPublic
        self.ReceiverPublicPath = /public/GenericFTReceiver
    }

}
