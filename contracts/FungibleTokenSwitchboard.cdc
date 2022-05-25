import FungibleToken from "./FungibleToken.cdc"

pub contract FungibleTokenSwitchboard{
    
    
    pub resource interface SwitchboardPublic{
        pub fun getVaultCapabilities() : [Capability<&AnyResource{FungibleToken.Receiver}>]
    }

    pub resource interface SwitchboardManager{
        pub fun addVaultCapability(capability : Capability<&AnyResource{FungibleToken.Receiver}>)
        pub fun removeVaultCapability(capability : Capability<&AnyResource{FungibleToken.Receiver}>)
    }
    
    /// Switchboard
    ///
    /// 
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic{
        
        pub var fungibleTokenReceiverCapabilities: [Capability<&{FungibleToken.Receiver}>]

        pub fun getVaultCapabilities() : [Capability<&AnyResource{FungibleToken.Receiver}>]{
            return self.fungibleTokenReceiverCapabilities
        }

        pub fun addVaultCapability(capability: Capability<&{FungibleToken.Receiver}>){
            self.fungibleTokenReceiverCapabilities.append(capability)
        }

        pub fun removeVaultCapability(capability: Capability<&{FungibleToken.Receiver}>){

        }
         
        pub fun deposit(from: @FungibleToken.Vault){

        }

        init(){
            self.fungibleTokenReceiverCapabilities = []
        }
    }

    // createNewSwitchboard creates a new Switchboard
    pub fun createNewSwitchboard(): @Switchboard{
        return <- create Switchboard()
    }


}