# Resource Interface `Balance`

```cadence
resource interface Balance {

    balance:  UFix64
}
```

The interface that contains the `balance` field of the Vault
and enforces that when new Vaults are created, the balance
is initialized correctly.
## Functions

### fun `getViews()`

```cadence
func getViews(): [Type]
```
Function that returns all the Metadata Views implemented by a Fungible Token

developers to know which parameter to pass to the resolveView() method.

Returns: An array of Types defining the implemented views. This value will be used by

---

### fun `resolveView()`

```cadence
func resolveView(_ Type): AnyStruct?
```
Function that resolves a metadata view for this fungible token by type.

Parameters:
  - view : _The Type of the desired view._

Returns: A structure representing the requested view.

---
