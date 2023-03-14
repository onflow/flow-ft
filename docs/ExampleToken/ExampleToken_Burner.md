# Resource `Burner`

```cadence
resource Burner {
}
```

Resource object that token admin accounts can hold to burn tokens.
## Functions

### fun `burnTokens()`

```cadence
func burnTokens(from FungibleToken.Vault)
```
Function that destroys a Vault instance, effectively burning the tokens.

Note: the burned tokens are automatically subtracted from the
total supply in the Vault destructor.

Parameters:
  - from : _The Vault resource containing the tokens to burn_

---
