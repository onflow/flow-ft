import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to send tokens to another account
/// through a switchboard, as long as they have set up their switchboard and have add the proper capability to it
///
/// The address of the receiver account, the amount to transfer and the PublicPath for the generic FT receiver will be
/// the parameters
///
transaction(to: Address, amount: UFix64, receiverPath: PublicPath) {

    // The signer's vault to withdraw from
    let sourceVault: auth(FungibleToken.Withdraw) &ExampleToken.Vault

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not get vault data view for the contract")

        // Get a reference to the signer's stored vault
        self.sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
			?? panic("Could not borrow reference to the owner's Vault!")

    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Receiver
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(receiverPath)
			?? panic("Could not borrow receiver reference to switchboard!")

        // Deposit the withdrawn tokens in the recipient's receiver
        receiverRef.deposit(from: <-self.sourceVault.withdraw(amount: amount))

    }

}
