# Resource Interface `Provider`

```cadence
resource interface Provider {
}
```

The interface that enforces the requirements for withdrawing
tokens from the implementing type.

It does not enforce requirements on `balance` here,
because it leaves open the possibility of creating custom providers
that do not necessarily need their own balance.
## Functions

### fun `withdraw()`

```cadence
func withdraw(amount UFix64): Vault
```
Subtracts tokens from the owner's Vault
and returns a Vault with the removed tokens.

The function's access level is public, but this is not a problem
because only the owner storing the resource in their account
can initially call this function.

The owner may grant other accounts access by creating a private
capability that allows specific other users to access
the provider resource through a reference.

The owner may also grant all accounts access by creating a public
capability that allows all users to access the provider
resource through a reference.

Parameters:
  - amount : _The amount of tokens to be withdrawn from the vault_

Returns: The Vault resource containing the withdrawn funds

---
