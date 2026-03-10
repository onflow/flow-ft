# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests (Cadence + Go)
make test

# Run only Cadence tests (faster, preferred for contract changes)
flow test --cover --covercode="contracts" tests/*.cdc

# Run only Go tests
cd lib/go && make test

# Run a single Go test by name
cd lib/go/test && go test -v -run TestTokenForwarding

# Regenerate Go assets after changing any .cdc file (required before Go tests will reflect contract changes)
cd lib/go && make generate
```

The `make generate` step embeds `.cdc` files as Go byte arrays via `go-bindata`. If you edit a contract or transaction and then run Go tests without regenerating, the tests will use stale code.

## Architecture

### Two test suites, one source of truth

The `.cdc` files in `contracts/` and `transactions/` are the canonical source. They are used by both:

1. **Cadence tests** (`tests/*.cdc`) — run directly by `flow test` against the Flow testing framework. These are the primary tests and easiest to write/read. All new tests should preferably be written in Cadence.
2. **Go tests** (`lib/go/test/`) — use `go-bindata`-embedded copies of the same `.cdc` files (via `lib/go/contracts/internal/assets/`). The Go layer also has template helpers in `lib/go/templates/` that inject contract addresses into transaction/script strings at test time.

### Contract hierarchy

`FungibleToken.cdc` is a **contract interface** (not a concrete contract). It defines the `Vault` resource interface, `Withdraw` entitlement, and standard events (`Withdrawn`, `Deposited`, `Burned`). Crucially, it enforces pre/post conditions on `withdraw` and `deposit` at the interface level — implementations get these for free.

`ExampleToken.cdc` is the reference implementation of `FungibleToken`. It is the token used in all tests and transaction templates.

`Burner.cdc` is a standalone utility that provides the `burn()` function and `Burnable` interface. Vaults implement `burnCallback()` (called by `Burner.burn()`) to update total supply when tokens are destroyed. Direct `destroy` on a vault is not the correct pattern — always use `Burner.burn()`.

### Utility contracts

- `TokenForwarding.cdc` — A `Forwarder` resource that implements `FungibleToken.Receiver` and forwards all deposits to a configured recipient capability. Used to redirect tokens transparently.
- `PrivateReceiverForwarder.cdc` — Like `TokenForwarding` but the `deposit` is `access(contract)`, so only a co-deployed `Sender` resource (held by an admin) can push tokens in. Used for privacy-preserving airdrops.
- `FungibleTokenSwitchboard.cdc` — A `Switchboard` resource that acts as a single `Receiver` capable of routing deposits to multiple underlying vaults by vault type. The generic receiver path `/public/GenericFTReceiver` points here.
- `FungibleTokenMetadataViews.cdc` — Defines the `FTView`, `FTDisplay`, `FTVaultData`, and `TotalSupply` metadata view structs. `FTVaultData` is particularly important: it carries storage/public paths and a `createEmptyVault` function, enabling generic account setup transactions without importing the token contract directly.

### flow.json and contract aliases

`flow.json` maps contract names to source files and **network aliases** (deployed addresses). In Cadence source, contracts are imported by name string (e.g. `import "FungibleToken"`). The Flow CLI resolves these names to addresses at deploy/test time using the aliases. The `testing` network alias points to `0000000000000007` (the test account). When adding a new contract, it needs an entry in both `flow.json` `contracts` and `deployments` sections.

### Imports folder

`imports/` contains pinned copies of external contracts (`MetadataViews`, `ViewResolver`, `FlowToken`, etc.) fetched via `flow dependencies`. These are resolved by `flow.json` aliases and should not be edited manually.

### Generic transfer transactions

`transactions/generic_transfer_with_address.cdc` and `transactions/generic_transfer_with_paths.cdc` allow transferring any FT without a token-specific import. They read `FTVaultData` from the contract's metadata views to find the correct storage/receiver paths. `generic_transfer_with_address.cdc` includes a post-withdraw type assertion to guard against malicious tokens returning incorrect metadata (see `contracts/test/MaliciousToken.cdc` for the attack this prevents).

## Key Patterns

- **Vault deposit**: Always force-cast the incoming vault to the concrete type before incrementing balance: `let vault <- from as! @ExampleToken.Vault`. The interface pre-condition already type-checks, so the cast is guaranteed to succeed.
- **Capability access control**: Use Cadence entitlements (`access(Owner)`, `access(FungibleToken.Withdraw)`) rather than capability types alone to restrict sensitive functions.
- **Account setup**: The recommended setup transaction is `transactions/metadata/setup_account_from_address.cdc`, which uses `FTVaultData` to set up an account for any token without importing that token's contract.
