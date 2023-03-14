# Resource `Vault`

```cadence
resource Vault {

    balance:  UFix64
}
```

Each user stores an instance of only the Vault in their storage
The functions in the Vault and governed by the pre and post conditions
in FungibleToken when they are called.
The checks happen at runtime whenever a function is called.

Resources can only be created in the context of the contract that they
are defined in, so there is no way for a malicious user to create Vaults
out of thin air. A special Minter resource needs to be defined to mint
new tokens.

Implemented Interfaces:
  - `FungibleToken.Provider`
  - `FungibleToken.Receiver`
  - `FungibleToken.Balance`
  - `MetadataViews.Resolver`


### Initializer

```cadence
func init(balance UFix64)
```


## Functions

### fun `withdraw()`

```cadence
func withdraw(amount UFix64): FungibleToken.Vault
```
Function that takes an amount as an argument
and withdraws that amount from the Vault.
It creates a new temporary Vault that is used to hold
the money that is being transferred. It returns the newly
created Vault to the context that called so it can be deposited
elsewhere.

Parameters:
  - amount : _The amount of tokens to be withdrawn from the vault_

Returns: The Vault resource containing the withdrawn funds

---

### fun `deposit()`

```cadence
func deposit(from FungibleToken.Vault)
```
Function that takes a Vault object as an argument and adds
its balance to the balance of the owners Vault.
It is allowed to destroy the sent Vault because the Vault
was a temporary holder of the tokens. The Vault's balance has
been consumed and therefore can be destroyed.

Parameters:
  - from : _The Vault resource containing the funds that will be deposited_

---

### fun `getViews()`

```cadence
func getViews(): [Type]
```
The way of getting all the Metadata Views implemented by ExampleToken

developers to know which parameter to pass to the resolveView() method.

Returns: An array of Types defining the implemented views. This value will be used by

---

### fun `resolveView()`

```cadence
func resolveView(_ Type): AnyStruct?
```
The way of getting a Metadata View out of the ExampleToken

Parameters:
  - view : _The Type of the desired view._

Returns: A structure representing the requested view.

---
