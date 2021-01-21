import FungibleToken from 0xFUNGIBLETOKENADDRESS
import ExampleToken from 0xTOKENADDRESS
import PrivateReceiverForwarder from 0xPRIVATEFORWARDINGADDRESS

transaction(amount: UFix64, to: [Address]) {

    // The Vault resource that holds the tokens that are being transferred
    let vaultRef: &ExampleToken.Vault

    let privateForwardingSender: &PrivateReceiverForwarder.Sender

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.borrow<&ExampleToken.Vault>(from: /storage/exampleTokenVault)
			?? panic("Could not borrow reference to the owner's Vault!")

        self.privateForwardingSender = signer.borrow<&PrivateReceiverForwarder.Sender>(from: PrivateReceiverForwarder.SenderStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

    }

    execute {

        for address in to {

            self.privateForwardingSender.sendPrivateTokens(address, tokens: <-self.vaultRef.withdraw(amount: amount))

        }
    }
}
