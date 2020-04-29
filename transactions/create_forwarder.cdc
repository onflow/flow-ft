/**

This transaction is a template for a transaction that could be used
to set up an account to forward deposited tokens to another receiver.

If anyone sends tokens to a user's Forwarder Receiver,
the Receiver will just forward those tokens to the Vault that has been
set as the recipient and emit an event that indicates
which user forwarded the tokens.

This way, if an off-chain service wants to monitor who is forwarding
tokens to it, it can watch events to see where the tokens came from.

Steps to set up accounts with token forwarder:

1. The Fungible Token contract interface should already be deployed somewhere
2. The applicable token contract should be deployed.
3. The recipient account should have a Vault for this token created
    and stored in its storage with a published Receiver
4. Deploy the `TokenForwarding.cdc` contract to a different account
5. For a new Account: Create the account normally,
    then run the `create_forwarder.cdc` transaction,
    getting the Receiver from the account that is the recipient.
*/

import FungibleToken from 0x01
import FlowToken from 0x02
import TokenForwarding from 0x03

transaction {

    prepare(signer: AuthAccount) {

        let recipient = getAccount(0x03)

        // Get a Receiver reference for the account that will be the recipient of the forwarded tokens

        let recipientReceiver = recipient
            .getCapability(/public/flowTokenReceiver)!
            .borrow<&{FungibleToken.Receiver}>()
            ?? panic("Could not borrow receiver reference from the capability")

        // Create a new Forwarder resource and store it in the signer's storage
        let forwarder <- TokenForwarding.createNewForwarder(recipient: recipientReceiver)
        signer.save(<-forwarder, to: /storage/flowTokenReceiver)

        // Publish a Receiver capability for the signer, which is linked to the Forwarder
        signer.link<&{FungibleToken.Receiver}>(
            /public/flowTokenReceiver,
            target: /storage/flowTokenReceiver
        )
    }
}
 