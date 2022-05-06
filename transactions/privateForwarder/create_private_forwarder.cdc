import FungibleToken from 0xFUNGIBLETOKENADDRESS
import ExampleToken from 0xTOKENADDRESS
import PrivateReceiverForwarder from 0xPRIVATEFORWARDINGADDRESS

// This transaction creates a new private receiver in an account that 
// doesn't already have a private receiver or a public token receiver
// but does already have a Vault

transaction {

    prepare(acct: AuthAccount) {
        receiverCapability = signer.link<&ExampleToken.Vault{FungibleToken.Receiver}>(
            /private/exampleTokenReceiver,
            target: /storage/exampleTokenVault
        )

        let vault <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        acct.save(<-vault, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        signer.link<&{PrivateReceiverForwarder.Forwarder}>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )
    }
}
