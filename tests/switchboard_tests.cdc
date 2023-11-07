import Test
import BlockchainHelpers
import "FungibleTokenMetadataViews"
import "ExampleToken"
import "FungibleTokenSwitchboard"

access(all) let admin = Test.getAccount(0x0000000000000007)
access(all) let recipient = Test.createAccount()

access(all)
fun setup() {
    var err = Test.deployContract(
        name: "FungibleTokenMetadataViews",
        path: "../contracts/FungibleTokenMetadataViews.cdc",
        arguments: []
    )
    Test.expect(err, Test.beNil())

    err = Test.deployContract(
        name: "ExampleToken",
        path: "../contracts/ExampleToken-v2.cdc",
        arguments: []
    )
    Test.expect(err, Test.beNil())

    err = Test.deployContract(
        name: "FungibleTokenSwitchboard",
        path: "../contracts/FungibleTokenSwitchboard.cdc",
        arguments: []
    )
    Test.expect(err, Test.beNil())

    err = Test.deployContract(
        name: "TokenForwarding",
        path: "../contracts/utility/TokenForwarding.cdc",
        arguments: []
    )
    Test.expect(err, Test.beNil())
}

access(all)
fun testSetupSwitchboard() {
    var txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/switchboard/setup_account.cdc",
        [],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the newly-setup switchboard cannot accept any types
    var scriptResult = executeScript(
        "scripts/get_supported_vault_types.cdc",
        [recipient.address, /public/GenericFTReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    var supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    Test.expect(supportedTypes, Test.beEmpty())

    txResult = executeTransaction(
        "../transactions/switchboard/add_vault_capability.cdc",
        [],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard can now accept one vault type
    scriptResult = executeScript(
        "scripts/get_supported_vault_types.cdc",
        [recipient.address, /public/GenericFTReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)
}

access(all)
fun testUseSwitchboard() {
    var txResult = executeTransaction(
        "../transactions/switchboard/safe_transfer_tokens_v2.cdc",
        [recipient.address, 10.0],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/switchboard/transfer_tokens.cdc",
        [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard account has a balance of 20.0
    let scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(20.0, balance)
}

access(all)
fun testRemoveVaultTypeFromSwitchboard() {
    var txResult = executeTransaction(
        "../transactions/switchboard/remove_vault_capability.cdc",
        [/public/exampleTokenReceiver],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/switchboard/transfer_tokens.cdc",
        [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic],
        admin
    )
    Test.expect(txResult, Test.beFailed())

    // Test that the switchboard can now accept zero vault types
    let scriptResult = executeScript(
        "scripts/get_supported_vault_types.cdc",
        [recipient.address, /public/GenericFTReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    Test.expect(supportedTypes, Test.beEmpty())
}

access(all)
fun testUseSwitchboardWithForwarder() {
    var txResult = executeTransaction(
        "../transactions/create_forwarder.cdc",
        [admin.address],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    txResult = executeTransaction(
        "../transactions/switchboard/batch_add_vault_wrapper_capabilities.cdc",
        [recipient.address],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard can now accept one vault types
    var scriptResult = executeScript(
        "scripts/get_supported_vault_types.cdc",
        [recipient.address, /public/GenericFTReceiver]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedTypes = scriptResult.returnValue! as! {Type: Bool}
    let expectedTypes = {Type<@ExampleToken.Vault>(): true}
    Test.assertEqual(expectedTypes, supportedTypes)

    txResult = executeTransaction(
        "../transactions/switchboard/transfer_tokens.cdc",
        [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard account has a balance of 20.0
    scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(20.0, balance)
}
