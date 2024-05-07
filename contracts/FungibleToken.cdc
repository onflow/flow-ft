/**

# The Flow Fungible Token standard

## `FungibleToken` contract

If a users wants to deploy a new token contract, their contract
needs to implement the FungibleToken interface and their tokens
need to implement the interfaces defined in this contract.

/// Contributors (please add to this list if you contribute!):
/// - Joshua Hannan - https://github.com/joshuahannan
/// - Bastian Müller - https://twitter.com/turbolent
/// - Dete Shirley - https://twitter.com/dete73
/// - Bjarte Karlsen - https://twitter.com/0xBjartek
/// - Austin Kline - https://twitter.com/austin_flowty
/// - Giovanni Sanchez - https://twitter.com/gio_incognito
/// - Deniz Edincik - https://twitter.com/bluesign
/// - Jonny - https://github.com/dryruner
///
/// Repo reference: https://github.com/onflow/flow-ft

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

import "ViewResolver"
import "Burner"

/// FungibleToken
///
/// Fungible Token implementations should implement the fungible token
/// interface.
access(all) contract interface FungibleToken: ViewResolver {

    // An entitlement for allowing the withdrawal of tokens from a Vault
    access(all) entitlement Withdraw

    /// The event that is emitted when tokens are withdrawn from a Vault
    access(all) event Withdrawn(type: String, amount: UFix64, from: Address?, fromUUID: UInt64, withdrawnUUID: UInt64, balanceAfter: UFix64)

    /// The event that is emitted when tokens are deposited to a Vault
    access(all) event Deposited(type: String, amount: UFix64, to: Address?, toUUID: UInt64, depositedUUID: UInt64, balanceAfter: UFix64)

    /// Event that is emitted when the global burn method is called with a non-zero balance
    access(all) event Burned(type: String, amount: UFix64, fromUUID: UInt64)

    /// Balance
    ///
    /// The interface that provides standard functions\
    /// for getting balance information
    ///
    access(all) resource interface Balance {
        access(all) var balance: UFix64
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
    access(all) resource interface Provider {

        /// Function to ask a provider if a specific amount of tokens
        /// is available to be withdrawn
        /// This could be useful to avoid panicing when calling withdraw
        /// when the balance is unknown
        /// Additionally, if the provider is pulling from multiple vaults
        /// it only needs to check some of the vaults until the desired amount
        /// is reached, potentially helping with performance.
        /// 
        access(all) view fun isAvailableToWithdraw(amount: UFix64): Bool

        /// withdraw subtracts tokens from the implementing resource
        /// and returns a Vault with the removed tokens.
        ///
        /// The function's access level is `access(Withdraw)`
        /// So in order to access it, one would either need the object itself
        /// or an entitled reference with `Withdraw`.
        ///
        access(Withdraw) fun withdraw(amount: UFix64): @{Vault} {
            post {
                // `result` refers to the return value
                result.balance == amount:
                    "Withdrawal amount must be the same as the balance of the withdrawn Vault"
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
    access(all) resource interface Receiver {

        /// deposit takes a Vault and deposits it into the implementing resource type
        ///
        access(all) fun deposit(from: @{Vault})

        /// getSupportedVaultTypes returns a dictionary of Vault types
        /// and whether the type is currently supported by this Receiver
        access(all) view fun getSupportedVaultTypes(): {Type: Bool}

        /// Returns whether or not the given type is accepted by the Receiver
        /// A vault that can accept any type should just return true by default
        access(all) view fun isSupportedVaultType(type: Type): Bool
    }

    /// Vault
    ///
    /// Ideally, this interface would also conform to Receiver, Balance, Transferor, Provider, and Resolver
    /// but that is not supported yet
    ///
    access(all) resource interface Vault: Receiver, Provider, Balance, ViewResolver.Resolver, Burner.Burnable {

        /// Field that tracks the balance of a vault
        access(all) var balance: UFix64

        /// Called when a fungible token is burned via the `Burner.burn()` method
        /// Implementations can do any bookkeeping or emit any events
        /// that should be emitted when a vault is destroyed.
        /// Many implementations will want to update the token's total supply
        /// to reflect that the tokens have been burned and removed from the supply.
        /// Implementations also need to set the balance to zero before the end of the function
        /// This is to prevent vault owners from spamming fake Burned events.
        access(contract) fun burnCallback() {
            pre {
                emit Burned(type: self.getType().identifier, amount: self.balance, fromUUID: self.uuid)
            }
            post {
                self.balance == 0.0: "The balance must be set to zero during the burnCallback method so that it cannot be spammed"
            }
            self.balance = 0.0
        }

        /// getSupportedVaultTypes returns a dictionary of vault types and whether this receiver accepts the indexed type
        /// The default implementation is included here because vaults are expected
        /// to only accepted their own type, so they have no need to provide an implementation
        /// for this function
        access(all) view fun getSupportedVaultTypes(): {Type: Bool} {
            // Below check is implemented to make sure that run-time type would
            // only get returned when the parent resource conforms with `FungibleToken.Vault`. 
            if self.getType().isSubtype(of: Type<@{FungibleToken.Vault}>()) {
                return {self.getType(): true}
            } else {
                // Return an empty dictionary as the default value for resource who don't
                // implement `FungibleToken.Vault`, such as `FungibleTokenSwitchboard`, `TokenForwarder` etc.
                return {}
            }
        }

        /// Checks if the given type is supported by this Vault
        access(all) view fun isSupportedVaultType(type: Type): Bool {
            return self.getSupportedVaultTypes()[type] ?? false
        }

        /// withdraw subtracts `amount` from the Vault's balance
        /// and returns a new Vault with the subtracted balance
        ///
        access(Withdraw) fun withdraw(amount: UFix64): @{Vault} {
            pre {
                self.balance >= amount:
                    "Amount withdrawn must be less than or equal than the balance of the Vault"
            }
            post {
                result.getType() == self.getType(): "Must return the same vault type as self"
                // use the special function `before` to get the value of the `balance` field
                // at the beginning of the function execution
                //
                self.balance == before(self.balance) - amount:
                    "New Vault balance must be the difference of the previous balance and the withdrawn Vault balance"
                emit Withdrawn(
                        type: result.getType().identifier,
                        amount: amount,
                        from: self.owner?.address,
                        fromUUID: self.uuid,
                        withdrawnUUID: result.uuid,
                        balanceAfter: self.balance
                )
            }
        }

        /// deposit takes a Vault and adds its balance to the balance of this Vault
        ///
        access(all) fun deposit(from: @{FungibleToken.Vault}) {
            // Assert that the concrete type of the deposited vault is the same
            // as the vault that is accepting the deposit
            pre {
                from.isInstance(self.getType()): 
                    "Cannot deposit an incompatible token type"
            }
            post {
                emit Deposited(
                        type: before(from.getType().identifier),
                        amount: before(from.balance),
                        to: self.owner?.address,
                        toUUID: self.uuid,
                        depositedUUID: before(from.uuid),
                        balanceAfter: self.balance
                )
                self.balance == before(self.balance) + before(from.balance):
                    "New Vault balance must be the sum of the previous balance and the deposited Vault"
            }
        }

        /// createEmptyVault allows any user to create a new Vault that has a zero balance
        ///
        access(all) fun createEmptyVault(): @{Vault} {
            post {
                result.balance == 0.0: "The newly created Vault must have zero balance"
                result.getType() == self.getType(): "The newly created Vault must have the same type as the creating vault"
            }
        }
    }

    /// createEmptyVault allows any user to create a new Vault that has a zero balance
    ///
    access(all) fun createEmptyVault(vaultType: Type): @{FungibleToken.Vault} {
        post {
            result.getType() == vaultType: "The returned vault does not match the desired type"
            result.balance == 0.0: "The newly created Vault must have zero balance"
        }
    }
}