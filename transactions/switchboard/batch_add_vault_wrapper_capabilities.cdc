import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that could be used by anyone
// to add several capabilities that point to fungible token vaults of a different   
// `Type` and belong to a certain `Address`, to their switchboard resource.
transaction (address: Address) {

    let vaultPaths: [PublicPath]
    let vaultTypes: [Type]
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

        // Store the Example Token receiver's public path in the array of public 
        // paths that will be passed to the switchboard method
        self.vaultPaths = []
        self.vaultPaths.append(ExampleToken.ReceiverPublicPath)

        // Store the Example Token's type in the array of types that will be passed 
        // to the switchboard method
        self.vaultTypes = []
        self.vaultTypes.append(Type<@ExampleToken.Vault>())
      
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
                                         (from: FungibleTokenSwitchboard.StoragePath) 
                                ?? panic("Could not borrow reference to switchboard")
    
    }

    execute {

      // Add the capability(ies) to the switchboard using addNewVaultWrappersByPath
      self.switchboardRef.addNewVaultWrappersByPath(paths: self.vaultPaths, 
                                            types: self.vaultTypes, address: address)

    }

}