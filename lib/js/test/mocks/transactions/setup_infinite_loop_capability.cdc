import "FungibleToken"
import "ExampleToken"
import "FungibleTokenSwitchboard"

transaction() {

    prepare(signer: AuthAccount) {

        signer.unlink(ExampleToken.ReceiverPublicPath)

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&ExampleToken.Vault{FungibleToken.Receiver}>(
            FungibleTokenSwitchboard.ReceiverPublicPath,
            target: ExampleToken.VaultStoragePath
        )
    }
}