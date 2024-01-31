import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

/// This transaction is a template for a transaction that could be used by anyone to add a new vault wrapper capability
/// to their switchboard resource
///
transaction {

    let tokenForwarderCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())
            ?? panic("Could not get vault data view for the contract")

        // Get the token forwarder capability from the signer's account
        self.tokenForwarderCapability = signer.capabilities.get<&{FungibleToken.Receiver}>(
                vaultData.receiverPath
            )

        // Check if the receiver capability exists
        assert(
            self.tokenForwarderCapability.check(),
            message: "Signer does not have a working fungible token receiver capability"
        )

        // Get a reference to the signers switchboard
        self.switchboardRef = signer.storage.borrow<&FungibleTokenSwitchboard.Switchboard>(
                from: FungibleTokenSwitchboard.StoragePath
            ) ?? panic("Could not borrow reference to switchboard")

    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVaultWrapper(
            capability: self.tokenForwarderCapability,
            type: Type<@ExampleToken.Vault>()
        )

    }

}
