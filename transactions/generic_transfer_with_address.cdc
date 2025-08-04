import "FungibleToken"
import "FungibleTokenMetadataViews"
import "MetadataViews"

#interaction (
  version: "1.1.0",
	title: "Generic FT Transfer with Contract Address and Name",
	description: "Transfer any Fungible Token by providing the vault type identifier",
	language: "en-US",
)

/// Can pass in any contract address and name to transfer a token from that contract
/// This lets you choose the token you want to send
///
/// Any contract can be chosen here, so wallets should check argument values
/// to make sure the intended token contract name and address is passed in
/// Contracts that are used must implement the FTVaultData Metadata View
///
/// Note: This transaction only will work for Fungible Tokens that
///       have their token's resource name set as "Vault".
///       Tokens with other names will need to use a different transaction
///       that additionally specifies the identifier
///
/// @param amount: The amount of tokens to transfer
/// @param to: The address to transfer the tokens to
/// @param ftTypeIdentifier: The type identifier name of the FT type to burn
/// Ex: "A.1654653399040a61.FlowToken.Vault"
///
transaction(amount: UFix64, to: Address, ftTypeIdentifier: String,) {

    // The Vault resource that holds the tokens that are being transferred
    let tempVault: @{FungibleToken.Vault}

    // FTVaultData struct to get paths from
    let vaultData: FungibleTokenMetadataViews.FTVaultData

    prepare(signer: auth(BorrowValue) &Account) {

        self.vaultData = MetadataViews.resolveContractViewFromTypeIdentifier(
            resourceTypeIdentifier: ftTypeIdentifier,
            viewType: Type<FungibleTokenMetadataViews.FTVaultData>()
        ) as? FungibleTokenMetadataViews.FTVaultData
            ?? panic("Could not construct valid FT type and view from identifier \(ftTypeIdentifier)")

        // Get a reference to the signer's stored vault
        let vaultRef = signer.storage.borrow<auth(FungibleToken.Withdraw) &{FungibleToken.Provider}>(from: self.vaultData.storagePath)
			?? panic("The signer does not store a FungibleToken.Provider object at the path "
                .concat(" \(self.vaultData.storagePath.toString())."))

        self.tempVault <- vaultRef.withdraw(amount: amount)

        let type = CompositeType(ftTypeIdentifier)!

        assert(
            self.tempVault.getType() == type,
            message: "The Vault that was withdrawn to transfer is not the type that was requested!"
        )
    }

    execute {
        let recipient = getAccount(to)
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(self.vaultData.receiverPath)
            ?? panic("Could not borrow a Receiver reference to the FungibleToken Vault in account \(to.toString()) at path \(self.vaultData.receiverPath.toString())"
                .concat(". Make sure you are sending to an address that has ")
                .concat("a FungibleToken Vault set up properly at the specified path."))

        // Transfer tokens from the signer's stored vault to the receiver capability
        receiverRef.deposit(from: <-self.tempVault)
    }
}