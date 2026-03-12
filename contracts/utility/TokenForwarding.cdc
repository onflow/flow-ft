/**

# Fungible Token Forwarding Contract

This contract shows how an account could set up a custom FungibleToken Receiver
to allow them to forward tokens to a different account whenever they receive tokens.

They can publish this Forwarder resource as a Receiver capability just like a Vault,
and the sender doesn't even need to know it is different.

When an account wants to create a Forwarder, they call the createNewForwarder
function and provide it with the Receiver capability that they want to forward
their tokens to.

*/

import "FungibleToken"

access(all) contract TokenForwarding {

    access(all) entitlement Owner

    /// Event that is emitted when tokens are deposited to the target receiver
    access(all) event ForwardedDeposit(amount: UFix64, depositedUUID: UInt64, from: Address?, to: Address?, toUUID: UInt64, depositedType: String)

    /// Event that is emitted when the recipient of a forwarder has changed
    access(all) event ForwarderRecipientUpdated(owner: Address?, oldRecipient: Address?, newRecipient: Address?, newReceiverType: String, newReceiverUUID: UInt64)

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

        /// The capability pointing to the receiver that deposited tokens will be forwarded to.
        /// Using a typed capability ensures compile-time safety and avoids requiring explicit
        /// type parameters on every borrow/check call.
        access(self) var recipient: Capability<&{FungibleToken.Receiver}>

        /// Forwards the deposited vault to the recipient receiver.
        /// Emits a ForwardedDeposit event before calling the external deposit
        /// to follow the Checks-Effects-Interactions pattern.
        access(all) fun deposit(from: @{FungibleToken.Vault}) {
            let receiverRef = self.recipient.borrow()
                ?? panic("TokenForwarding.Forwarder.deposit: Could not borrow a Receiver reference from the recipient capability. "
                    .concat("This is likely because the recipient account has removed their Vault or public capability. ")
                    .concat("The owner of this Forwarder should call changeRecipient to update it to a valid receiver."))

            let balance = from.balance
            let uuid = from.uuid

            // Emit before the external call (Checks-Effects-Interactions pattern)
            emit ForwardedDeposit(
                amount: balance,
                depositedUUID: uuid,
                from: self.owner?.address,
                to: receiverRef.owner?.address,
                toUUID: receiverRef.uuid,
                depositedType: from.getType().identifier
            )

            receiverRef.deposit(from: <-from)
        }

        /// Helper function to check whether set `recipient` capability
        /// is not latent or the capability tied to a type is valid.
        access(all) fun check(): Bool {
            return self.recipient.check()
        }

        /// Gets the receiver assigned to a recipient capability.
        /// This is necessary because without it, it is not possible to look under the hood and see if a capability
        /// is of an expected type or not. This helps guard against infinitely chained TokenForwarding or other invalid
        /// malicious kinds of updates that could prevent listings from being made that are valid on storefronts.
        ///
        /// @return an optional receiver capability for consumers of the TokenForwarding to check/validate on their own
        access(all) fun safeBorrow(): &{FungibleToken.Receiver}? {
            return self.recipient.borrow()
        }

        /// Changes the recipient of the forwarder to the provided capability.
        /// The old capability may be stale (e.g. the recipient deleted their vault),
        /// so we use an optional borrow instead of a force-unwrap to avoid permanently
        /// bricking the forwarder in that case.
        access(Owner) fun changeRecipient(_ newRecipient: Capability<&{FungibleToken.Receiver}>) {
            pre {
                newRecipient.borrow() != nil:
                    "TokenForwarding.Forwarder.changeRecipient: Could not borrow a Receiver reference from the new Capability. "
                    .concat("This is likely because the recipient account ")
                    .concat(newRecipient.address.toString())
                    .concat(" has not set up the FungibleToken Vault or public capability correctly. ")
                    .concat("Verify that the address is correct and the account has the correct Vault and capability.")
            }
            let newRef = newRecipient.borrow()!
            emit ForwarderRecipientUpdated(
                owner: self.owner?.address,
                oldRecipient: self.recipient.address,
                newRecipient: newRecipient.address,
                newReceiverType: newRef.getType().identifier,
                newReceiverUUID: newRef.uuid
            )
            self.recipient = newRecipient
        }

        /// A getter function that returns the token types supported by this resource,
        /// which can be deposited using the 'deposit' function.
        /// Delegates to the recipient's own getSupportedVaultTypes() so that chained
        /// forwarders correctly report the underlying vault type rather than their own type.
        ///
        /// @return Dictionary of FT types that can be deposited.
        access(all) view fun getSupportedVaultTypes(): {Type: Bool} {
            // Single borrow eliminates the TOCTOU between a separate check() and borrow()
            if let vaultRef = self.recipient.borrow() {
                return vaultRef.getSupportedVaultTypes()
            }
            return {}
        }

        /// Returns whether or not the given type is accepted by the Receiver.
        /// A vault that can accept any type should just return true by default.
        access(all) view fun isSupportedVaultType(type: Type): Bool {
            return self.getSupportedVaultTypes()[type] ?? false
        }

        init(recipient: Capability<&{FungibleToken.Receiver}>) {
            pre {
                recipient.borrow() != nil:
                    "TokenForwarding.Forwarder.init: Could not borrow a Receiver reference from the Capability. "
                    .concat("This is likely because the recipient account ")
                    .concat(recipient.address.toString())
                    .concat(" has not set up the FungibleToken Vault or public capability correctly. ")
                    .concat("Verify that the address is correct and the account has the correct Vault and capability. ")
            }
            self.recipient = recipient
        }
    }

    /// Creates a new Forwarder that will forward all deposited tokens to the given recipient.
    access(all) fun createNewForwarder(recipient: Capability<&{FungibleToken.Receiver}>): @Forwarder {
        return <-create Forwarder(recipient: recipient)
    }
}
