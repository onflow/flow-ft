import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that
// could be used by anyone to remove fungible token vault
// capability from their switchboard resource
transaction {

    let exampleTokenVaultCapabilty: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

      // Get the example token vault capability from the signer's account
      self.exampleTokenVaultCapabilty = signer.getCapability
                    <&{FungibleToken.Receiver}>(ExampleToken.ReceiverPublicPath)
      
      // Get a reference to the signers switchboard  
      self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
        (from: FungibleTokenSwitchboard.StoragePath) 
          ?? panic("Could not borrow reference to switchboard")

    }

    execute {

      // Remove the capability from the switchboard using the 
      // removeVault method
      self.switchboardRef.removeVault(capability: self.exampleTokenVaultCapabilty)
    
    }

}
