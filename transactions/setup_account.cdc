// This transaction is a template for a transaction to allow
// anyone to add a Vault resource to their account so that
// they can use the exampleToken

import "FungibleToken"
import "ExampleToken"
import "ViewResolver"
import "FungibleTokenMetadataViews"

transaction () {

    prepare(signer: auth(BorrowValue, IssueStorageCapabilityController, PublishCapability, SaveValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction."))

        // Return early if the account already stores a ExampleToken Vault
        if signer.storage.borrow<&ExampleToken.Vault>(from: vaultData.storagePath) != nil {
            return
        }

        let vault <- ExampleToken.createEmptyVault(vaultType: Type<@ExampleToken.Vault>())

        // Create a new ExampleToken Vault and put it in storage
        signer.storage.save(<-vault, to: vaultData.storagePath)

        // Create a public capability to the Vault that exposes the Vault interfaces
        let vaultCap = signer.capabilities.storage.issue<&ExampleToken.Vault>(
            vaultData.storagePath
        )
        signer.capabilities.publish(vaultCap, at: vaultData.metadataPath)

        // Create a public Capability to the Vault's Receiver functionality
        let receiverCap = signer.capabilities.storage.issue<&ExampleToken.Vault>(
            vaultData.storagePath
        )
        signer.capabilities.publish(receiverCap, at: vaultData.receiverPath)
    }
}
