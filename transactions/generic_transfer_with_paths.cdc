import "FungibleToken"

/// Can pass in any storage path and receiver path identifier instead of just the default.
/// This lets you choose the token you want to send as well the capability you want to send it to.
///
/// Any token path can be passed as an argument here, so wallets should
/// should check argument values to make sure the intended token path is passed in
///
transaction(amount: UFix64, to: Address, senderPathIdentifier: String, receiverPathIdentifier: String) {

    // The Vault resource that holds the tokens that are being transferred
    let tempVault: @{FungibleToken.Vault}

    prepare(signer: auth(BorrowValue) &Account) {

        let storagePath = StoragePath(identifier: senderPathIdentifier)
            ?? panic("Could not construct a storage path from the provided path identifier string")

        // Get a reference to the signer's stored vault
        let vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &{FungibleToken.Provider}>(from: storagePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        self.tempVault <- vaultRef.withdraw(amount: amount)
    }

    execute {
        let publicPath = PublicPath(identifier: receiverPathIdentifier)
            ?? panic("Could not construct a public path from the provided path identifier string")

        let recipient = getAccount(to)
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(publicPath)
            ?? panic("Could not borrow reference to the recipient's Receiver!")

        // Transfer tokens from the signer's stored vault to the receiver capability
        receiverRef.deposit(from: <-self.tempVault)
    }
}