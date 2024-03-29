import "FungibleToken"
import "FungibleTokenSwitchboard"
import "ExampleToken"

/// This transaction is a template for a transaction that could be used by anyone to remove fungible token vault
/// capability from their switchboard resource
///
transaction(path: PublicPath) {

    let exampleTokenVaultCapabilty: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue) &Account) {

        // Get the capability from the signer's account
        self.exampleTokenVaultCapabilty = signer.capabilities.get<&{FungibleToken.Receiver}>(path)
            ?? panic("Signer does not have Receiver Capability at given path")

        // Get a reference to the signers switchboard
        self.switchboardRef = signer.storage.borrow<auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath
            ) ?? panic("Could not borrow reference to switchboard")

    }

    execute {

      // Remove the capability from the switchboard using the .removeVault() method
      self.switchboardRef.removeVault(capability: self.exampleTokenVaultCapabilty)

    }

}
