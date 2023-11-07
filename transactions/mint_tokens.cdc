import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"

/// This transaction is what the minter Account uses to mint new tokens
/// They provide the recipient address and amount to mint, and the tokens
/// are transferred to the address after minting

transaction(recipient: Address, amount: UFix64) {

    /// Reference to the Example Token Minter Resource object
    let tokenMinter: &ExampleToken.Minter

    /// Reference to the Fungible Token Receiver of the recipient
    let tokenReceiver: &{FungibleToken.Vault}

    /// The total supply of tokens before the burn
    let supplyBefore: UFix64

    prepare(signer: auth(BorrowValue) &Account) {
        self.supplyBefore = ExampleToken.totalSupply

        // Borrow a reference to the admin object
        self.tokenMinter = signer.storage.borrow<&ExampleToken.Minter>(from: ExampleToken.AdminStoragePath)
            ?? panic("Signer is not the token admin")

        // Get the account of the recipient and borrow a reference to their receiver
        self.tokenReceiver = getAccount(recipient).capabilities.borrow<&{FungibleToken.Vault}>(
                ExampleToken.VaultPublicPath
            ) ?? panic("Unable to borrow receiver reference")
    }

    execute {

        // Create mint tokens
        let mintedVault <- self.tokenMinter.mintTokens(amount: amount)

        // Deposit them to the receiever
        self.tokenReceiver.deposit(from: <-mintedVault)
    }

    post {
        ExampleToken.totalSupply == self.supplyBefore + amount: "The total supply must be increased by the amount"
    }
}