import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that
// could be used by anyone to add a new fungible token vault
// capability to their switchboard resource
transaction {

    let exampleTokenVaultCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

        // Get the example token vault capability from the signer's account
        self.exampleTokenVaultCapability = 
            signer.getCapability<&{FungibleToken.Receiver}>
                                (ExampleToken.ReceiverPublicPath)
        
        // Check if the receiver capability exists
        assert(self.exampleTokenVaultCapability.check(), 
            message: "Signer does not have a Example Token receiver capability")
        
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
            (from: FungibleTokenSwitchboard.StoragePath) 
            ?? panic("Could not borrow reference to switchboard")
    
    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVault(capability: self.exampleTokenVaultCapability)
    
    }

}
