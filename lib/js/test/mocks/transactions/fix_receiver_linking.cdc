import "FungibleToken"
import "ExampleToke"
import "FungibleTokenSwitchboard"

transaction() {

    prepare(signer: AuthAccount) {

        signer.unlink(FungibleTokenSwitchboard.ReceiverPublicPath)

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&ExampleToken.Vault{FungibleToken.Receiver}>(
            ExampleToken.ReceiverPublicPath,
            target: ExampleToken.VaultStoragePath
        )
    }
}