// This transaction is a template for a transaction to allow
// anyone to add a Vault resource to their account so that
// they can use the exampleToken

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import ViewResolver from "ViewResolver"

transaction () {

    prepare(signer: AuthAccount) {

        // Return early if the account already stores a ExampleToken Vault
        if signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) != nil {
            return
        }

        let vault <- ExampleToken.createEmptyVault()

        // Create a new ExampleToken Vault and put it in storage
        signer.save(
            <-vault,
            to: ExampleToken.VaultStoragePath
        )

        // Create a public capability to the Vault that exposes the Receiver, Balance, and Resolver interfaces
        signer.link<&{FungibleToken.Receiver, FungibleToken.Balance, ViewResolver.Resolver}>(
            ExampleToken.VaultPublicPath,
            target: ExampleToken.VaultStoragePath
        )
    }
}
