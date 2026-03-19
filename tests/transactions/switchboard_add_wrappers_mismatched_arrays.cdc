import "FungibleTokenSwitchboard"
import "ExampleToken"

/// Test helper transaction that calls addNewVaultWrappersByPath with two paths
/// but only one type, triggering the paths.length == types.length pre-condition.
transaction(address: Address) {
    let switchboardRef: auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard

    prepare(signer: auth(BorrowValue) &Account) {
        self.switchboardRef = signer.storage.borrow<auth(FungibleTokenSwitchboard.Owner) &FungibleTokenSwitchboard.Switchboard>(
            from: FungibleTokenSwitchboard.StoragePath)
            ?? panic("No Switchboard found at storage path")
    }

    execute {
        // Two paths, one type — lengths are mismatched
        self.switchboardRef.addNewVaultWrappersByPath(
            paths: [/public/exampleTokenReceiver, /public/exampleTokenVault],
            types: [Type<@ExampleToken.Vault>()],
            address: address
        )
    }
}
