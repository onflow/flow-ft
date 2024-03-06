import "FungibleToken"
import "FungibleTokenMetadataViews"

/// This transaction is what an account would run
/// to set itself up to manage fungible tokens. This function
/// uses views to know where to set up the vault
/// in storage and to create the empty vault.

transaction(contractAddress: Address, contractName: String) {

    prepare(signer: auth(SaveValue, Capabilities) &Account) {
        // Borrow a reference to the vault stored on the passed account at the passed publicPath
        let resolverRef = getAccount(contractAddress)
            .contracts.borrow<&{FungibleToken}>(name: contractName)
            ?? panic("Could not borrow a reference to the fungible token contract")

        // Use that reference to retrieve the FTView 
        let ftVaultData = resolverRef.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve the FTVaultData view for the given Fungible token contract")

        // Create a new empty vault using the createEmptyVault function inside the FTVaultData
        let emptyVault <-ftVaultData.createEmptyVault()

        // Save it to the account
        signer.storage.save(<-emptyVault, to: ftVaultData.storagePath)
        
        // Create a public capability for the vault which includes the .Resolver interface
        let vaultCap = signer.capabilities.storage.issue<&{FungibleToken.Vault}>(ftVaultData.storagePath)
        signer.capabilities.publish(vaultCap, at: ftVaultData.metadataPath)

        // Create a public capability for the vault exposing the receiver interface
        let receiverCap = signer.capabilities.storage.issue<&{FungibleToken.Receiver}>(ftVaultData.storagePath)
        signer.capabilities.publish(receiverCap, at: ftVaultData.receiverPath)

    }
}
 