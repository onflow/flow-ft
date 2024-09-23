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
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
            ?? panic("The signer does not store an ExampleToken.Vault object at the path "
                    .concat(vaultData.storagePath.toString())
                    .concat("The signer must initialize their account with this vault first!"))

        self.privateForwardingSender = signer.storage.borrow<&PrivateReceiverForwarder.Sender>(from: PrivateReceiverForwarder.SenderStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

    }

    execute {

        for address in addressAmountMap.keys {

            self.privateForwardingSender.sendPrivateTokens(address, tokens: <-self.vaultRef.withdraw(amount: addressAmountMap[address]!))

        }
    }
}
