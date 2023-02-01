# Fungible Token Standard

This is a description of the Flow standard for fungible token contracts.  It is meant to contain the minimum requirements to implement a safe, secure, easy to understand, and easy to use fungible token contract. It also includes an example implementation to show how a concrete smart contract would actually implement the interface.

## What is Flow?

Flow is a new blockchain for open worlds. Read more about it [here](https://www.onflow.org/).

## What is Cadence?

Cadence is a new Resource-oriented programming language 
for developing smart contracts for the Flow Blockchain.
Read more about it [here](https://docs.onflow.org/docs) and see its implementation [here](https://github.com/onflow/cadence)

We recommend that anyone who is reading this should have already
completed the [Cadence Tutorials](https://docs.onflow.org/docs/getting-started-1) 
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

(note: default deployment of `FungibleTokenMetadataViews`
and `FungibleTokenSwitchboard` is still pending for emulator/canary, so you will still have to deploy those yourself on those networks)

| Network         | Contract Address     |
| --------------- | -------------------- |
| Emulator/Canary | `0xee82856bf20e2aa6` |
| Testnet         | `0x9a0766d93b6608b7` |
| Sandboxnet      | `0xe20612a0776ca4bf` |
| Mainnet         | `0xf233dcee88fe0abe` |


## Basics of the Standard:

The code for the standard is in `contracts/FungibleToken.cdc`. An example implementation of the standard that simulates what a simple token would be like is in `contracts/ExampleToken.cdc`. 

The exact smart contract that is used for the official Flow Network Token is in `contracts/FlowToken.cdc`

Example transactions that users could use to interact with fungible tokens are located in the `transactions/` directory. These templates are mostly generic and can be used with any fungible token implementation by providing the correct addresses, names, and values.

The standard consists of a contract interface called `FungibleToken` that requires implementing contracts to define a `Vault` resource that represents the tokens that an account owns. Each account that owns tokens will have a `Vault` stored in its account storage.  Users call functions on each other's `Vault`s to send and receive tokens.  

Right now we are using unsigned 64-bit fixed point numbers `UFix64` as the type to represent token balance information. This type has 8 decimal places and cannot represent negative numbers.

## Core Features (All contained in the main FungibleToken interface)

1- Getting metadata for the token smart contract via the fields of the contract:

- `pub var totalSupply: UFix64`
    - The only required field of the contract.  It would be incremented when new tokens are minted and decremented when they are destroyed.
- Event that gets emitted when the contract is initialized
    - `pub event TokensInitialized(initialSupply: UFix64)`

2- Retrieving the token fields of a `Vault` in an account that owns tokens.

- Balance interface
    - `pub var balance: UFix64`
        - The only required field of the `Vault` type

3- Withdrawing a specific amount of tokens *amount* using the *withdraw* function of the owner's `Vault`

- Provider interface
    - `pub fun withdraw(amount: UFix64): @FungibleToken.Vault`
        - Conditions
            - the returned Vault's balance must equal the amount withdrawn
            - The amount withdrawn must be less than or equal to the balance
            - The resulting balance must equal the initial balance - amount
    - Users can give other accounts a reference to their `Vault` cast as a `Provider` to allow them to withdraw and send tokens for them.  A contract can define any custom logic to govern the amount of tokens that can be withdrawn at a time with a `Provider`.  This can mimic the `approve`, `transferFrom` functionality of ERC20.
- withdraw event
    - Indicates how much was withdrawn and from what account the `Vault` is stored in.
      If the `Vault` is not in account storage when the event is emitted,
      `from` will be `nil`.
    - `pub event TokensWithdrawn(amount: UFix64, from: Address?)`

4 - Depositing a specific amount of tokens *from* using the *deposit* function of the recipient's `Vault`

- `Receiver` interface
    - `pub fun deposit(from: @FungibleToken.Vault)`
    - Conditions
        - `from` balance must be non-zero
        - The resulting balance must be equal to the initial balance + the balance of `from`
- deposit event
    - Indicates how much was deposited and to what account the `Vault` is stored in.
      If the `Vault` is not in account storage when the event is emitted,
      `to` will be `nil`.
    - `pub event TokensDeposited(amount: UFix64, to: Address?)`
- Users could create custom `Receiver`s to trigger special code when transfers to them happen, like forwarding the tokens
  to another account, splitting them up, and much more.

- It is important that if you are making your own implementation of the fungible token interface that
  you cast the input to `deposit` as the type of your token.
  `let vault <- from as! @ExampleToken.Vault`
  The interface specifies the argument as `@FungibleToken.Vault`, any resource that satisfies this can be sent to the deposit function. The interface checks that the concrete types match, but you'll still need to cast the `Vault` before storing it.

5 - Creating an empty Vault resource

- `pub fun createEmptyVault(): @FungibleToken.Vault`
- Defined in the contract 
  To create an empty `Vault`, the caller calls the function in the contract and stores the Vault in their storage.
- Conditions:
    - the balance of the returned Vault must be 0

6 - Destroying a Vault

If a `Vault` is explicitly destroyed using Cadence's `destroy` keyword, the balance of the destroyed vault must be subtracted from the total supply.

7 - Standard for Token Metadata

- not sure what this should be yet
- Could be a dictionary, could be an IPFS hash, could be json, etc.
- need suggestions!


## Comparison to Similar Standards in Ethereum

This spec covers much of the same ground that a spec like ERC-20 covers, but without most of the downsides.  

- Tokens cannot be sent to accounts or contracts that don't have owners or don't understand how to use them, because an account has to have a `Vault` in its storage to receive tokens.  No `safetransfer` is needed.
- If the recipient is a contract that has a stored `Vault`, the tokens can just be deposited to that Vault without having to do a clunky `approve`, `transferFrom`
- Events are defined in the contract for withdrawing and depositing, so a recipient will always be notified that someone has sent them tokens with the deposit event.
- The `approve`, `transferFrom` pattern is not included, so double spends are not permitted
- Transfers can trigger actions because users can define custom `Receivers` to execute certain code when a token is sent.
- Cadence integer types protect against overflow and underflow, so a `SafeMath`-equivalent library is not needed.

## FT Metadata

FT Metadata is represented in a flexible and modular way using both the [standard proposed in FLIP-0636](https://github.com/onflow/flow/blob/master/flips/20210916-nft-metadata.md) and the [standard proposed in FLIP-1087](https://github.com/onflow/flips/blob/main/flips/20220811-fungible-tokens-metadata.md).

When writing an NFT contract, you should implement the [`MetadataViews.Resolver`](contracts/utility/MetadataViews.cdc#L20-L23) interface, which allows your `Vault` resource to implement one or more metadata types called views.

Views do not specify or require how to store your metadata, they only specify
the format to query and return them, so projects can still be flexible with how they store their data.

### Fungible token Metadata Views

The [Example Token contract](contracts/ExampleToken.cdc) defines three new views that can used to communicate any fungible token information:

1. `FTView` A view that wraps the two other views that actually contain the data.
1. `FTDisplay` The view that contains all the information that will be needed by other dApps to display the fungible token: name, symbol, description, external URL, logos and links to social media.
1. `FTVaultData` The view that can be used by other dApps to interact programmatically with the fungible token, providing the information about the public and private paths used by default by the token, the public and private linked types for exposing capabilities and the function for creating new empty vaults. You can use this view to [setup an account using the vault stored in other account without the need of importing the actual token contract.](transactions/setup_account_from_vault_reference.cdc)

### How to implement metadata

The [Example Token contract](contracts/ExampleToken.cdc) shows how to implement metadata views for fungible tokens.

### How to read metadata

In this repository you can find examples on how to read metadata, accessing the `ExampleToken` display (name, symbol, logos, etc.) and its vault data (paths, linked types and the method to create a new vault).

First step will be to borrow a reference to the token's vault stored in some account:

```cadence
let vaultRef = account
    .getCapability(ExampleToken.VaultPublicPath)
    .borrow<&{MetadataViews.Resolver}>()
    ?? panic("Could not borrow a reference to the vault resolver")
```

Latter using that reference you can call methods defined in the [Fungible Token Metadata Views contract](contracts/FungibleTokenMetadataViews.cdc) that will return you the structure containing the desired information:

```cadence
let ftView = FungibleTokenMetadataViews.getFTView(viewResolver: vaultRef)
```

Alternatively you could call directly the `resolveView(_ view: Type): AnyStruct?` method on the `ExampleToken.Vault`, but the `getFTView(viewResolver: &{MetadataViews.Resolver}): FTView`, `getFTDisplay(_ viewResolver: &{MetadataViews.Resolver}): FTDisplay?` and `getFTVaultData(_ viewResolver: &{MetadataViews.Resolver}): FTVaultData?` defined on the `FungibleMetadataViews` contract will ease the process of dealing with optional types when retrieving this views.

Finally you can return the whole of structure or just log some values from the views depending on what you are aiming for:

```cadence
return ftView
````

```cadence
/*
When you retrieve a FTView both the FTDisplay and the FTVaultData views contained on it are optional values, meaning that the token could not be implementing then.
*/
log(ftView.ftDisplay!.symbol)
```

## Bonus Features

**Minting and Burning are not included in the standard but are included in the FlowToken example contract to illustrate what minting and burning might look like for a token in Flow.**

8 - Minting or Burning a specific amount of tokens using a specific minter resource that an owner can control

- `MintandBurn` Resource
    - function to mintTokens
    - tokens minted event
    - Each minter has a set amount of tokens that they are allowed to mint. This cannot be changed and a new minter needs to be created to add more allowance.
    - function to burnTokens
    - tokens Burnt event
    - Each time tokens are minted or burnt, that value is added or subtracted to or from the total supply.


**The following features could each be defined as a separate interface. It would be good to make standards for these, but not necessary to include in the main standard interface and are not currently defined in this example.**

9 - Withdrawing a specific amount of tokens from someone else's `Vault` by using their `provider` reference.

- approved withdraw event
- Providing a resource that only approves an account to send a specific amount per transaction or per day/month/etc.
- Returning the amount of tokens that an account can send for another account.
- Reading the balance of the account that you have permission to send tokens for
- Owner is able to increase and decrease the approval at will, or revoke it completely
    - This is much harder than anticipated

11 - Pausing Token transfers (maybe a way to prevent the contract from being imported)

12 - Cloning the token to create a new token with the same distribution

13 - Restricted ownership (For accredited investors and such)
- allowlisting
- denylisting

# How to use the Fungible Token contract

To use the Flow Token contract as is, you need to follow these steps:

1. If you are using the Playground, you need to deploy the `FungibleToken` definition to account 1 yourself and import it in `ExampleToken`. It is a pre-deployed interface in the emulator, testnet, and mainnet and you can import definition from those accounts:
    - `0xee82856bf20e2aa6` on emulator
    - `0x9a0766d93b6608b7` on testnet
    - `0xf233dcee88fe0abe` on mainnet
2. Deploy the `ExampleToken` definition
3. You can use the `get_balance.cdc` or `get_supply.cdc` scripts to read the 
   balance of a user's `Vault` or the total supply of all tokens, respectively.
4. Use the `setupAccount.cdc` on any account to set up the account to be able to
   use `FlowTokens`.
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
 
 This can be accomplished by executing the transaction found in this repository `transactions/switchboard/setup_account.cdc`. This transaction will create and save a Switchboard resource to the signer's account,
 and it also will create the needed public capabilities to access it. After setting up their switchboard, in order to make it support receiving a certain token, users will need to add the desired token's receiver capability to their switchboard resource.
 
 ## Adding a new vault to the switchboard
 When a user wants to receive a new fungible token through their switchboard, they will need to add a new public capability linked to said FT to their switchboard resource. This can be accomplished in two different ways:
 
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
    This function will panic if is not possible to `.borrow()` a reference to a `&{FungibleToken.Receiver}` from the passed capability. It will also panic if there is already a capability stored for the same `Type` of resource exposed by the capability.

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
    This function won't panic, instead it will just not add to the `@Switchboard` any capability which can not be retrieved from any of the provided `PublicPath`s. It will also ignore any type of `&{FungibleToken.Receiver}` that is already present on the `@Switchboard`

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
 If a user no longer wants to be able to receive deposits from a certain FT, or if they want to update the provided capability for one of them, they will need to remove the vault from the switchboard.
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
 The Fungible Token Switchboard provides two different ways of depositing tokens to it, using the `deposit(from: @FungibleToken.Vault)` method enforced by the `{FungibleToken.Receiver}` or using the `safeDeposit(from: @FungibleToken.Vault): @FungibleToken`:

 1. Using the first method will be just the same as depositing to `&{FungibleToken.Receiver}`. The path for the Switchboard receiver is defined in `FungibleTokenSwitchboard.ReceiverPublicPath`,
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

 2. The `safeDeposit(from: @FungibleToken.Vault): @FungibleToken` works in a similar way, with the difference that it will not panic if the desired FT Vault can not be obtained from the Switchboard. The method will return the passed vault, empty if the funds were deposited successfully or still containing the funds if the transfer of the funds was not possible. Keep in mind that when using this method on a transaction you will always have to deal with the returned resource. An example of this can be found on `transactions/switchboard/safe_transfer_tokens.cdc`:
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

You can find automated tests in the `lib/go/test/token_test.go` file. It uses the transaction templates that are contained in the `lib/go/templates/transaction_templates.go` file. Currently, these rely on a dependency from a private dapper labs repository to run, so external users will not be able to run them. We are working on making all of this public so anyone can run tests, but haven't completed this work yet.

## License 

The works in these folders are under the [Unlicense](https://github.com/onflow/flow-ft/blob/master/LICENSE):

- [/contracts](https://github.com/onflow/flow-ft/blob/master/contracts/)


