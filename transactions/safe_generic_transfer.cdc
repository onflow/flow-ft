import "FungibleToken"

/// Can pass in any storage path and receiver path instead of just the default.
/// This lets you choose the token you want to send as well the capability you want to send it to.
///
/// Any token path can be passed as an argument here, so wallets should
/// should check argument values to make sure the intended token path is passed in
///
transaction(amount: UFix64, to: Address, senderPath: StoragePath, receiverPath: PublicPath) {

    // The Vault resource that holds the tokens that are being transferred
    let tempVault: @FungibleToken.Vault

    // Borrowed teference receive tokens if receiving account doesn't support the sending token
    let senderReceiverRef: &{Token.Receiver}

    prepare(signer: auth(BorrowValue) &Account) {

        // Get a reference to the signer's stored vault
        let vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &{FungibleToken.Provider}>(from: senderPath)
			?? panic("Could not borrow reference to the owner's Vault!")
        
        self.senderReceiverRef = signer.storage.borrow<&{FungibleToken.Receiver}>(from: senderPath)
			?? panic("Could not borrow {FungibleToken.Receiver} reference to the owner's Vault!")

        self.tempVault <- vaultRef.withdraw(amount: amount)
    }

    execute {
        let receiverRef = getAccount(to).capabilities.borrow<&{FungibleToken.Receiver}>(receiverPath)!
        let supportedVaultTypes = receiverRef.getSupportedVaultTypes()
        // Only transfer tokens when the receiver is willing to receive the targeted FT.
        if supportedVaultTypes.containsKey(self.tempVault.getType()) {
            // Transfer tokens from the signer's stored vault to the receiver capability
            receiverRef.deposit(from: <-self.tempVault)
        } else {
            self.senderReceiverRef.deposit(from: <-self.tempVault)
        }
    }
}