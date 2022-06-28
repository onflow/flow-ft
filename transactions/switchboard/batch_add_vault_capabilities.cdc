// This transaction is a template for a transaction that
// could be used by anyone to add a new fungible token vault
// capability to their switchboard resource

import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

transaction (address: Address) {

    let exampleTokenVaultPath: PublicPath
    let vaultPaths: [PublicPath]
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {
      // Get the example token vault path from the contract
      self.exampleTokenVaultPath = ExampleToken.ReceiverPublicPath
      // And store it in the array of public paths that will be passed to the
      // switchboard method
      self.vaultPaths = []
      self.vaultPaths.append(self.exampleTokenVaultPath)
      // Get a reference to the signers switchboard
      self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
        (from: FungibleTokenSwitchboard.StoragePath) 
          ?? panic("Could not borrow reference to switchboard")
    }

    execute {
      // Add the capability to the switchboard using addNewVault method
      self.switchboardRef.addNewVaultsByPath (paths: self.vaultPaths, address: address)
    }

}