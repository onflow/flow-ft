// This transaction is a template for a transaction to allow 
// anyone to add a Vault resource to their account so that 
// they can use the exampleToken

import FungibleToken from "./../contracts/FungibleToken.cdc"
import ExampleToken from "./../contracts/ExampleToken.cdc"
import MetadataViews from "./../contracts/utility/MetadataViews.cdc"

transaction () {

    prepare(signer: AuthAccount) {

        // Return early if the account already stores a ExampleToken Vault
        if signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) != nil {
            return
        }

        // Create a new ExampleToken Vault and put it in storage
        signer.save(
            <-ExampleToken.createEmptyVault(),
            to: ExampleToken.VaultStoragePath
        )

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&ExampleToken.Vault{FungibleToken.Receiver}>(
            ExampleToken.ReceiverPublicPath,
            target: ExampleToken.VaultStoragePath
        )

        // Create a public capability to the Vault that exposes the Balance and Resolver interfaces
        signer.link<&ExampleToken.Vault{FungibleToken.Balance, MetadataViews.Resolver}>(
            ExampleToken.VaultPublicPath,
            target: ExampleToken.VaultStoragePath
        )
    }
}
