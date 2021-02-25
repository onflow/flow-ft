import FungibleToken from 0xFUNGIBLETOKENADDRESS
import ExampleToken from 0xTOKENADDRESS
import PrivateReceiverForwarder from 0xPRIVATEFORWARDINGADDRESS

// This transaction is used to create a user's Flow account with a private forwarder

transaction {
    prepare(payer: AuthAccount) {
        // Pay for the account creation with the Dapper payer
        let acct = AuthAccount(payer: payer)

        // Destroy Dapper user's FLOW token receiver
        acct.unlink(/public/flowTokenReceiver)

        // Create a private receiver
        acct.link<&{FungibleToken.Receiver}>(
            /private/flowTokenReceiver,
            target: /storage/flowTokenVault
        )
        let receiverCapability = acct.getCapability<&{FungibleToken.Receiver}>(/private/flowTokenReceiver)

        // Use the private receiver to create a private forwarder
        let forwarder <- PrivateReceiverForwarder.createNewForwarder(recipient: receiverCapability)

        acct.save(<-forwarder, to: PrivateReceiverForwarder.PrivateReceiverStoragePath)

        acct.link<&PrivateReceiverForwarder.Forwarder>(
            PrivateReceiverForwarder.PrivateReceiverPublicPath,
            target: PrivateReceiverForwarder.PrivateReceiverStoragePath
        )
    }
}