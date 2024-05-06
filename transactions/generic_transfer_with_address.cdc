import "FungibleToken"
import "FungibleTokenMetadataViews"

/// Can pass in any contract address and name to transfer a token from that contract
/// This lets you choose the token you want to send
///
/// Any contract can be chosen here, so wallets should check argument values
/// to make sure the intended token contract name and address is passed in
///
transaction(amount: UFix64, to: Address, contractAddress: Address, contractName: String) {

    // The Vault resource that holds the tokens that are being transferred
    let tempVault: @{FungibleToken.Vault}

    // FTVaultData struct to get paths from
    let vaultData: FungibleTokenMetadataViews.FTVaultData

    prepare(signer: auth(BorrowValue) &Account) {

        // Borrow a reference to the vault stored on the passed account at the passed publicPath
        let resolverRef = getAccount(contractAddress)
            .contracts.borrow<&{FungibleToken}>(name: contractName)
            ?? panic("Could not borrow a reference to the fungible token contract")

        // Use that reference to retrieve the FTView 
        self.vaultData = resolverRef.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve the FTVaultData view for the given Fungible token contract")

        // Get a reference to the signer's stored vault
        let vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &{FungibleToken.Provider}>(from: self.vaultData.storagePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        self.tempVault <- vaultRef.withdraw(amount: amount)
    }

    execute {
        let recipient = getAccount(to)
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(self.vaultData.receiverPath)
            ?? panic("Could not borrow reference to the recipient's Receiver!")

        // Transfer tokens from the signer's stored vault to the receiver capability
        receiverRef.deposit(from: <-self.tempVault)
    }
}