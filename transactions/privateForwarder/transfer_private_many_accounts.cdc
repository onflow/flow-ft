import "FungibleToken"
import "ExampleToken"
import "PrivateReceiverForwarder"
import "FungibleTokenMetadataViews"

/// This transaction transfers to many addresses through their private receivers

transaction(addressAmountMap: {Address: UFix64}) {

    // The Vault resource that holds the tokens that are being transferred
    let vaultRef: auth(FungibleToken.Withdraw) &ExampleToken.Vault

    let privateForwardingSender: &PrivateReceiverForwarder.Sender

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken contract needs to implement the FTVaultData Metadata view in order to execute this transaction.")

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
            ?? panic("The signer does not store an ExampleToken.Vault object at the path \(vaultData.storagePath). The signer must initialize their account with this vault first!")

        self.privateForwardingSender = signer.storage.borrow<&PrivateReceiverForwarder.Sender>(from: PrivateReceiverForwarder.SenderStoragePath)
            ?? panic("The signer does not store a PrivateReceiverForwarder.Sender object at the path \(PrivateReceiverForwarder.SenderStoragePath). The signer must initialize their account with this object first!")

    }

    execute {

        for address in addressAmountMap.keys {

            self.privateForwardingSender.sendPrivateTokens(address, tokens: <-self.vaultRef.withdraw(amount: addressAmountMap[address]!))

        }
    }
}
