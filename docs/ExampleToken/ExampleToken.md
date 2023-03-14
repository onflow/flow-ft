# Contract `ExampleToken`

```cadence
contract ExampleToken {

    totalSupply:  UFix64

    VaultStoragePath:  StoragePath

    ReceiverPublicPath:  PublicPath

    VaultPublicPath:  PublicPath

    AdminStoragePath:  StoragePath
}
```


Implemented Interfaces:
  - `FungibleToken`

## Structs & Resources

### resource `Vault`

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

[More...](ExampleToken_Vault.md)

---

### resource `Administrator`

```cadence
resource Administrator {
}
```

[More...](ExampleToken_Administrator.md)

---

### resource `Minter`

```cadence
resource Minter {

    allowedAmount:  UFix64
}
```
Resource object that token admin accounts can hold to mint new tokens.

[More...](ExampleToken_Minter.md)

---

### resource `Burner`

```cadence
resource Burner {
}
```
Resource object that token admin accounts can hold to burn tokens.

[More...](ExampleToken_Burner.md)

---
## Functions

### fun `createEmptyVault()`

```cadence
func createEmptyVault(): Vault
```
Function that creates a new Vault with a balance of zero
and returns it to the calling context. A user must call this function
and store the returned Vault in their storage in order to allow their
account to be able to receive deposits of this token type.

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
The event that is emitted when tokens are deposited to a Vault

---

### event `TokensMinted`

```cadence
event TokensMinted(amount UFix64)
```
The event that is emitted when new tokens are minted

---

### event `TokensBurned`

```cadence
event TokensBurned(amount UFix64)
```
The event that is emitted when tokens are destroyed

---

### event `MinterCreated`

```cadence
event MinterCreated(allowedAmount UFix64)
```
The event that is emitted when a new minter resource is created

---

### event `BurnerCreated`

```cadence
event BurnerCreated()
```
The event that is emitted when a new burner resource is created

---
