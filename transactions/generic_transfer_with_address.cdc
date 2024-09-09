import "FungibleToken"
import "FungibleTokenMetadataViews"

#interaction (
  version: "1.0.0",
	title: "Generic FT Transfer with Contract Address and Name",
	description: "Transfer any Fungible Token by providing the contract address and name",
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
/// @param contractAddress: The address of the contract that defines the tokens being transferred
/// @param contractName: The name of the contract that defines the tokens being transferred. Ex: "FlowToken"
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

        // Get the string representation of the address without the 0x
        var addressString = contractAddress.toString()
        if addressString.length == 18 {
            addressString = addressString.slice(from: 2, upTo: 18)
        }
        let typeString: String = "A.".concat(addressString).concat(".").concat(contractName).concat(".Vault")
        let type = CompositeType(typeString)
        assert(
            type != nil,
            message: "Could not create a type out of the contract name and address!"
        )

        assert(
            self.tempVault.getType() == type!,
            message: "The Vault that was withdrawn to transfer is not the type that was requested!"
        )
    }

    execute {
        let recipient = getAccount(to)
        let receiverRef = recipient.capabilities.borrow<&{FungibleToken.Receiver}>(self.vaultData.receiverPath)
            ?? panic("Could not borrow reference to the recipient's Receiver!")

        // Transfer tokens from the signer's stored vault to the receiver capability
        receiverRef.deposit(from: <-self.tempVault)
    }
}