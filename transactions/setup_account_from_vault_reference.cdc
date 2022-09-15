import FungibleToken from "../contracts/FungibleToken.cdc"
import FungibleTokenMetadataViews from "../contracts/FungibleTokenMetadataViews.cdc"
import MetadataViews from "../contracts/utilityContracts/MetadataViews.cdc"

/// This transaction is what an account would run
/// to set itself up to manage fungible tokens. This function
/// uses views to know where to set up the vault
/// in storage and to create the empty vault.

transaction(address: Address, publicPath: PublicPath) {

    prepare(signer: AuthAccount) {
        // Borrow a reference to the vault stored on the passed account at the passed publicPath
        // only caring about it conforming to the MetadataViews.Resolver interface
        let resolverRef = getAccount(address)
            .getCapability(publicPath)
            .borrow<&{MetadataViews.Resolver}>()
            ?? panic("Could not borrow a reference to the vault view resolver")

        // Use that reference to retrieve the FTView 
        let ftView = resolverRef.resolveView(Type<FungibleTokenMetadataViews.FTView>())! as! FungibleTokenMetadataViews.FTView

        // Get the FTVaultData view from from the FTView
        let ftVaultData = ftView.ftVaultData ?? panic ("The stored vault didn't have the vault data view")

        // Create a new empty vault using the createEmptyVault function inside the FTVaultData
        let emptyVault <-ftVaultData.createEmptyVault()

        // Save it to the account
        signer.save(<-emptyVault, to: ftVaultData.storagePath)

        // Create a public capability for the vault exposing the public interfaces
        signer.link<&{FungibleToken.Receiver, FungibleToken.Balance, MetadataViews.Resolver}>(
            ftVaultData.publicPath,
            target: ftVaultData.storagePath
        )

    }
}
 