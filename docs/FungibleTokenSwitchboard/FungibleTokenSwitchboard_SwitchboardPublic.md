# Resource Interface `SwitchboardPublic`

```cadence
resource interface SwitchboardPublic {
}
```

The interface that enforces the method to allow anyone to check on the
available capabilities of a switchboard resource and also exposes the
deposit methods to deposit funds on it.
## Functions

### fun `getVaultTypes()`

```cadence
func getVaultTypes(): [Type]
```

---

### fun `deposit()`

```cadence
func deposit(from FungibleToken.Vault)
```

---

### fun `safeDeposit()`

```cadence
func safeDeposit(from FungibleToken.Vault): FungibleToken.Vault?
```

---
