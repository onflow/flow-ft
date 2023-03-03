// This transaction is a template for a transaction to allow 
// anyone to add a Vault resource to their account so that 
// they can use the DemoToken

import Token from "../contracts/Token.cdc"
import DemoToken from "../contracts/DemoToken.cdc"
import MetadataViews from "../../../../../contracts/utility/MetadataViews.cdc"

transaction () {

    prepare(signer: AuthAccount) {

        // Return early if the account already stores a DemoToken Vault
        if signer.borrow<&DemoToken.Vault>(from: DemoToken.VaultStoragePath) != nil {
            return
        }

        // Create a new DemoToken Vault and put it in storage
        signer.save(
            <-DemoToken.createEmptyVault(),
            to: DemoToken.VaultStoragePath
        )

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&DemoToken.Vault{Token.Receiver}>(
            DemoToken.ReceiverPublicPath,
            target: DemoToken.VaultStoragePath
        )

        // Create a public capability to the Vault that exposes the Balance and Resolver interfaces
        signer.link<&DemoToken.Vault{Token.Balance, MetadataViews.Resolver}>(
            DemoToken.VaultPublicPath,
            target: DemoToken.VaultStoragePath
        )
    }
}
