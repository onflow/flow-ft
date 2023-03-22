import FungibleToken from "../../../../../contracts/FungibleToken.cdc"
import ExampleToken from "../../../../../contracts/ExampleToken.cdc"
import FungibleTokenSwitchboard from "../../../../../contracts/FungibleTokenSwitchboard.cdc"

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