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

/// Verifies that changeRecipient succeeds when the old recipient's vault has
/// been removed, making the stored capability stale. The ForwarderRecipientUpdated
/// event should emit nil for oldRecipient and the correct address for newRecipient.
access(all)
fun testChangeRecipientSucceedsWhenOldCapabilityIsStale() {
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

    // Remove initialRecipient's vault, making forwarderOwner's stored capability stale
    txResult = executeTransaction(
        "transactions/unload_example_token_vault.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/tokenForwarder/change_recipient.cdc",
        [newRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    let recipientUpdatedEvents = Test.eventsOfType(Type<TokenForwarding.ForwarderRecipientUpdated>())
    Test.assertEqual(1, recipientUpdatedEvents.length)

    let recipientUpdatedEvent = recipientUpdatedEvents[0] as! TokenForwarding.ForwarderRecipientUpdated
    Test.assertEqual(nil, recipientUpdatedEvent.oldRecipient)
    Test.assertEqual(newRecipient.address, recipientUpdatedEvent.newRecipient!)
}

/// Verifies that tokens sent to a Forwarder are routed to the current recipient
/// after changeRecipient has been called with a stale old capability.
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

    // Remove initialRecipient's vault, making the stored capability stale
    txResult = executeTransaction(
        "transactions/unload_example_token_vault.cdc",
        [],
        initialRecipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/tokenForwarder/change_recipient.cdc",
        [newRecipient.address],
        forwarderOwner
    )
    Test.expect(txResult, Test.beSucceeded())

    // Tokens deposited to forwarderOwner's receiver should be forwarded to newRecipient
    txResult = executeTransaction(
        "../transactions/mint_tokens.cdc",
        [forwarderOwner.address, 100.0],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    let scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [newRecipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())
    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(100.0, balance)
}

/// Verifies that getSupportedVaultTypes on a Forwarder returns the vault type
/// of the underlying recipient, not the Forwarder's own concrete type.
access(all)
fun testGetSupportedVaultTypesForSingleForwarder() {
    let vaultAccount = Test.createAccount()
    let forwarderAccount = Test.createAccount()

    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        vaultAccount
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/tokenForwarder/create_forwarder.cdc",
        [vaultAccount.address],
        forwarderAccount
    )
    Test.expect(txResult, Test.beSucceeded())

    let scriptResult = executeScript(
        "../transactions/scripts/get_supported_vault_types.cdc",
        [forwarderAccount.address, /public/exampleTokenReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)
}

/// Verifies that getSupportedVaultTypes correctly resolves the underlying vault
/// type through a chain of Forwarders (Forwarder → Forwarder → Vault),
/// returning the depositable vault type rather than an intermediate forwarder type.
access(all)
fun testGetSupportedVaultTypesForChainedForwarders() {
    let vaultAccount = Test.createAccount()
    let forwarderA = Test.createAccount()
    let forwarderB = Test.createAccount()

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

    let scriptResult = executeScript(
        "../transactions/scripts/get_supported_vault_types.cdc",
        [forwarderB.address, /public/exampleTokenReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)

    Test.assert(
        supportedTypes[Type<@TokenForwarding.Forwarder>()] == nil,
        message: "getSupportedVaultTypes must not return TokenForwarding.Forwarder as a supported type"
    )
}
