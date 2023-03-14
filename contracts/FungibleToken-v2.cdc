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

import FungibleTokenMetadataViews from "./FungibleTokenMetadataViews.cdc"

/// FungibleToken
///
/// Fungible Token implementations are no longer required to implement the fungible token
/// interface. We still have it as an interface here because there are some useful
/// utility methods that many projects will still want to have on their contracts,
/// but they are by no means required. all that is required is that the token
/// implements the `Vault` interface
pub contract interface FungibleToken {

    /// TokensWithdrawn
    ///
    /// The event that is emitted when tokens are withdrawn from a Vault
    pub event TokensWithdrawn(amount: UFix64, from: Address?, type: Type, ftView: FungibleTokenMetadataViews.FTView)

    /// TokensDeposited
    ///
    /// The event that is emitted when tokens are deposited to a Vault
    pub event TokensDeposited(amount: UFix64, to: Address?, type: Type, ftView: FungibleTokenMetadataViews.FTView)

    /// TokensTransferred
    ///
    /// The event that is emitted when tokens are transferred from one account to another
    pub event TokensTransferred(amount: UFix64, from: Address?, to: Address?, type: Type, ftView: FungibleTokenMetadataViews.FTView)

    /// TokensMinted
    ///
    /// The event that is emitted when new tokens are minted
    pub event TokensMinted(amount: UFix64, type: Type, ftView: FungibleTokenMetadataViews.FTView)

    /// Contains the total supply of the fungible tokens defined in this contract
    access(contract) var totalSupply: {Type: UFix64}

    /// Function to return the types that the contract implements
    pub fun getVaultTypes(): {Type: FungibleTokenMetadataViews.FTView} {
        post {
            result.length > 0: "Must indicate what fungible token types this contract defines"
        }
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
            }
        }
    }
    
    pub resource interface Transferable {
        /// Function for a direct transfer instead of having to do a deposit and withdrawal
        ///
        pub fun transfer(amount: UFix64, recipient: Capability<&{FungibleToken.Receiver}>)
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
        pub fun getSupportedVaultTypes(): {Type: Bool}
    }

    /// Balance
    ///
    /// This interface is now a general purpose metadata interface because
    /// a public interface is needed to get metadata, but adding a whole new interface
    /// for every account to upgrade to is probably too much of a breaking change
    pub resource interface Balance {

        /// Method to get the balance
        /// The balance could be a derived field,
        /// so there is no need to require an explicit field
        pub fun getBalance(): UFix64

        pub fun getSupportedVaultTypes(): {Type: Bool}

        /// MetadataViews Methods
        ///
        pub fun getViews(): [Type] {
            return []
        }

        pub fun resolveView(_ view: Type): AnyStruct? {
            return nil
        }
    }

    /// Vault
    ///
    /// Ideally, this interface would also conform to Receiver, Balance, Transferable, and Provider,
    /// but that is not supported yet
    ///
    pub resource interface Vault { //: Receiver, Balance, Transferable, Provider {

        /// Get the balance of the vault
        pub fun getBalance(): UFix64

        /// getSupportedVaultTypes optionally returns a list of vault types that this receiver accepts
        pub fun getSupportedVaultTypes(): {Type: Bool}

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
            }
            post {
                self.getBalance() == before(self.getBalance()) + before(from.getBalance()):
                    "New Vault balance must be the sum of the previous balance and the deposited Vault"
            }
        }

        /// Function for a direct transfer instead of having to do a deposit and withdrawal
        ///
        pub fun transfer(amount: UFix64, recipient: Capability<&{FungibleToken.Receiver}>) {
            post {
                self.getBalance() == before(self.getBalance()) - amount:
                    "New Vault balance from the sender must be the difference of the previous balance and the withdrawn Vault balance"
            }
        }

        /// createEmptyVault allows any user to create a new Vault that has a zero balance
        ///
        pub fun createEmptyVault(): @AnyResource{Vault} {
            post {
                result.getBalance() == 0.0: "The newly created Vault must have zero balance"
            }
        }
    }

    /// createEmptyVault allows any user to create a new Vault that has a zero balance
    ///
    pub fun createEmptyVault(vaultType: Type): @AnyResource{Vault}? {
        post {
            result.getBalance() == 0.0: "The newly created Vault must have zero balance"
        }
    }
}