# Resource `Administrator`

```cadence
resource Administrator {
}
```

## Functions

### fun `createNewMinter()`

```cadence
func createNewMinter(allowedAmount UFix64): Minter
```
Function that creates and returns a new minter resource

Parameters:
  - allowedAmount : _The maximum quantity of tokens that the minter could create_

Returns: The Minter resource that would allow to mint tokens

---

### fun `createNewBurner()`

```cadence
func createNewBurner(): Burner
```
Function that creates and returns a new burner resource

Returns: The Burner resource

---
