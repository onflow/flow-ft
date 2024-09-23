import "FungibleToken"
import "ExampleToken"
import "FungibleTokenMetadataViews"

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
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction"))

        // Get a reference to the signer's stored vault
        self.sourceVault = signer.storage.borrow<auth(FungibleToken.Withdraw) &ExampleToken.Vault>(from: vaultData.storagePath)
			?? panic("The signer does not store a ExampleToken Vault object at the path "
                .concat(vaultData.storagePath.toString()).concat("For the ExampleToken contract. ")
                .concat("The signer must initialize their account with this object first!"))

    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Receiver
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(receiverPath)
            ?? panic("Could not borrow a Receiver reference to the FungibleToken Vault in account "
                .concat(to.toString()).concat(" at path ").concat(receiverPath.toString())
                .concat(". Make sure you are querying an address that has ")
                .concat("a FungibleToken Vault set up properly at the specified path."))

        // Deposit the withdrawn tokens in the recipient's receiver
        receiverRef.deposit(from: <-self.sourceVault.withdraw(amount: amount))

    }

}
