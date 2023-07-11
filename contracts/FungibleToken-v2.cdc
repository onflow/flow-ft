/**

# The Flow Fungible Token standard

## `FungibleToken` contract

The Fungible Token standard is no longer an interface
that all fungible token contracts would have to conform to.

If a users wants to deploy a new token contract, their contract
does not need to implement the FungibleToken interface, but their tokens
do need to implement the interfaces defined in this contract.

## `Vault` resource interface

Each fungible token resource type needs to implement the `Vault` resource interface.

## `Provider`, `Receiver`, and `Balance` resource interfaces

These interfaces declare pre-conditions and post-conditions that restrict
the execution of the functions in the Vault.

They are separate because it gives the user the ability to share
a reference to their Vault that only exposes the fields functions
in one or more of the interfaces.

It also gives users the ability to make custom resources that implement
these interfaces to do various things with the tokens.
For example, a faucet can be implemented by conforming
to the Provider interface.

*/

import ViewResolver from "./utility/ViewResolver.cdc"

/// FungibleToken
///
/// Fungible Token implementations are no longer required to implement the fungible token
/// interface. We still have it as an interface here because there are some useful
/// utility methods that many projects will still want to have on their contracts,
/// but they are by no means required. all that is required is that the token
/// implements the `Vault` interface
pub contract FungibleToken {

    /// The event that is emitted when tokens are withdrawn from a Vault
    pub event Withdraw(amount: UFix64, from: Address?, type: String)
    access(self) fun emitWithdrawEvent(amount: UFix64, from: Address?, type: String): Bool {
        emit Withdraw(amount: amount, from: from, type: type)
        return true
    }

    /// The event that is emitted when tokens are deposited to a Vault
    pub event Deposit(amount: UFix64, to: Address?, type: String)
    access(self) fun emitDepositEvent(amount: UFix64, to: Address?, type: String): Bool {
        emit Deposit(amount: amount, to: to, type: type)
        return true
    }

    /// The event that is emitted when tokens are transferred from one account to another
    pub event Transfer(amount: UFix64, from: Address?, to: Address?, type: String)
    access(self) fun emitTransferEvent(amount: UFix64, from: Address?, to: Address?, type: String): Bool {
        emit Transfer(amount: amount, from: from, to: to, type: type)
        return true
    }

    /// Event emitted when tokens are destroyed
    pub event Burn(amount: UFix64, type: String)

    access(self) fun emitBurnEvent(amount: UFix64, type: String): Bool {
        if amount >= 0.0 {
            emit Burn(amount: amount, type: type)
        }
        return true
    }

    /// Provider
    ///
    /// The interface that enforces the requirements for withdrawing
    /// tokens from the implementing type.
    ///
    /// It does not enforce requirements on `balance` here,
    /// because it leaves open the possibility of creating custom providers
    /// that do not necessarily need their own balance.
    ///
    pub resource interface Provider {

        /// withdraw subtracts tokens from the owner's Vault
        /// and returns a Vault with the removed tokens.
        ///
        /// The function's access level is public, but this is not a problem
        /// because only the owner storing the resource in their account
        /// can initially call this function.
        ///
        /// The owner may grant other accounts access by creating a private
        /// capability that allows specific other users to access
        /// the provider resource through a reference.
        ///
        /// The owner may also grant all accounts access by creating a public
        /// capability that allows all users to access the provider
        /// resource through a reference.
        ///
        pub fun withdraw(amount: UFix64): @AnyResource{Vault} {
            post {
                // `result` refers to the return value
                result.getBalance() == amount:
                    "Withdrawal amount must be the same as the balance of the withdrawn Vault"
                FungibleToken.emitWithdrawEvent(amount: amount, from: self.owner?.address, type: self.getType().identifier)
            }
        }
    }

    /// Receiver
    ///
    /// The interface that enforces the requirements for depositing
    /// tokens into the implementing type.
    ///
    /// We do not include a condition that checks the balance because
    /// we want to give users the ability to make custom receivers that
    /// can do custom things with the tokens, like split them up and
    /// send them to different places.
    ///
    pub resource interface Receiver {

        /// deposit takes a Vault and deposits it into the implementing resource type
        ///
        pub fun deposit(from: @AnyResource{Vault})

        /// getSupportedVaultTypes optionally returns a list of vault types that this receiver accepts
        pub fun getSupportedVaultTypes(): {Type: Bool} {
            // Below check is implemented to make sure that run-time type would
            // only get returned when the parent resource conforms with `FungibleToken.Vault`. 
            if self.getType().isSubtype(of: Type<@AnyResource{FungibleToken.Vault}>()) {
                return {self.getType(): true}
            } else {
                // Return an empty dictionary as the default value for resource who don't
                // implement `FungibleToken.Vault`, such as `FungibleTokenSwitchboard`, `TokenForwarder` etc.
                return {}
            }
        }

        /// Returns whether or not the given type is accepted by the Receiver
        /// A vault that can accept any type should just return true by default
        pub fun isSupportedVaultType(type: Type): Bool {
            return false
        }
    }

    pub resource interface Transferor {
        /// Function for a direct transfer instead of having to do a deposit and withdrawal
        ///
        pub fun transfer(amount: UFix64, receiver: Capability<&{FungibleToken.Receiver}>) {
            pre {
                receiver.check(): "Could not borrow a reference to the NFT receiver"
            }
        }
    }

    /// Balance
    ///
    /// This interface is now a general purpose metadata interface because
    /// a public interface is needed to get metadata, but adding a whole new interface
    /// for every account to upgrade to is probably too much of a breaking change
    pub resource interface Balance { //: ViewResolver.Resolver {

        /// Method to get the balance
        /// The balance could be a derived field,
        /// so there is no need to require an explicit field
        pub fun getBalance(): UFix64

        pub fun getSupportedVaultTypes(): {Type: Bool}
        pub fun isSupportedVaultType(type: Type): Bool

        /// ViewResolver Methods
        ///
        pub fun getViews(): [Type]
        pub fun resolveView(_ view: Type): AnyStruct?
    }

    /// Vault
    ///
    /// Ideally, this interface would also conform to Receiver, Balance, Transferor, Provider, and Resolver
    /// but that is not supported yet
    ///
    pub resource interface Vault { //: Receiver, Balance, Transferor, Provider, ViewResolver.Resolver {

        /// Get the balance of the vault
        pub fun getBalance(): UFix64

        /// getSupportedVaultTypes optionally returns a list of vault types that this receiver accepts
        pub fun getSupportedVaultTypes(): {Type: Bool}

        pub fun isSupportedVaultType(type: Type): Bool

        /// Returns the storage path where the vault should typically be stored
        pub fun getDefaultStoragePath(): StoragePath? {
            return nil
        }

        /// Returns the public path where this vault should have a public capability
        pub fun getDefaultPublicPath(): PublicPath? {
            return nil
        }

        pub fun getViews(): [Type]
        pub fun resolveView(_ view: Type): AnyStruct?

        /// withdraw subtracts `amount` from the Vault's balance
        /// and returns a new Vault with the subtracted balance
        ///
        pub fun withdraw(amount: UFix64): @AnyResource{Vault} {
            pre {
                self.getBalance() >= amount:
                    "Amount withdrawn must be less than or equal than the balance of the Vault"
            }
            post {
                // use the special function `before` to get the value of the `balance` field
                // at the beginning of the function execution
                //
                self.getBalance() == before(self.getBalance()) - amount:
                    "New Vault balance must be the difference of the previous balance and the withdrawn Vault balance"
            }
        }

        /// deposit takes a Vault and adds its balance to the balance of this Vault
        ///
        pub fun deposit(from: @AnyResource{FungibleToken.Vault}) {
            // Assert that the concrete type of the deposited vault is the same
            // as the vault that is accepting the deposit
            pre {
                from.isInstance(self.getType()): 
                    "Cannot deposit an incompatible token type"
                FungibleToken.emitDepositEvent(amount: from.getBalance(), to: self.owner?.address, type: from.getType().identifier)
            }
            post {
                self.getBalance() == before(self.getBalance()) + before(from.getBalance()):
                    "New Vault balance must be the sum of the previous balance and the deposited Vault"
            }
        }

        /// Function for a direct transfer instead of having to do a deposit and withdrawal
        ///
        pub fun transfer(amount: UFix64, receiver: Capability<&{FungibleToken.Receiver}>) {
            post {
                self.getBalance() == before(self.getBalance()) - amount:
                    "New Vault balance from the sender must be the difference of the previous balance and the withdrawn Vault balance"
                FungibleToken.emitTransferEvent(amount: amount, from: self.owner?.address, to: receiver.borrow()?.owner?.address, type: self.getType().identifier)
            }
        }

        /// createEmptyVault allows any user to create a new Vault that has a zero balance
        ///
        pub fun createEmptyVault(): @AnyResource{Vault} {
            post {
                result.getBalance() == 0.0: "The newly created Vault must have zero balance"
            }
        }

        destroy() {
            pre {
                FungibleToken.emitBurnEvent(amount: self.getBalance(), type: self.getType().identifier)
            }
        }
    }
}