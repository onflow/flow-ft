import FungibleToken from "./FungibleToken.cdc"

pub contract FungibleTokenSwitchboard {

    pub event FungibleTokenSwitchboardInitialized()
    
    pub event SwitchboardInitialized(switchboardResourceID: UInt64)

    //pub event vaultCapabilityAdded()
    //pub event vaultCapabilityRemoved()
    
    pub let SwitchboardStoragePath: StoragePath
    pub let SwitchboardPublicPath: PublicPath
    pub let SwitchboardReceiverPublicPath: PublicPath
    
    pub resource interface SwitchboardPublic {
        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>}
    }
    
    /// Switchboard
    ///
    ///
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic {
        
        pub var fungibleTokenReceiverCapabilities: {Type: Capability<&{FungibleToken.Receiver}>}

        pub fun addVaultCapability(capability: Capability<&{FungibleToken.Receiver}>) {
            let vaultRef = capability.borrow() ?? panic ("Cannot borrow reference to vault from capability")
            self.fungibleTokenReceiverCapabilities[vaultRef.getType()] = capability
        }

        pub fun removeVaultCapability(capability: Capability<&{FungibleToken.Receiver}>) {
            self.fungibleTokenReceiverCapabilities.remove(key: capability.getType())
        }
         
        pub fun deposit(from: @FungibleToken.Vault) {
            let depositedVaultCapability = self.fungibleTokenReceiverCapabilities[from.getType()] ?? 
                panic ("The deposited vault is not available on this switchboard")
            let vaultRef = depositedVaultCapability.borrow() ?? 
                panic ("Can not borrow a reference to the the vault")
            vaultRef.deposit(from: <- from)
        }

        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>} {
            return self.fungibleTokenReceiverCapabilities
        }

        init() {
            self.fungibleTokenReceiverCapabilities = {}
            emit SwitchboardInitialized(switchboardResourceID: self.uuid)
        }
    }

    pub fun createSwitchboard(): @Switchboard {
        return <- create Switchboard()
    }

    init() {
        self.SwitchboardStoragePath = StoragePath(identifier: "fungibleTokenSwitchboard")!
        self.SwitchboardPublicPath = PublicPath(identifier: "fungibleTokenSwitchboardPublic")!
        self.SwitchboardReceiverPublicPath = PublicPath(identifier: "fungibleTokenSwitchboardReceiverPublic")!
        emit FungibleTokenSwitchboardInitialized()
    }

}
 