/**

# Fungible Token Forwarding Contract

This contract shows how an account could set up a custom FungibleToken Receiver
to allow them to forward tokens to a different account whenever they receive tokens.

They can publish this Forwarder resource as a Receiver capability just like a Vault,
and the sender doesn't even need to know it is different.

When an account wants to create a Forwarder, they call the createNewForwarder
function and provide it with the Receiver reference that they want to forward
their tokens to.

*/

import "FungibleToken"

pub contract TokenForwarding {

    // Event that is emitted when tokens are deposited to the target receiver
    pub event ForwardedDeposit(amount: UFix64, from: Address?)

    pub resource interface ForwarderPublic {

        /// Helper function to check whether set `recipient` capability
        /// is not latent or the capability tied to a type is valid.
        pub fun check(): Bool

        /// Gets the receiver assigned to a recipient capability.
        /// This is necessary because without it, it is not possible to look under the hood and see if a capability
        /// is of an expected type or not. This helps guard against infinitely chained TokenForwarding or other invalid 
        /// malicious kinds of updates that could prevent listings from being made that are valid on storefronts.
        ///
        /// @return an optional receiver capability for consumers of the TokenForwarding to check/validate on their own
        pub fun safeBorrow(): &{FungibleToken.Receiver}?
    }

    pub resource Forwarder: FungibleToken.Receiver, ForwarderPublic {

        // This is where the deposited tokens will be sent.
        // The type indicates that it is a reference to a receiver
        //
        access(self) var recipient: Capability

        // deposit
        //
        // Function that takes a Vault object as an argument and forwards
        // it to the recipient's Vault using the stored reference
        //
        pub fun deposit(from: @FungibleToken.Vault) {
            let receiverRef = self.recipient.borrow<&{FungibleToken.Receiver}>()!

            let balance = from.balance

            receiverRef.deposit(from: <-from)

            emit ForwardedDeposit(amount: balance, from: self.owner?.address)
        }

        /// Helper function to check whether set `recipient` capability
        /// is not latent or the capability tied to a type is valid.
        pub fun check(): Bool {
            return self.recipient.check<&{FungibleToken.Receiver}>()
        }

        /// Gets the receiver assigned to a recipient capability.
        /// This is necessary because without it, it is not possible to look under the hood and see if a capability
        /// is of an expected type or not. This helps guard against infinitely chained TokenForwarding or other invalid 
        /// malicious kinds of updates that could prevent listings from being made that are valid on storefronts.
        ///
        /// @return an optional receiver capability for consumers of the TokenForwarding to check/validate on their own
        pub fun safeBorrow(): &{FungibleToken.Receiver}? {
            return self.recipient.borrow<&{FungibleToken.Receiver}>()
        }

        // changeRecipient changes the recipient of the forwarder to the provided recipient
        //
        pub fun changeRecipient(_ newRecipient: Capability) {
            pre {
                newRecipient.borrow<&{FungibleToken.Receiver}>() != nil: "Could not borrow Receiver reference from the Capability"
            }
            self.recipient = newRecipient
        }

        /// A getter function that returns the token types supported by this resource,
        /// which can be deposited using the 'deposit' function.
        ///
        /// @return Array of FT types that can be deposited.
        pub fun getSupportedVaultTypes(): {Type: Bool} {
            if !self.recipient.check<&{FungibleToken.Receiver}>() {
                return {}
            }
            let vaultRef = self.recipient.borrow<&{FungibleToken.Receiver}>()!
            let supportedVaults: {Type: Bool} = {}
            supportedVaults[vaultRef.getType()] = true
            return supportedVaults
        }

        init(recipient: Capability) {
            pre {
                recipient.borrow<&{FungibleToken.Receiver}>() != nil: "Could not borrow Receiver reference from the Capability"
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
