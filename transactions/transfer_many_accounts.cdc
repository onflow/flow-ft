import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"

/// Transfers tokens to a list of addresses specified in the `addressAmountMap` parameter

transaction(addressAmountMap: {Address: UFix64}) {

    /// FTVaultData metadata view for the token being used
    let vaultData: FungibleTokenMetadataViews.FTVaultData

    /// The Vault resource that holds the tokens that are being transferred
    let vaultRef: auth(FungibleToken.Withdraw) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {
        self.vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: self.vaultData.storagePath)
            ?? panic("The signer does not store an ExampleToken.Vault object at the path "
                    .concat(self.vaultData.storagePath.toString())
                    .concat("The signer must initialize their account with this vault first!"))
    }

    execute {

        for address in addressAmountMap.keys {

            // Withdraw tokens from the signer's stored vault
            let sentVault <- self.vaultRef.withdraw(amount: addressAmountMap[address]!)

            // Get the recipient's public account object
            let recipient = getAccount(address)

            // Get a reference to the recipient's Receiver
            let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(self.vaultData.receiverPath)
                ?? panic("Could not borrow a Receiver reference to the FungibleToken Vault in account "
                    .concat(address.toString()).concat(" at path ").concat(self.vaultData.receiverPath.toString())
                    .concat(". Make sure you are sending to an address that has ")
                    .concat("a FungibleToken Vault set up properly at the specified path."))

            // Deposit the withdrawn tokens in the recipient's receiver
            receiverRef.deposit(from: <-sentVault)

        }
    }
}
