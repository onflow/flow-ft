/*

# Fungible Token Private Receiver Contract

This contract implements a special resource and receiver interface
whose deposit function is only callable by an admin through a public capability.

*/

import "FungibleToken"

access(all) contract PrivateReceiverForwarder {

    /// Event that is emitted when tokens are deposited to the target receiver
    access(all) event PrivateDeposit(amount: UFix64, depositedUUID: UInt64, from: Address?, to: Address?, toUUID: UInt64)

    access(all) let SenderStoragePath: StoragePath

    access(all) let PrivateReceiverStoragePath: StoragePath
    access(all) let PrivateReceiverPublicPath: PublicPath

    access(all) resource Forwarder {

        /// The capability pointing to the receiver that deposited tokens will be forwarded to.
        /// Immutable: there is no changeRecipient function, so the recipient is fixed at creation.
        access(self) let recipient: Capability<&{FungibleToken.Receiver}>

        /// Forwards the deposited vault to the recipient receiver.
        /// Only callable from within this contract (via the Sender resource).
        /// Emits PrivateDeposit before calling the external deposit to follow
        /// the Checks-Effects-Interactions pattern.
        access(contract) fun deposit(from: @{FungibleToken.Vault}) {
            let receiverRef = self.recipient.borrow()
                ?? panic("PrivateReceiverForwarder.Forwarder.deposit: Could not borrow a Receiver reference from the recipient capability. "
                    .concat("The recipient needs to have the correct Fungible Token Vault initialized in their account with a public Receiver Capability."))

            let balance = from.balance
            let uuid = from.uuid

            // Emit before the external call (Checks-Effects-Interactions pattern)
            emit PrivateDeposit(
                amount: balance,
                depositedUUID: uuid,
                from: self.owner?.address,
                to: receiverRef.owner?.address,
                toUUID: receiverRef.uuid
            )

            receiverRef.deposit(from: <-from)
        }

        init(recipient: Capability<&{FungibleToken.Receiver}>) {
            pre {
                recipient.borrow() != nil:
                    "PrivateReceiverForwarder.Forwarder.init: Could not borrow a Receiver reference from the recipient Capability. "
                    .concat("The recipient needs to have the correct Fungible Token Vault initialized in their account with a public Receiver Capability.")
            }
            self.recipient = recipient
        }
    }

    /// Creates a new Forwarder that will accept private token deposits and forward them to the given recipient.
    access(all) fun createNewForwarder(recipient: Capability<&{FungibleToken.Receiver}>): @Forwarder {
        return <-create Forwarder(recipient: recipient)
    }


    access(all) resource Sender {
        /// Sends tokens to the PrivateReceiverForwarder at the given address.
        /// The recipient account must have a Forwarder stored at PrivateReceiverPublicPath.
        access(all) fun sendPrivateTokens(_ address: Address, tokens: @{FungibleToken.Vault}) {

            let account = getAccount(address)

            let privateReceiver = account.capabilities.borrow<&PrivateReceiverForwarder.Forwarder>(
                    PrivateReceiverForwarder.PrivateReceiverPublicPath
                ) ?? panic("PrivateReceiverForwarder.Sender.sendPrivateTokens: Could not borrow a reference to the private forwarder in the account "
                            .concat(address.toString())
                            .concat(". Make sure this account has a Forwarder initialized in its storage with a public capability at ")
                            .concat(PrivateReceiverForwarder.PrivateReceiverPublicPath.toString()))

            privateReceiver.deposit(from: <-tokens)
        }
    }

    init(senderPath: StoragePath, storagePath: StoragePath, publicPath: PublicPath) {

        self.SenderStoragePath = senderPath

        self.PrivateReceiverStoragePath = storagePath
        self.PrivateReceiverPublicPath = publicPath

        self.account.storage.save(<-create Sender(), to: self.SenderStoragePath)

    }
}
