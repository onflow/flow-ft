import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

transaction {

    let exampleTokenVaultCapabilty: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {
      self.exampleTokenVaultCapabilty = signer.getCapability<&{FungibleToken.Receiver}>(ExampleToken.ReceiverPublicPath)
        
      self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
        (from: FungibleTokenSwitchboard.SwitchboardStoragePath) ?? panic("Could not borrow reference to switchboard")

    }

    execute {
      self.switchboardRef.removeVaultCapability(capability: self.exampleTokenVaultCapabilty)
    }

}