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

/// FungibleToken
///
pub contract FungibleToken {

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

        /// getAcceptedTypes optionally returns a list of vault types that this receiver accepts
        pub fun getAcceptedTypes(): {Type: Bool}
    }

    /// Balance
    ///
    pub resource interface Balance {

        /// Method to get the balance
        /// The balance could be a derived field,
        /// so there is no need to require an explicit field
        pub fun getBalance(): UFix64
    }

    /// Represents generic information about a vaults defined in the contract
    /// not information about a specific vault
    ///
    pub struct VaultInfo {

        /// The type of the vault represented
        pub let type: Type

        /// Storage and Public Paths
        pub let StoragePath: StoragePath
        pub let PublicReceiverBalancePath: PublicPath
        pub let PrivateProviderPath: PrivatePath

        init(type: Type, StoragePath: StoragePath, PublicReceiverBalancePath: PublicPath, PrivateProviderPath: PrivatePath) {
            self.type = type
            self.StoragePath = StoragePath
            self.PublicReceiverBalancePath = PublicReceiverBalancePath
            self.PrivateProviderPath = PrivateProviderPath
        }
    }

    /// Vault
    ///
    /// Ideally, this interface would also conform to Receiver, Balance, Transferable, and Provider,
    /// but that is not supported yet
    ///
    pub resource interface Vault { //: Receiver, Balance, Transferable, Provider {

        /// Storage and Public Paths
        pub let StoragePath: StoragePath
        pub let PublicReceiverBalancePath: PublicPath
        pub let PrivateProviderPath: PrivatePath

        /// Get the balance of the vault
        pub fun getBalance(): UFix64

        /// Return information about the vault's type and paths
        pub fun getVaultInfo(): FungibleToken.VaultInfo

        /// getAcceptedTypes optionally returns a list of vault types that this receiver accepts
        pub fun getAcceptedTypes(): {Type: Bool}

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
}