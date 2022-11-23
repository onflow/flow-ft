# Resource `Vault`

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

Implemented Interfaces:
  - `Provider`
  - `Receiver`
  - `Balance`


### Initializer

```cadence
func init(balance UFix64)
```


## Functions

### fun `withdraw()`

```cadence
func withdraw(amount UFix64): Vault
```

---

### fun `deposit()`

```cadence
func deposit(from Vault)
```
Takes a Vault and deposits it into the implementing resource type

Parameters:
  - from : _The Vault resource containing the funds that will be deposited_

---
