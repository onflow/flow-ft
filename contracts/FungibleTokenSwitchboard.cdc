import FungibleToken from "./FungibleToken.cdc"

pub contract FungibleTokenSwitchboard {
    
    pub let SwitchboardStoragePath: StoragePath
    pub let SwitchboardPublicPath: PublicPath
    
    pub resource interface SwitchboardPublic {
        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>}
    }
    
    /// Switchboard
    ///
    /// 
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic {
        
        pub var fungibleTokenReceiverCapabilities: {Type: Capability<&{FungibleToken.Receiver}>}

        pub fun getVaultCapabilities(): {Type: Capability<&{FungibleToken.Receiver}>}{
            return self.fungibleTokenReceiverCapabilities
        }

        pub fun addVaultCapability(capability: Capability<&{FungibleToken.Receiver}>){
            let vaultRef = capability.borrow()
            self.fungibleTokenReceiverCapabilities[vaultRef.getType()] = capability
        }

        pub fun removeVaultCapability(capability: Capability<&{FungibleToken.Receiver}>){
            self.fungibleTokenReceiverCapabilities.remove(key: capability.getType())
        }
         
        pub fun deposit(from: @FungibleToken.Vault){
            let depositedVaultCapability = self.fungibleTokenReceiverCapabilities[from.getType()] ?? 
                panic ("The deposited vault is not available on this switchboard")
            let vaultRef = depositedVaultCapability.borrow() ?? 
                panic ("Can not borrow a reference to the the vault")
            vaultRef.deposit(from: <- from)
        }

        init(){
            self.fungibleTokenReceiverCapabilities = {}
        }
    }

    init(){
        self.SwitchboardStoragePath = StoragePath(identifier: "fungibleTokenSwitchboard")!
        self.SwitchboardPublicPath = PublicPath(identifier: "fungibleTokenSwitchboardPublic")!
        let switchboard <- create Switchboard()        
        self.account.save(<- switchboard, to: self.SwitchboardStoragePath)
        self.account.link<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>(
            self.SwitchboardPublicPath,
            target: self.SwitchboardStoragePath
        )
    }

}
 