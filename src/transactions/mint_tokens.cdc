
// This transaction is a template for a transaction that
// could be used by the admin account to mint new tokens
// and deposit them in another account
//
// The minting amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x02
import ExampleToken from 0x03

transaction {
    let tokenAdmin: &ExampleToken.Administrator
    let tokenReceiver: &ExampleToken.Vault{FungibleToken.Receiver}

    prepare(signer: AuthAccount) {
        self.tokenAdmin = signer
        .borrow<&ExampleToken.Administrator>(from: /storage/exampleTokenAdmin) 
        ?? panic("Signer is not the token admin")

        self.tokenReceiver = getAccount(0x04)
        .getCapability(/public/exampleTokenReceiver)!
        .borrow<&ExampleToken.Vault{FungibleToken.Receiver}>()
        ?? panic("Unable to borrow receiver reference")
    }

    execute {
        let minter <- self.tokenAdmin.createNewMinter(allowedAmount: 100.0)
        let mintedVault <- minter.mintTokens(amount: 10)

        self.tokenReceiver.deposit(from: <-mintedVault)

        destroy minter
    }
}
 