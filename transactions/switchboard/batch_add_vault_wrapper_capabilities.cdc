import "FungibleTokenSwitchboard"
import "ExampleToken"
import "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to add several capabilities that point
/// to fungible token vaults of a different `Type` and belong to a certain `Address`, to their switchboard resource.
///
transaction (address: Address) {

    let vaultPaths: [PublicPath]
    let vaultTypes: [Type]
    let switchboardRef:  auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not resolve FTVaultData view. The ExampleToken"
                .concat(" contract needs to implement the FTVaultData Metadata view in order to execute this transaction."))

        // Store the Example Token receiver's public path in the array of public 
        // paths that will be passed to the switchboard method
        self.vaultPaths = []
        self.vaultPaths.append(vaultData.receiverPath)

        // Store the Example Token's type in the array of types that will be passed 
        // to the switchboard method
        self.vaultTypes = []
        self.vaultTypes.append(Type<@ExampleToken.Vault>())
      
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.storage.borrow<auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath)
			?? panic("The signer does not store a FungibleToken Switchboard object at the path "
                .concat(FungibleTokenSwitchboard.StoragePath.toString())
                .concat(". The signer must initialize their account with this object first!"))
    
    }

    execute {

        // Add the capability(ies) to the switchboard using addNewVaultWrappersByPath
        self.switchboardRef.addNewVaultWrappersByPath(
            paths: self.vaultPaths, 
            types: self.vaultTypes,
            address: address
        )

    }

}