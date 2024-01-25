/*

# Fungible Token Private Receiver Contract

This contract implements a special resource and receiver interface 
whose deposit function is only callable by an admin through a public capability.

*/

import FungibleToken from "FungibleToken"

access(all) contract PrivateReceiverForwarder {

    // Event that is emitted when tokens are deposited to the target receiver
    access(all) event PrivateDeposit(amount: UFix64, depositedUUID: UInt64, from: Address?, to: Address?, toUUID: UInt64)

    access(all) let SenderStoragePath: StoragePath

    access(all) let PrivateReceiverStoragePath: StoragePath
    access(all) let PrivateReceiverPublicPath: PublicPath

    access(all) resource Forwarder {

        // This is where the deposited tokens will be sent.
        // The type indicates that it is a reference to a receiver
        //
        access(self) var recipient: Capability<&{FungibleToken.Receiver}>

        // deposit
        //
        // Function that takes a Vault object as an argument and forwards
        // it to the recipient's Vault using the stored reference
        //
        access(contract) fun deposit(from: @{FungibleToken.Vault}) {
            let receiverRef = self.recipient.borrow()!

            let balance = from.balance

            let uuid = from.uuid

            receiverRef.deposit(from: <-from)

            emit PrivateDeposit(amount: balance, depositedUUID: uuid, from: self.owner?.address, to: receiverRef.owner?.address, toUUID: receiverRef.uuid)
        }

        init(recipient: Capability<&{FungibleToken.Receiver}>) {
            pre {
                recipient.borrow() != nil: "Could not borrow Receiver reference from the Capability"
            }
            self.recipient = recipient
        }
    }

    // createNewForwarder creates a new Forwarder reference with the provided recipient
    //
    access(all) fun createNewForwarder(recipient: Capability<&{FungibleToken.Receiver}>): @Forwarder {
        return <-create Forwarder(recipient: recipient)
    }


    access(all) resource Sender {
        access(all) fun sendPrivateTokens(_ address: Address, tokens: @{FungibleToken.Vault}) {

            let account = getAccount(address)

            let privateReceiver = account.capabilities.borrow<&PrivateReceiverForwarder.Forwarder>(
                    PrivateReceiverForwarder.PrivateReceiverPublicPath
                ) ?? panic("Could not borrow reference to private forwarder")

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