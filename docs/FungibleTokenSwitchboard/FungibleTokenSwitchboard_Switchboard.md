# Resource `Switchboard`

```cadence
resource Switchboard {

    receiverCapabilities:  {Type: Capability<&{FungibleToken.Receiver}>}
}
```

The resource that stores the multiple fungible token receiver
capabilities, allowing the owner to add and remove them and anyone to
deposit any fungible token among the available types.

Implemented Interfaces:
  - `FungibleToken.Receiver`
  - `SwitchboardPublic`


### Initializer

```cadence
func init()
```


## Functions

### fun `addNewVault()`

```cadence
func addNewVault(capability Capability<&{FungibleToken.Receiver}>)
```
Adds a new fungible token receiver capability to the switchboard
resource.

token vault deposit function through `{FungibleToken.Receiver}` that
will be added to the switchboard.

Parameters:
  - capability : _The capability to expose a certain fungible_

---

### fun `addNewVaultsByPath()`

```cadence
func addNewVaultsByPath(paths [PublicPath], address Address)
```
Adds a number of new fungible token receiver capabilities by using
the paths where they are stored.

Parameters:
  - paths : _The paths where the public capabilities are stored._
  - address : _The address of the owner of the capabilities._

---

### fun `addNewVaultWrapper()`

```cadence
func addNewVaultWrapper(capability Capability<&{FungibleToken.Receiver}>, type Type)
```
Adds a new fungible token receiver capability to the switchboard
resource specifying which `Type`of `@FungibleToken.Vault` can be
deposited to it. Use it to include in your switchboard "wrapper"
receivers such as a `@TokenForwarding.Forwarder`. It can also be
used to overwrite the type attached to a certain capability without
having to remove that capability first.

token vault deposit function through `{FungibleToken.Receiver}` that
will be added to the switchboard.

capability, rather than the `Type` from the reference borrowed from
said capability

Parameters:
  - capability : _The capability to expose a certain fungible_
  - type : _The type of fungible token that can be deposited to that_

---

### fun `removeVault()`

```cadence
func removeVault(capability Capability<&{FungibleToken.Receiver}>)
```
Removes a fungible token receiver capability from the switchboard
resource.

removed from the switchboard.

Parameters:
  - capability : _The capability to a fungible token vault to be_

---

### fun `deposit()`

```cadence
func deposit(from FungibleToken.Vault)
```
Takes a fungible token vault and routes it to the proper fungible
token receiver capability for depositing it.

Parameters:
  - from : _The deposited fungible token vault resource._

---

### fun `safeDeposit()`

```cadence
func safeDeposit(from FungibleToken.Vault): FungibleToken.Vault?
```
Takes a fungible token vault and tries to route it to the proper
fungible token receiver capability for depositing the funds,
avoiding panicking if the vault is not available.

deposited.

funds if the deposit was successful, or still containing the funds
if the reference to the needed vault was not found.

Parameters:
  - vaultType : _The type of the ft vault that wants to be_

Returns: The deposited fungible token vault resource, without the

---

### fun `getVaultTypes()`

```cadence
func getVaultTypes(): [Type]
```
A getter function to know which tokens a certain switchboard
resource is prepared to receive.

`{FungibleToken.Receiver}` capabilities that can be effectively
borrowed.

Returns: The keys from the dictionary of stored

---
