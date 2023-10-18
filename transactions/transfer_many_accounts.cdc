import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

/// Transfers tokens to a list of addresses specified in the `addressAmountMap` parameter

transaction(addressAmountMap: {Address: UFix64}) {

    // The Vault resource that holds the tokens that are being transferred
    let vaultRef: auth(FungibleToken.Withdrawable) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {
        // Get a reference to the signer's stored vault
        self.vaultRef = signer.storage.borrow<auth(FungibleToken.Withdrawable) &ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")
    }

    execute {

        for address in addressAmountMap.keys {

            // Withdraw tokens from the signer's stored vault
            let sentVault <- self.vaultRef.withdraw(amount: addressAmountMap[address]!)

            // Get the recipient's public account object
            let recipient = getAccount(address)

            // Get a reference to the recipient's Receiver
            let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(ExampleToken.ReceiverPublicPath)
                ?? panic("Could not borrow receiver reference to the recipient's Vault")

            // Deposit the withdrawn tokens in the recipient's receiver
            receiverRef.deposit(from: <-sentVault)

        }
    }
}
