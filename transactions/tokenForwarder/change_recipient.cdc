/**

This transaction changes the recipient for a token forwarder recource
to a different account

*/

import "FungibleToken"
import "ExampleToken"
import "TokenForwarding"

transaction(newRecipient: Address) {

    prepare(signer: auth(BorrowValue) &Account) {

        // Get the receiver capability for the account being forwarded to
        let recipient = getAccount(newRecipient).capabilities.get<&{FungibleToken.Receiver}>(ExampleToken.ReceiverPublicPath)

        // Get a reference to the signer's forwarder
        let forwarderRef = signer.storage.borrow<auth(TokenForwarding.Owner) &TokenForwarding.Forwarder>(from: /storage/exampleTokenForwarder)
			?? panic("Could not borrow reference to the owner's forwarder!")

        forwarderRef.changeRecipient(recipient)
    }
}
