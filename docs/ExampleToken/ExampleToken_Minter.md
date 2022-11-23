# Resource `Minter`

```cadence
resource Minter {

    allowedAmount:  UFix64
}
```

Resource object that token admin accounts can hold to mint new tokens.

### Initializer

```cadence
func init(allowedAmount UFix64)
```


## Functions

### fun `mintTokens()`

```cadence
func mintTokens(amount UFix64): ExampleToken.Vault
```
Function that mints new tokens, adds them to the total supply,
and returns them to the calling context.

Parameters:
  - amount : _The quantity of tokens to mint_

Returns: The Vault resource containing the minted tokens

---
