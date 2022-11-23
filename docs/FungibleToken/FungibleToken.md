# Contract Interface `FungibleToken`

```cadence
contract interface FungibleToken {

    totalSupply:  UFix64
}
```

The interface that Fungible Token contracts implement.
## Interfaces
    
### resource interface `Provider`

```cadence
resource interface Provider {
}
```
The interface that enforces the requirements for withdrawing
tokens from the implementing type.

It does not enforce requirements on `balance` here,
because it leaves open the possibility of creating custom providers
that do not necessarily need their own balance.

[More...](FungibleToken_Provider.md)

---
    
### resource interface `Receiver`

```cadence
resource interface Receiver {
}
```
The interface that enforces the requirements for depositing
tokens into the implementing type.

We do not include a condition that checks the balance because
we want to give users the ability to make custom receivers that
can do custom things with the tokens, like split them up and
send them to different places.

[More...](FungibleToken_Receiver.md)

---
    
### resource interface `Balance`

```cadence
resource interface Balance {

    balance:  UFix64
}
```
The interface that contains the `balance` field of the Vault
and enforces that when new Vaults are created, the balance
is initialized correctly.

[More...](FungibleToken_Balance.md)

---
## Structs & Resources

### resource `Vault`

```cadence
resource Vault {

    balance:  UFix64
}
```
The resource that contains the functions to send and receive tokens.
The declaration of a concrete type in a contract interface means that
every Fungible Token contract that implements the FungibleToken interface
must define a concrete `Vault` resource that conforms to the `Provider`, `Receiver`,
and `Balance` interfaces, and declares their required fields and functions

[More...](FungibleToken_Vault.md)

---
## Functions

### fun `createEmptyVault()`

```cadence
func createEmptyVault(): Vault
```
Allows any user to create a new Vault that has a zero balance

Returns: The new Vault resource

---
## Events

### event `TokensInitialized`

```cadence
event TokensInitialized(initialSupply UFix64)
```
The event that is emitted when the contract is created

---

### event `TokensWithdrawn`

```cadence
event TokensWithdrawn(amount UFix64, from Address?)
```
The event that is emitted when tokens are withdrawn from a Vault

---

### event `TokensDeposited`

```cadence
event TokensDeposited(amount UFix64, to Address?)
```
The event that is emitted when tokens are deposited into a Vault

---
