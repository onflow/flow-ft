/**

# Fungible Token Forwarding Contract

This contract shows how an account could set up a custom Token Receiver
to allow them to forward tokens to a different account whenever they receive tokens.

They can publish this Forwarder resource as a Receiver capability just like a Vault,
and the sender doesn't even need to know it is different.

When an account wants to create a Forwarder, they call the createNewForwarder
function and provide it with the Receiver reference that they want to forward
their tokens to.

*/

import "Token"

pub contract TokenForwarding {

    // Event that is emitted when tokens are deposited to the target receiver
    pub event ForwardedDeposit(amount: UFix64, from: Address?)

    pub resource Forwarder: Token.Receiver {

        // This is where the deposited tokens will be sent.
        // The type indicates that it is a reference to a receiver
        //
        access(self) var recipient: Capability

        // deposit
        //
        // Function that takes a Vault object as an argument and forwards
        // it to the recipient's Vault using the stored reference
        //
        pub fun deposit(from: @Token.Vault) {
            let receiverRef = self.recipient.borrow<&{Token.Receiver}>()!

            let balance = from.balance

            receiverRef.deposit(from: <-from)

            emit ForwardedDeposit(amount: balance, from: self.owner?.address)
        }

        // changeRecipient changes the recipient of the forwarder to the provided recipient
        //
        pub fun changeRecipient(_ newRecipient: Capability) {
            pre {
                newRecipient.borrow<&{Token.Receiver}>() != nil: "Could not borrow Receiver reference from the Capability"
            }
            self.recipient = newRecipient
        }

        /// A getter function that returns the token types supported by this resource,
        /// which can be deposited using the 'deposit' function.
        ///
        /// @return Dictionary of FT types that can be deposited.
        pub fun getSupportedVaultTypes(): {Type: Bool} { 
            if !self.recipient.check<&{Token.Receiver}>() {
                return {}
            }
            let vaultRef = self.recipient.borrow<&{Token.Receiver}>()!
            return {vaultRef.getType(): true}
        }

        init(recipient: Capability) {
            pre {
                recipient.borrow<&{Token.Receiver}>() != nil: "Could not borrow Receiver reference from the Capability"
            }
            self.recipient = recipient
        }
    }

    // createNewForwarder creates a new Forwarder reference with the provided recipient
    //
    pub fun createNewForwarder(recipient: Capability): @Forwarder {
        return <-create Forwarder(recipient: recipient)
    }
}
