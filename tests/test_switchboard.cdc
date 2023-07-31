import Test

pub let blockchain = Test.newEmulatorBlockchain()
pub let admin = blockchain.createAccount()
pub let recipient = blockchain.createAccount()

pub fun setup() {
    blockchain.useConfiguration(Test.Configuration({
        "FungibleTokenMetadataViews": admin.address,
        "ExampleToken": admin.address,
        "FungibleTokenSwitchboard": admin.address,
        "TokenForwarding": admin.address
    }))

    var code = Test.readFile("../contracts/FungibleTokenMetadataViews.cdc")
    var err = blockchain.deployContract(
        name: "FungibleTokenMetadataViews",
        code: code,
        account: admin,
        arguments: []
    )
    Test.expect(err, Test.beNil())

    code = Test.readFile("../contracts/ExampleToken.cdc")
    err = blockchain.deployContract(
        name: "ExampleToken",
        code: code,
        account: admin,
        arguments: []
    )

    Test.expect(err, Test.beNil())

    code = Test.readFile("../contracts/FungibleTokenSwitchboard.cdc")
    err = blockchain.deployContract(
        name: "FungibleTokenSwitchboard",
        code: code,
        account: admin,
        arguments: []
    )

    Test.expect(err, Test.beNil())

    code = Test.readFile("../contracts/utility/TokenForwarding.cdc")
    err = blockchain.deployContract(
        name: "TokenForwarding",
        code: code,
        account: admin,
        arguments: []
    )

    Test.expect(err, Test.beNil())
}

pub fun testSetupSwitchboard() {
    var code = Test.readFile("../transactions/setup_account.cdc")
    var tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: []
    )
    var txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())


    code = Test.readFile("../transactions/switchboard/setup_account.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: []
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the newly-setup switchboard cannot accept any types
    let getSupportedVaultsCode = Test.readFile("scripts/get_supported_vault_types.cdc")

    var scriptResult = blockchain.executeScript(
        getSupportedVaultsCode,
        [recipient.address, /public/GenericFTReceiver]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    var numTypes = (scriptResult.returnValue as! Int?)!
    Test.assertEqual(0, numTypes)

    code = Test.readFile("../transactions/switchboard/add_vault_capability.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: []
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard can now accept one vault type
    scriptResult = blockchain.executeScript(
        getSupportedVaultsCode,
        [recipient.address, /public/GenericFTReceiver]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    numTypes = (scriptResult.returnValue as! Int?)!
    Test.assertEqual(1, numTypes)
}

pub fun testUseSwitchboard() {
    var code = Test.readFile("../transactions/switchboard/safe_transfer_tokens_v2.cdc")
    var tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 10.0]
    )
    var txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    code = Test.readFile("../transactions/switchboard/transfer_tokens.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic]
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard account has a balance of 20.0
    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    let scriptResult = blockchain.executeScript(
        code,
        [recipient.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(20.0, balance)

}

pub fun testRemoveVaultTypeFromSwitchboard() {
    var code = Test.readFile("../transactions/switchboard/remove_vault_capability.cdc")
    var tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [/public/exampleTokenReceiver]
    )
    var txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    code = Test.readFile("../transactions/switchboard/transfer_tokens.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic]
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beFailed())

    let getSupportedVaultsCode = Test.readFile("scripts/get_supported_vault_types.cdc")

    // Test that the switchboard can now accept zero vault types
    let scriptResult = blockchain.executeScript(
        getSupportedVaultsCode,
        [recipient.address, /public/GenericFTReceiver]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let numTypes = (scriptResult.returnValue as! Int?)!
    Test.assertEqual(0, numTypes)

}

pub fun testUseSwitchboardWithForwarder() {
    var code = Test.readFile("../transactions/create_forwarder.cdc")
    var tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [admin.address]
    )
    var txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    code = Test.readFile("../transactions/switchboard/batch_add_vault_wrapper_capabilities.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [recipient.address]
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard can now accept one vault types
    let getSupportedVaultsCode = Test.readFile("scripts/get_supported_vault_types.cdc")
    var scriptResult = blockchain.executeScript(
        getSupportedVaultsCode,
        [recipient.address, /public/GenericFTReceiver]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let numTypes = (scriptResult.returnValue as! Int?)!
    Test.assertEqual(1, numTypes)

    code = Test.readFile("../transactions/switchboard/transfer_tokens.cdc")
    tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 10.0, /public/fungibleTokenSwitchboardPublic]
    )
    txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the switchboard account has a balance of 20.0
    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    scriptResult = blockchain.executeScript(
        code,
        [recipient.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(20.0, balance)
}