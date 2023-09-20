// This transaction is a template for a transaction to allow
// anyone to add a Vault resource to their account so that
// they can use the exampleToken

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import ViewResolver from "ViewResolver"

transaction () {

    prepare(signer: auth(BorrowValue) &Account) {

        // Return early if the account already stores a ExampleToken Vault
        if signer.storage.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath) != nil {
            return
        }

        let vault <- ExampleToken.createEmptyVault()

        // Create a new ExampleToken Vault and put it in storage
        signer.storage.save(
            <-vault,
            to: ExampleToken.VaultStoragePath
        )

        // Create a public capability to the Vault that exposes the Receiver, Balance, and Resolver interfaces
        let vaultCap = signer.link<&{FungibleToken.Receiver, FungibleToken.Balance, ViewResolver.Resolver}>(
            ExampleToken.VaultStoragePath
        )
        signer.capabilities.publish<&{FungibleToken.Receiver, FungibleToken.Balance, ViewResolver.Resolver}>(
            vaultCap,
            at: ExampleToken.VaultPublicPath,
        )
    }
}
