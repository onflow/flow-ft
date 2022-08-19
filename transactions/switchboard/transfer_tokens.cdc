import FungibleToken from "./../../contracts/FungibleToken.cdc"
import ExampleToken from "./../../contracts/ExampleToken.cdc"

// This transaction is a template for a transaction that
// could be used by anyone to send tokens to another account
// through a switchboard, as long as they have set up their
// switchboard and have add the proper capability to it
//
// The address of the receiver account, the amount to transfer
// and the PublicPath for the generic FT receiver will be the
// parameters
transaction(to: Address, amount: UFix64, receiverPath: PublicPath) {

    // The vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        let vaultRef = signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        self.sentVault <- vaultRef.withdraw(amount: amount)
    
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Receiver
        let receiverRef = recipient
            .getCapability(receiverPath)
            .borrow<&{FungibleToken.Receiver}>()
			?? panic("Could not borrow receiver reference to switchboard!")

        // Deposit the withdrawn tokens in the recipient's receiver
        receiverRef.deposit(from: <-self.sentVault)
    
    }

}
