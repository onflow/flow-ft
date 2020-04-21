/**

# The Custodial Deposit example contract

An example contract for accepting deposits from any user
so that a custodial service could credit to their account.

This is an example implementation of a Custodial receiver contract
that a custodial account with an off-chain app can use to accept
deposits from external accounts and credit them to the relevant
user account off-chain

*/

import FungibleToken from 0x01
import FlowToken from 0x02

pub contract CustodialDeposit {

    // Event that is emitted when the contract is created
    pub event Initialized(initialSupply: UFix64)

    // Event that is emitted when tokens are withdrawn from a Vault
    pub event Withdraw(amount: UFix64)

    // Event that is emitted when tokens are deposited to a Vault
    pub event TaggedDeposit(amount: UFix64, tag: String)

    pub resource interface DepositPublic {
        pub fun taggedDeposit(from: @FungibleToken.Vault, tag: String)
    }

    pub resource DepositResource: FungibleToken.Provider, DepositPublic {

        // Field to hold the vault that holds the tokens that people deposit
        pub let vault: @FlowToken.Vault

        // taggedDeposit
        //
        // Function that deposits a Vault into the stored Vault
        // and emits an event with the tag for the user
        //
        pub fun taggedDeposit(from: @FungibleToken.Vault, tag: String) {
            emit TaggedDeposit(amount: from.balance, tag: tag)
            self.vault.deposit(from: <-from)
        }

        // withdraw
        //
        // Function that takes an integer amount as an argument
        // and withdraws that amount from the Vault.
        // It creates a new temporary Vault that is used to hold
        // the money that is being transferred. It returns the newly
        // created Vault to the context that called so it can be deposited
        // elsewhere.
        //
        pub fun withdraw(amount: UFix64): @FungibleToken.Vault {
            emit Withdraw(amount: amount)
            return <-self.vault.withdraw(amount: amount)
        }

        // getBalance
        //
        // returns the balance of the stored Vault
        pub fun getBalance(): UFix64 {
            return self.vault.balance
        }

        init(initVault: @FlowToken.Vault) {
            self.vault <- initVault
        }

        destroy() {
            destroy self.vault
        }
    }

    access(self) fun initializeDepositResource(vault: @FlowToken.Vault) {
        let balance = vault.balance

        self.account.save(
            <-create DepositResource(initVault: <-vault),
            to: /storage/depositResource
        )

        self.account.link<&{DepositPublic}>(
            /public/depositResourcePublic,
            target: /storage/depositResource
        )

        emit Initialized(initialSupply: balance)
    }

    // The init function creates a new instance of the DepositResource resource
    // and stores it in account storage, whether or not a Vault resource
    // already exists in storage
    init() {

        // To initialize the DepositResource,
        // take the stored Vault out of storage and use it,
        // or if no vault is stored, create a new empty vault

        if let storedVault <- self.account.load<@FlowToken.Vault>(from: /storage/flowTokenVault) {
            self.initializeDepositResource(vault: <-storedVault)
        } else {
            let newEmptyVault <-FlowToken.createEmptyVault()
            self.initializeDepositResource(vault: <-newEmptyVault)
        }
    }
}
