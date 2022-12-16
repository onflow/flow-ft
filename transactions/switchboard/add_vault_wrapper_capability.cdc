import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that
// could be used by anyone to add a new vault wrapper
// capability to their switchboard resource
transaction {

    let tokenForwarderCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

        // Get the token forwarder capability from the signer's account
        self.tokenForwarderCapability = 
            signer.getCapability<&{FungibleToken.Receiver}>
                                (ExampleToken.ReceiverPublicPath)
        
        // Check if the receiver capability exists
        assert(self.tokenForwarderCapability.check(), 
            message: "Signer does not have a working fungible token receiver capability")
        
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
            (from: FungibleTokenSwitchboard.StoragePath) 
            ?? panic("Could not borrow reference to switchboard")
    
    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVaultWrapper(capability: self.tokenForwarderCapability, type: Type<@ExampleToken.Vault>())
    
    }

}
