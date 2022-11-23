# Contract `FungibleTokenSwitchboard`

```cadence
contract FungibleTokenSwitchboard {

    StoragePath:  StoragePath

    PublicPath:  PublicPath

    ReceiverPublicPath:  PublicPath
}
```

The contract that allows an account to receive payments in multiple fungible
tokens using a single `{FungibleToken.Receiver}` capability.
This capability should ideally be stored at the
`FungibleTokenSwitchboard.ReceiverPublicPath = /public/GenericFTReceiver`
but it can be stored anywhere.
## Interfaces
    
### resource interface `SwitchboardPublic`

```cadence
resource interface SwitchboardPublic {
}
```
The interface that enforces the method to allow anyone to check on the
available capabilities of a switchboard resource and also exposes the
deposit methods to deposit funds on it.

[More...](FungibleTokenSwitchboard_SwitchboardPublic.md)

---
## Structs & Resources

### resource `Switchboard`

```cadence
resource Switchboard {

    receiverCapabilities:  {Type: Capability<&{FungibleToken.Receiver}>}
}
```
The resource that stores the multiple fungible token receiver
capabilities, allowing the owner to add and remove them and anyone to
deposit any fungible token among the available types.

[More...](FungibleTokenSwitchboard_Switchboard.md)

---
## Functions

### fun `createSwitchboard()`

```cadence
func createSwitchboard(): Switchboard
```
Function that allows to create a new blank switchboard. A user must call
this function and store the returned resource in their storage.

---
## Events

### event `VaultCapabilityAdded`

```cadence
event VaultCapabilityAdded(type Type, switchboardOwner Address?, capabilityOwner Address?)
```
The event that is emitted when a new vault capability is added to a
switchboard resource.

---

### event `VaultCapabilityRemoved`

```cadence
event VaultCapabilityRemoved(type Type, switchboardOwner Address?, capabilityOwner Address?)
```
The event that is emitted when a vault capability is removed from a
switchboard resource.

---

### event `NotCompletedDeposit`

```cadence
event NotCompletedDeposit(type Type, amount UFix64, switchboardOwner Address?)
```
The event that is emitted when a deposit can not be completed.

---
