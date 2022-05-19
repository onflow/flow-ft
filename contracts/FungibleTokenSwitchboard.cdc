import FungibleToken from "./FungibleToken.cdc"

pub contract FungibleTokenSwitchboard{
    
    
    pub resource interface SwitchboardPublic{
        pub fun getVaultCapabilities() : [Capability<&AnyResource{FungibleToken.Receiver}>]
    }
    
    /// Switchboard
    ///
    /// 
    pub resource Switchboard: FungibleToken.Receiver, SwitchboardPublic{
        
        pub var switches: [Capability<&{FungibleToken.Receiver}>]


        pub fun deposit(from: @FungibleToken.Vault){

        }

        pub fun addVaultCapability(switch: Capability<&{FungibleToken.Receiver}>){
            self.switches.append(switch)
        }

        pub fun getVaultCapabilities() : [Capability<&AnyResource{FungibleToken.Receiver}>]{
            return self.switches
        }

        init(){
            self.switches = []
        }
    }

    // createNewSwitchboard creates a new Switchboard
    pub fun createNewSwitchboard(): @Switchboard{
        return <- create Switchboard()
    }


}
