import FungibleToken from "../contracts/FungibleToken.cdc"
import FungibleTokenMetadataViews from "../contracts/FungibleTokenMetadataViews.cdc"
import MetadataViews from "../contracts/utilityContracts/MetadataViews.cdc"

/// This transaction is what an account would run
/// to set itself up to manage fungible tokens. This function
/// uses views to know where to set up the vault
/// in storage and to create the empty vault.

transaction(address: Address, publicPath: PublicPath) {

    prepare(signer: AuthAccount) {
        let resolverRef = getAccount(address)
            .getCapability(publicPath)
            .borrow<&{MetadataViews.Resolver}>()
            ?? panic("Could not borrow a reference to the vault view resolver")

        let ftVaultData = resolverRef.resolveView(Type<FungibleTokenMetadataViews.FTVaultData>())! as! FungibleTokenMetadataViews.FTVaultData

        // Create a new empty vault
        let emptyVault <-ftVaultData.createEmptyVault()

        // Save it to the account
        signer.save(<-emptyVault, to: ftVaultData.storagePath)

        // create a public capability for the vault
        signer.link<&{FungibleToken.Receiver, FungibleToken.Balance, MetadataViews.Resolver}>(
            ftVaultData.publicPath,
            target: ftVaultData.storagePath
        )

    }
}
 