import "FungibleToken"
import "FungibleTokenSwitchboard"
import "ExampleToken"
import "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to add a new vault wrapper capability
/// to their switchboard resource
/// They would just need to change the contract name they are importing from
/// to the token contract that they want to support with their switchboard
///
transaction {

    let tokenReceiverCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
            ?? panic("Could not get vault data view for the contract")

        // Get the token forwarder capability from the signer's account
        self.tokenReceiverCapability = signer.capabilities.get<&{FungibleToken.Receiver}>(
            vaultData.receiverPath)

        // Check if the receiver capability exists
        assert(
            self.tokenReceiverCapability.check(),
            message: "Signer does not have a working fungible token receiver capability"
        )

        // Get a reference to the signers switchboard
        self.switchboardRef = signer.storage.borrow<auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath
            ) ?? panic("Could not borrow reference to switchboard")

    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVaultWrapper(
            capability: self.tokenReceiverCapability,
            type: Type<@ExampleToken.Vault>()
        )

    }

}
