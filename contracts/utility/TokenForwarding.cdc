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

access(all) contract TokenForwarding {

    // Event that is emitted when tokens are deposited to the target receiver
    access(all) event ForwardedDeposit(amount: UFix64, depositedUUID: UInt64, from: Address?, to: Address?, toUUID: UInt64, depositedType: Type)

    access(all) resource interface ForwarderPublic {

        /// Helper function to check whether set `recipient` capability
        /// is not latent or the capability tied to a type is valid.
        access(all) fun check(): Bool

        /// Gets the receiver assigned to a recipient capability.
        /// This is necessary because without it, it is not possible to look under the hood and see if a capability
        /// is of an expected type or not. This helps guard against infinitely chained TokenForwarding or other invalid 
        /// malicious kinds of updates that could prevent listings from being made that are valid on storefronts.
        ///
        /// @return an optional receiver capability for consumers of the TokenForwarding to check/validate on their own
        access(all) fun safeBorrow(): &{FungibleToken.Receiver}?
    }

    access(all) resource Forwarder: FungibleToken.Receiver, ForwarderPublic {

        // This is where the deposited tokens will be sent.
        // The type indicates that it is a reference to a receiver
        //
        access(self) var recipient: Capability

        // deposit
        //
        // Function that takes a Vault object as an argument and forwards
        // it to the recipient's Vault using the stored reference
        //
        access(all) fun deposit(from: @{FungibleToken.Vault}) {
            let receiverRef = self.recipient.borrow<&{FungibleToken.Receiver}>()!

            let balance = from.balance

            let uuid = from.uuid

            emit ForwardedDeposit(amount: balance, depositedUUID: uuid, from: self.owner?.address, to: receiverRef.owner?.address, toUUID: receiverRef.uuid, depositedType: from.getType())

            receiverRef.deposit(from: <-from)
        }

        /// Helper function to check whether set `recipient` capability
        /// is not latent or the capability tied to a type is valid.
        access(all) fun check(): Bool {
            return self.recipient.check<&{FungibleToken.Receiver}>()
        }

        /// Gets the receiver assigned to a recipient capability.
        /// This is necessary because without it, it is not possible to look under the hood and see if a capability
        /// is of an expected type or not. This helps guard against infinitely chained TokenForwarding or other invalid 
        /// malicious kinds of updates that could prevent listings from being made that are valid on storefronts.
        ///
        /// @return an optional receiver capability for consumers of the TokenForwarding to check/validate on their own
        access(all) fun safeBorrow(): &{FungibleToken.Receiver}? {
            return self.recipient.borrow<&{FungibleToken.Receiver}>()
        }

        // changeRecipient changes the recipient of the forwarder to the provided recipient
        //
        access(all) fun changeRecipient(_ newRecipient: Capability) {
            pre {
                newRecipient.borrow<&{FungibleToken.Receiver}>() != nil: "Could not borrow Receiver reference from the Capability"
            }
            self.recipient = newRecipient
        }

        /// A getter function that returns the token types supported by this resource,
        /// which can be deposited using the 'deposit' function.
        ///
        /// @return Array of FT types that can be deposited.
        access(all) view fun getSupportedVaultTypes(): {Type: Bool} {
            if !self.recipient.check<&{FungibleToken.Receiver}>() {
                return {}
            }
            let vaultRef = self.recipient.borrow<&{FungibleToken.Receiver}>()!
            let supportedVaults: {Type: Bool} = {}
            supportedVaults[vaultRef.getType()] = true
            return supportedVaults
        }

        /// Returns whether or not the given type is accepted by the Receiver
        /// A vault that can accept any type should just return true by default
        access(all) view fun isSupportedVaultType(type: Type): Bool {
            let supportedVaults = self.getSupportedVaultTypes()
            if let supported = supportedVaults[type] {
                return supported
            } else { return false }
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
    access(all) fun createNewForwarder(recipient: Capability): @Forwarder {
        return <-create Forwarder(recipient: recipient)
    }
}
