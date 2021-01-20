import FungibleToken from 0xFUNGIBLETOKENADDRESS
import ExampleToken from 0xTOKENADDRESS
import PrivateReceiverForwarder from 0xPRIVATEFORWARDINGADDRESS

transaction(receiver: Address) {

    prepare(acct: AuthAccount) {
        let recipient = getAccount(receiver)
            .getCapability<&{FungibleToken.Receiver}>(/public/exampleTokenReceiver)

        let vault <- PrivateReceiverForwarder.createNewForwarder(recipient: recipient)
        acct.save(<-vault, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        signer.link<&{PrivateReceiverForwarder.Forwarder}>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )
    }
}
