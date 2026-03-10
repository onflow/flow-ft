import Test
import BlockchainHelpers
import "test_helpers.cdc"
import "FungibleToken"
import "ExampleToken"
import "TokenForwarding"
import "FungibleTokenMetadataViews"

access(all) let admin = Test.getAccount(0x0000000000000007)

access(all)
fun setup() {
    deploy("Burner", "../contracts/utility/Burner.cdc")
    deploy("FungibleToken", "../contracts/FungibleToken.cdc")
    deploy("FungibleTokenMetadataViews", "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", "../contracts/ExampleToken.cdc")
    deploy("TokenForwarding", "../contracts/utility/TokenForwarding.cdc")
}

/// Bug fix #1: changeRecipient should succeed even when the old recipient
/// capability is stale (e.g. the previous recipient deleted their vault).
///
/// Before the fix, the force-unwrap on `self.recipient.borrow()!` would panic,
/// permanently bricking the Forwarder. After the fix, an optional borrow is used
/// so changeRecipient always succeeds regardless of the old cap's validity.
access(all)
fun testChangeRecipientSucceedsWhenOldCapabilityIsStale() {
    let initialRecipient = Test.createAccount()
    let forwarderOwner = Test.createAccount()
    let newRecipient = Test.createAccount()

    // Both the initial and new recipients set up ExampleToken vaults
    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        newRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // forwarderOwner creates a Forwarder pointing to initialRecipient
    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [initialRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    // initialRecipient destroys their vault, making forwarderOwner's
    // stored recipient capability stale
    txResult = executeTransaction(
        "transactions/unload_example_token_vault.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // changeRecipient must succeed even though the old capability is now stale.
    // Before the fix this would panic with a force-unwrap failure.
    txResult = executeTransaction(
        "../transactions/tokenForwarder/change_recipient.cdc",
        [newRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    // The event should record nil for oldRecipient (stale cap) and the correct newRecipient
    let recipientUpdatedEvents = Test.eventsOfType(Type<TokenForwarding.ForwarderRecipientUpdated>())
    Test.assertEqual(1, recipientUpdatedEvents.length)

    let recipientUpdatedEvent = recipientUpdatedEvents[0] as! TokenForwarding.ForwarderRecipientUpdated
    Test.assertEqual(nil, recipientUpdatedEvent.oldRecipient)
    Test.assertEqual(newRecipient.address, recipientUpdatedEvent.newRecipient!)
}

/// Verify that after changeRecipient, the Forwarder actually routes tokens
/// to the new recipient and not the old (now-stale) one.
access(all)
fun testTokensForwardedToNewRecipientAfterChange() {
    let initialRecipient = Test.createAccount()
    let forwarderOwner = Test.createAccount()
    let newRecipient = Test.createAccount()

    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        newRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [initialRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    // Destroy the initial recipient's vault to make the capability stale
    txResult = executeTransaction(
        "transactions/unload_example_token_vault.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // Update the Forwarder to point to the new (valid) recipient
    txResult = executeTransaction(
        "../transactions/tokenForwarder/change_recipient.cdc",
        [newRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    // Mint tokens directly to the forwarderOwner's receiver path (the Forwarder).
    // They should be forwarded through to newRecipient.
    txResult = executeTransaction(
        "../transactions/mint_tokens.cdc",
        [forwarderOwner.address, 100.0],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    // newRecipient should have received the forwarded tokens
    let scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [newRecipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())
    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(100.0, balance)
}

/// Bug fix #2: getSupportedVaultTypes on a chained Forwarder should return the
/// underlying vault type, not the intermediate forwarder's concrete type.
///
/// Before the fix, `getSupportedVaultTypes` returned `{vaultRef.getType(): true}`,
/// which yields `TokenForwarding.Forwarder` when forwarders are chained rather than
/// the actual depositable vault type. After the fix it delegates to
/// `vaultRef.getSupportedVaultTypes()`, propagating the correct type through the chain.
access(all)
fun testGetSupportedVaultTypesForSingleForwarder() {
    let vaultAccount = Test.createAccount()
    let forwarderAccount = Test.createAccount()

    // vaultAccount sets up a real ExampleToken vault
    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        vaultAccount
    )
    Test.expect(txResult, Test.beSucceeded())

    // forwarderAccount creates a Forwarder pointing to vaultAccount's vault
    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [vaultAccount.address],
        forwarderAccount
    )
    Test.expect(txResult, Test.beSucceeded())

    // getSupportedVaultTypes on the Forwarder must return the underlying vault type
    let scriptResult = executeScript(
        "../transactions/scripts/get_supported_vault_types.cdc",
        [forwarderAccount.address, /public/exampleTokenReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)
}

access(all)
fun testGetSupportedVaultTypesForChainedForwarders() {
    let vaultAccount = Test.createAccount()
    let forwarderA = Test.createAccount()
    let forwarderB = Test.createAccount()

    // vaultAccount sets up a real ExampleToken vault
    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        vaultAccount
    )
    Test.expect(txResult, Test.beSucceeded())

    // forwarderA: Forwarder → vaultAccount's ExampleToken.Vault
    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [vaultAccount.address],
        forwarderA
    )
    Test.expect(txResult, Test.beSucceeded())

    // forwarderB: Forwarder → forwarderA's Forwarder (chained)
    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [forwarderA.address],
        forwarderB
    )
    Test.expect(txResult, Test.beSucceeded())

    // Before the fix, querying forwarderB returned {TokenForwarding.Forwarder: true}.
    // After the fix, it must recursively resolve to {ExampleToken.Vault: true}.
    let scriptResult = executeScript(
        "../transactions/scripts/get_supported_vault_types.cdc",
        [forwarderB.address, /public/exampleTokenReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)

    // Sanity check: the incorrect pre-fix result is NOT present
    Test.assert(
        supportedTypes[Type<@TokenForwarding.Forwarder>()] == nil,
        message: "getSupportedVaultTypes must not return TokenForwarding.Forwarder as a supported type"
    )
}
