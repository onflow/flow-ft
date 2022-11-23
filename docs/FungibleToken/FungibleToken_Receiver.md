# Resource Interface `Receiver`

```cadence
resource interface Receiver {
}
```

The interface that enforces the requirements for depositing
tokens into the implementing type.

We do not include a condition that checks the balance because
we want to give users the ability to make custom receivers that
can do custom things with the tokens, like split them up and
send them to different places.
## Functions

### fun `deposit()`

```cadence
func deposit(from Vault)
```
Takes a Vault and deposits it into the implementing resource type

Parameters:
  - from : _The Vault resource containing the funds that will be deposited_

---
