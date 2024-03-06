# Fungible Token Standard

This is a description of the Flow standard for fungible token contracts. 
It is meant to contain the minimum requirements to implement a safe, secure, easy to understand,
and easy to use fungible token contract.
It also includes an example implementation to show how a 
concrete smart contract would actually implement the interface.

The version of the contracts in the `master` branch is the
Cadence 1.0 version of the contracts and is not the same
as the ones that are currently deployed to testnet and mainnet.
See the `cadence-0.42` branch for the currently deployed versions.

## What is Flow?

Flow is a new blockchain for open worlds. Read more about it [here](https://www.flow.com/).

## What is Cadence?

Cadence is a new Resource-oriented programming language 
for developing smart contracts for the Flow Blockchain.
Read more about it [here](https://developers.flow.com/) and see its implementation [here](https://github.com/onflow/cadence)

We recommend that anyone who is reading this should have already
completed the [Cadence Tutorials](https://cadence-lang.org/docs/tutorial/first-steps) 
so they can build a basic understanding of the programming language.

Resource-oriented programming, and by extension Cadence, 
is the perfect programming environment for currencies, because users are able
to store their tokens directly in their accounts and transact
peer-to-peer. Please see the [blog post about resources](https://medium.com/dapperlabs/resource-oriented-programming-bee4d69c8f8e)
to understand why they are perfect for digital assets.

## Import Addresses

The `FungibleToken`, `FungibleTokenMetadataViews`, and `FungibleTokenSwitchboard` contracts are already deployed
on various networks. You can import them in your contracts from these addresses.
There is no need to deploy them yourself.

| Network                      | Contract Address     |
| ---------------------------- | -------------------- |
| Emulator                     | `0xee82856bf20e2aa6` |
| Testnet/Previewnet/Crescendo | `0x9a0766d93b6608b7` |
| Sandboxnet                   | `0xe20612a0776ca4bf` |
| Mainnet                      | `0xf233dcee88fe0abe` |

The `Burner` contract is also deployed to these addresses, but should not be used until after the Cadence 1.0 network upgrade.

## Basics of the Standard:

The code for the standard is in [`contracts/FungibleToken.cdc`](contracts/FungibleToken.cdc). An example implementation of the standard that simulates what a simple token would be like is in [`contracts/ExampleToken.cdc`](contracts/FungibleToken.cdc). 

The exact smart contract that is used for the official Flow Network Token (`FlowToken`) is in [the `flow-core-contracts` repository](https://github.com/onflow/flow-core-contracts/blob/master/contracts/FlowToken.cdc).

Example transactions that users could use to interact with fungible tokens are located in the `transactions/` directory. These templates are mostly generic and can be used with any fungible token implementation by providing the correct addresses, names, and values.

The standard consists of a contract interface called `FungibleToken` that defines important
functionality for token implementations. Contracts are expected to define a resource
that implement the `FungibleToken.Vault` resource interface.
A `Vault` represents the tokens that an account owns. Each account that owns tokens
will have a `Vault` stored in its account storage. 
Users call functions on each other's `Vault`s to send and receive tokens.  

The standard uses unsigned 64-bit fixed point numbers `UFix64` as the type to represent token balance information. This type has 8 decimal places and cannot represent negative numbers.

## Core Features (All contained in the main FungibleToken interface)

### `Balance` Interface

Specifies that the implementing type must have a `UFix64` `balance` field.
  - `access(all) var balance: UFix64`

### `Provider` Interface
Defines a [`withdraw ` function](contracts/FungibleToken.cdc#L95) for withdrawing a specific amount of tokens *amount*.
  - `access(all) fun withdraw(amount: UFix64): @{FungibleToken.Vault}`
      - Conditions
          - the returned Vault's balance must equal the amount withdrawn
          - The amount withdrawn must be less than or equal to the balance
          - The resulting balance must equal the initial balance - amount
  - Users can give other accounts a reference to their `Vault` cast as a `Provider`
  to allow them to withdraw and send tokens for them. 
  A contract can define any custom logic to govern the amount of tokens
  that can be withdrawn at a time with a `Provider`. 
  This can mimic the `approve`, `transferFrom` functionality of ERC20.
- [`FungibleToken.Withdrawn` event](contracts/FungibleToken.cdc#L50)
    - Event that is emitted automatically to indicate how much was withdrawn
    and from what account the `Vault` is stored in.
      If the `Vault` is not in account storage when the event is emitted,
      `from` will be `nil`.
    - Contracts do not have to emit their own events,
    the standard events will automatically be emitted.

Defines [an `isAvailableToWithdraw()` function](contracts/FungibleToken.cdc#L95)
to ask a `Provider` if the specified number of tokens can be withdrawn from the implementing type.

### `Receiver` Interface
Defines functionality to depositing fungible tokens into a resource object.
- [`deposit()` function](contracts/FungibleToken.cdc#L119):
  - `access(all) fun deposit(from: @{FungibleToken.Vault})`
  - Conditions
      - `from` balance must be non-zero
      - The resulting balance must be equal to the initial balance + the balance of `from`
  - It is important that if you are making your own implementation of the fungible token interface that
  you cast the input to `deposit` as the type of your token.
  `let vault <- from as! @ExampleToken.Vault`
  The interface specifies the argument as `@FungibleToken.Vault`, any resource that satisfies this can be sent to the deposit function. The interface checks that the concrete types match, but you'll still need to cast the `Vault` before storing it.
- deposit event
    - [`FungibleToken.Deposited` event](contracts/FungibleToken.cdc#L53) from the standard
    that indicates how much was deposited and to what account the `Vault` is stored in.
      - If the `Vault` is not in account storage when the event is emitted,
        `to` will be `nil`.
      - This event is emitted automatically on any deposit, so projects do not need
        to define and emit their own events.

Defines Functionality for Getting Supported Vault Types
- Some resource types can accept multiple different vault types in their deposit functions,
  so the `getSupportedVaultTypes()` and `isSupportedVaultType()` functions allow callers
  to query a resource that implements `Receiver` to see if the `Receiver` accepts
  their desired `Vault` type in its deposit function.

Users could create custom `Receiver`s to trigger special code when transfers to them happen,
like forwarding the tokens to another account, splitting them up, and much more.

### `Vault` Interface
[Interface](contracts/FungibleToken.cdc#L134) that inherits from `Provider`, `Receiver`, `Balance`, `ViewResolver.Resolver`,
and `Burner.Burnable` and provides additional pre and post conditions.

The `ViewResolver.Resolver` interface defines functionality for retrieving metadata
about a particular resource object. [Fungible Token metadata](README.md#ft-metadata) is described below.

See the comments in [the `Burner` contract](contracts/Burner.cdc) for context about it.
Basically, it defines functionality for tokens to have custom logic when those tokens
are destroyed.

### Creating an empty Vault resource
Defines functionality in the contract to create a new empty vault of
of the contract's defined type.
- `access(all) fun createEmptyVault(vaultType: Type): @{FungibleToken.Vault}`
- Defined in the contract 
- To create an empty `Vault`, the caller calls the function and provides the Vault Type
  that they want. They get a vault back and can store it in their storage.
- Conditions:
    - the balance of the returned Vault must be 0


## Comparison to Similar Standards in Ethereum

This spec covers much of the same ground that a spec like ERC-20 covers, but without most of the downsides.  

- Tokens cannot be sent to accounts or contracts that don't have owners or don't understand how to use them, because an account has to have a `Vault` in its storage to receive tokens.  No `safetransfer` is needed.
- If the recipient is a contract that has a stored `Vault`, the tokens can just be deposited to that Vault without having to do a clunky `approve`, `transferFrom`
- Events are defined in the contract for withdrawing and depositing, so a recipient will always be notified that someone has sent them tokens with the deposit event.
- The `approve`, `transferFrom` pattern is not included, so double spends are not permitted
- Transfers can trigger actions because users can define custom `Receivers` to execute certain code when a token is sent.
- Cadence integer types protect against overflow and underflow, so a `SafeMath`-equivalent library is not needed.

## FT Metadata

FT Metadata is represented in a flexible and modular way using both
the [standard proposed in FLIP-0636](https://github.com/onflow/flips/blob/main/application/20210916-nft-metadata.md)
and the [standard proposed in FLIP-1087](https://github.com/onflow/flips/blob/main/application/20220811-fungible-tokens-metadata.md).

[A guide for NFT metadata](https://developers.flow.com/build/advanced-concepts/metadata-views)
is provided on the docs site. Many of the concepts described there also apply
to fungible tokens, so it is useful to read for any Cadence developer.

When writing an FT contract interface, your contract will implement
the `FungibleToken` contract interface which already inherits
from [the `ViewResolver` contract interface](https://github.com/onflow/flow-nft/blob/master/contracts/ViewResolver.cdc),
so you will be required to implement the metadata functions.
Additionally, your `Vault` will also implement the `ViewResolver.Resolver` by default,
which allows your `Vault` resource to implement one or more metadata types called views.

Views do not specify or require how to store your metadata, they only specify
the format to query and return them, so projects can still be flexible with how they store their data.

### Fungible token Metadata Views

The [FungibleTokenMetadataViews contract](contracts/FungibleTokenMetadataViews.cdc) defines four new views that can used to communicate any fungible token information:

1. `FTView`: A view that wraps the two other views that actually contain the data.
2. `FTDisplay`: The view that contains all the information that will be needed by other dApps to display the fungible token: name, symbol, description, external URL, logos and links to social media.
3. `FTVaultData`: The view that can be used by other dApps to interact programmatically with the fungible token, providing the information about the public and private paths used by default by the token, the public and private linked types for exposing capabilities and the function for creating new empty vaults. You can use this view to [setup an account using the vault stored in other account without the need of importing the actual token contract.](transactions/setup_account_from_vault_reference.cdc)
4. `TotalSupply`: Specifies the total supply of the given token.

### How to implement metadata

The [Example Token contract](contracts/ExampleToken.cdc) shows how to implement metadata views for fungible tokens.

### How to read metadata

In this repository you can find examples on how to read metadata, accessing the `ExampleToken` display (name, symbol, logos, etc.) and its vault data (paths, linked types and the method to create a new vault).

Latter using that reference you can call methods defined in the [Fungible Token Metadata Views contract](contracts/FungibleTokenMetadataViews.cdc) that will return you the structure containing the desired information:



## Bonus Features

The following features could each be defined as a separate standard. It would be good to make standards for these, but not necessary to include in the main standard interface and are not currently defined in this example.

- Scoped Provider 
This refers to restricting a `Provider` capability to only be able to withdraw a specific amount of tokens from someone else's `Vault`
This is currently being worked on.

- Pausing Token transfers (maybe a way to prevent the contract from being imported)
- Cloning the token to create a new token with the same distribution
- Restricted ownership (For accredited investors and such)
- allowlisting
- denylisting

# How to use the Fungible Token contract

To use the Flow Token contract as is, you need to follow these steps:

1. If you are using any network or the playground, there is no need to deploy
the `FungibleToken` definition to accounts yourself.
It is a pre-deployed interface in the emulator, testnet, mainnet,
and playground and you can import definition from those accounts:
    - `0xee82856bf20e2aa6` on emulator
    - `0x9a0766d93b6608b7` on testnet
    - `0xf233dcee88fe0abe` on mainnet
2. Deploy the `ExampleToken` definition, making sure to import the `FungibleToken` interface.
3. You can use the `get_balance.cdc` or `get_supply.cdc` scripts to read the 
   balance of a user's `Vault` or the total supply of all tokens, respectively.
4. Use the `setup_account.cdc` on any account to set up the account to be able to
   use `ExampleToken`.
5. Use the `transfer_tokens.cdc` transaction file to send tokens from one user with
   a `Vault` in their account storage to another user with a `Vault` in their account storage.
6. Use the `mint_tokens.cdc` transaction with the admin account to mint new tokens.
7. Use the `burn_tokens.cdc` transaction with the admin account to burn tokens.
8. Use the `create_minter.cdc` transaction to create a new MintandBurn resource
   and store it in a new Admin's account.

# Fungible Token Switchboard

`FungibleTokenSwitchboard.cdc`, allows users to receive payments in different fungible tokens using a single `&{FungibleToken.Receiver}` placed in a standard receiver path `/public/GenericFTReceiver`.

 ## How to use it

 Users willing to use the Fungible Token Switchboard will need to setup their accounts by creating a new `FungibleTokenSwitchboard.Switchboard` resource and saving it to their accounts at the `FungibleTokenSwitchboard.StoragePath` path.
 
 This can be accomplished by executing the transaction found in this repository `transactions/switchboard/setup_account.cdc`.
 This transaction will create and save a Switchboard resource to the signer's account,
 and it also will create the needed public capabilities to access it.
 After setting up their switchboard, in order to make it support receiving a certain token,
 users will need to add the desired token's receiver capability to their switchboard resource.
 
 ## Adding a new vault to the switchboard
 When a user wants to receive a new fungible token through their switchboard,
 they will need to add a new public capability linked to said FT
 to their switchboard resource. This can be accomplished in two different ways:
 
 1. Adding a single capability using `addNewVault(capability: Capability<&{FungibleToken.Receiver}>)`
    * Before calling this method on a transaction you should first retrieve the capability to the token's vault you are
    willing to add to the switchboard, as is done in the template transaction `transactions/switchboard/add_vault_capability.cdc`.

    ```cadence
    transaction {
        let exampleTokenVaultCapabilty: Capability<&{FungibleToken.Receiver}>
        let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

        prepare(signer: AuthAccount) {
          // Get the example token vault capability from the signer's account
          self.exampleTokenVaultCapability = 
            signer.getCapability<&{FungibleToken.Receiver}>
                                    (ExampleToken.ReceiverPublicPath)
          // Get a reference to the signers switchboard
          self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
            (from: FungibleTokenSwitchboard.StoragePath) 
              ?? panic("Could not borrow reference to switchboard")
        }

        execute {
          // Add the capability to the switchboard using addNewVault method
          self.switchboardRef.addNewVault(capability: self.exampleTokenVaultCapability)
        }
    }
    ```
    This function will panic if is not possible to `.borrow()` a reference
    to a `&{FungibleToken.Receiver}` from the passed capability.
    It will also panic if there is already a capability stored
    for the same `Type` of resource exposed by the capability.

 2. Adding one or more capabilities using the paths where they are stored using `addNewVaultsByPath(paths: [PublicPath], address: Address)`
    * When using this method, an array of `PublicPath` objects should be passed along with the `Address` of the account from where the vaults' capabilities should be retrieved.

    ```cadence
    transaction (address: Address) {

        let exampleTokenVaultPath: PublicPath
        let vaultPaths: [PublicPath]
        let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

        prepare(signer: AuthAccount) {
          // Get the example token vault path from the contract
          self.exampleTokenVaultPath = ExampleToken.ReceiverPublicPath
          // And store it in the array of public paths that will be passed to the
          // switchboard method
          self.vaultPaths = []
          self.vaultPaths.append(self.exampleTokenVaultPath)
          // Get a reference to the signers switchboard
          self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
            (from: FungibleTokenSwitchboard.StoragePath) 
              ?? panic("Could not borrow reference to switchboard")
        }

        execute {
          // Add the capability to the switchboard using addNewVault method
          self.switchboardRef.addNewVaultsByPath(paths: self.vaultPaths, 
                                                        address: address)
        }
    }
    ```
    This function won't panic, instead it will just not add to the `@Switchboard`
    any capability which can not be retrieved from any of the provided `PublicPath`s.
    It will also ignore any type of `&{FungibleToken.Receiver}` that is already present on the `@Switchboard`

  3. Adding a capability to a receiver specifying which type of token will be deposited there 
  using `addNewVaultWrapper(capability: Capability<&{FungibleToken.Receiver}>, type: Type)`. 
  This method can be used to link a token forwarder or any other wrapper to the switchboard. 
  Once the `Forwarder` has been properly created containing the capability to an actual `@FungibleToken.Vault`,
  this method can be used to link the `@Forwarder` to the switchboard to deposit the specified type of Fungible Token.
  In the template transaction  `switchboard/add_vault_wrapper_capability.cdc`,
  we assume that the signer has a forwarder containing a capability to an `@ExampleToken.Vault` resource:

  ```cadence
  transaction {

    let tokenForwarderCapability: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

        // Get the token forwarder capability from the signer's account
        self.tokenForwarderCapability = 
            signer.getCapability<&{FungibleToken.Receiver}>
                                (ExampleToken.ReceiverPublicPath)
        
        // Check if the receiver capability exists
        assert(self.tokenForwarderCapability.check(), 
            message: "Signer does not have a working fungible token receiver capability")
        
        // Get a reference to the signers switchboard
        self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
            (from: FungibleTokenSwitchboard.StoragePath) 
            ?? panic("Could not borrow reference to switchboard")
    
    }

    execute {

        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVaultWrapper(capability: self.tokenForwarderCapability, type: Type<@ExampleToken.Vault>())
    
    }

  }
  ```

 ## Removing a vault from the switchboard
 If a user no longer wants to be able to receive deposits from a certain FT,
 or if they want to update the provided capability for one of them,
 they will need to remove the vault from the switchboard.
 This can be accomplished by using `removeVault(capability: Capability<&{FungibleToken.Receiver}>)`. 
This can be observed in the template transaction `transactions/switchboard/remove_vault_capability.cdc`:
 ```cadence
 transaction {
    let exampleTokenVaultCapabilty: Capability<&{FungibleToken.Receiver}>
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {
      // Get the example token vault capability from the signer's account
      self.exampleTokenVaultCapability = signer.getCapability
                    <&{FungibleToken.Receiver}>(ExampleToken.ReceiverPublicPath)
      // Get a reference to the signers switchboard  
      self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
        (from: FungibleTokenSwitchboard.StoragePath) 
          ?? panic("Could not borrow reference to switchboard")

    }

    execute {
      // Remove the capability from the switchboard using the 
      // removeVault method
      self.switchboardRef.removeVault(capability: self.exampleTokenVaultCapability)
    }
 }
 ```
 This function will panic if is not possible to `.borrow()` a reference to a `&{FungibleToken.Receiver}` from the passed capability.

 ## Transferring tokens through the switchboard
 The Fungible Token Switchboard provides two different ways of depositing tokens to it,
 using the `deposit(from: @FungibleToken.Vault)` method enforced by
 the `{FungibleToken.Receiver}` or using the `safeDeposit(from: @FungibleToken.Vault): @FungibleToken`:

 1. Using the first method will be just the same as depositing to `&{FungibleToken.Receiver}`.
 The path for the Switchboard receiver is defined in `FungibleTokenSwitchboard.ReceiverPublicPath`,
 the generic receiver path `/public/GenericFTReceiver` that can also be obtained from the NFT MetadataViews contract.
 An example of how to do this can be found in the transaction template on this repo `transactions/switchboard/transfer_tokens.cdc`
 ```cadence
 transaction(to: Address, amount: UFix64) {
    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        let vaultRef = signer.borrow<&ExampleToken.Vault>
                                    (from: ExampleToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        self.sentVault <- vaultRef.withdraw(amount: amount)
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Switchboard Receiver
        let switchboardRef = recipient.getCapability
            (FungibleTokenSwitchboard.ReceiverPublicPath)
            .borrow<&{FungibleToken.Receiver}>()
			    ?? panic("Could not borrow receiver reference to switchboard!")

        // Deposit the withdrawn tokens in the recipient's switchboard receiver
        switchboardRef.deposit(from: <-self.sentVault)
    }
 }
 ```

 2. The `safeDeposit(from: @FungibleToken.Vault): @FungibleToken` works in a similar way,
 with the difference that it will not panic if the desired FT Vault
 can not be obtained from the Switchboard. The method will return the passed vault,
 empty if the funds were deposited successfully or still containing the funds
 if the transfer of the funds was not possible. Keep in mind that 
 when using this method on a transaction you will always have to deal
 with the returned resource. An example of this can be found on `transactions/switchboard/safe_transfer_tokens.cdc`:
 ```cadence
 transaction(to: Address, amount: UFix64) {
    // The reference to the vault from the payer's account
    let vaultRef: &ExampleToken.Vault
    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault


    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.borrow<&ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        self.sentVault <- self.vaultRef.withdraw(amount: amount)
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        // Get a reference to the recipient's Switchboard Receiver
        let switchboardRef = recipient.getCapability(FungibleTokenSwitchboard.PublicPath)
            .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
			?? panic("Could not borrow receiver reference to switchboard!")

        // Deposit the withdrawn tokens in the recipient's switchboard receiver,
        // then deposit the returned vault in the signer's vault
        self.vaultRef.deposit(from: <- switchboardRef.safeDeposit(from: <-self.sentVault))
    }
 }
 ```

# Running Automated Tests

There are two sets of tests in the repo, Cadence tests and Go tests.
The Cadence tests are much more straightforward nad are all written in Cadence,
so we recommend following those.

## Cadence Testing Framework

The Cadence tests are located in the `tests/` repository. They are written in Cadence
and can be run directly from the command line using the Flow CLI.
Make sure you are using [the latest Cadence 1.0 CLI verion](https://forum.flow.com/t/update-on-cadence-1-0/5197/10).
```
flow test --cover --covercode="contracts" tests/*.cdc
```

## Go tests

You can find automated tests in the `lib/go/test/token_test.go` file.
It uses the transaction templates that are contained in the
`lib/go/templates/transaction_templates.go` file. You can run them by navigating
to the `lib/go/test/` directory and running `go test -v`.
If you make changes to the contracts or transactions in between running tests,
you will need to run `make generate` from the `lib/go/` directory
to generate the assets used in the tests.

## License 

The works in these folders are under the [Unlicense](https://github.com/onflow/flow-ft/blob/master/LICENSE):

- [/contracts](https://github.com/onflow/flow-ft/blob/master/contracts/)


