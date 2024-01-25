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

import FungibleToken from "FungibleToken"
import ExampleToken from "ExampleToken"
import TokenForwarding from "TokenForwarding"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"

transaction(receiver: Address) {

    prepare(acct: auth(BorrowValue, IssueStorageCapabilityController, PublishCapability, SaveValue, UnpublishCapability) &Account) {

        let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>())
            ?? panic("Could not get vault data view for the contract")
    
        let vaultRef = account.capabilities.borrow<&{FungibleToken.Vault}>(vaultData.metadataPath)
            ?? panic("Could not borrow Balance reference to the Vault")

        // Create the forwarder and save it to the account that is doing the forwarding
        let vault <- TokenForwarding.createNewForwarder(recipient: recipient)
        acct.storage.save(<-vault, to: /storage/exampleTokenForwarder)

        // Unlink the existing capability
        acct.capabilities.unpublish(ExampleToken.ReceiverPublicPath)

        // Link the new forwarding receiver capability
        let tokenReceiverCap = acct.capabilities.storage.issue<&{FungibleToken.Receiver}>(
                /storage/exampleTokenForwarder
            )
        acct.capabilities.publish(tokenReceiverCap, at: ExampleToken.ReceiverPublicPath)

        // Link the new ForwarderPublic capability
        let tokenForwarderCap = acct.capabilities.storage.issue<&{TokenForwarding.ForwarderPublic}>(
                /storage/exampleTokenForwarder
            )
        acct.capabilities.publish(tokenForwarderCap, at: /public/exampleTokenForwarder)
    }
}
