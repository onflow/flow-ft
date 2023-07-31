import Test

pub let blockchain = Test.newEmulatorBlockchain()
pub let admin = blockchain.createAccount()
pub let recipient = blockchain.createAccount()

pub fun setup() {
    blockchain.useConfiguration(Test.Configuration({
        "FungibleTokenMetadataViews": admin.address,
        "ExampleToken": admin.address,
        "FungibleTokenSwitchboard": admin.address
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
    let getSupportedVaultsCode = Test.readFile("../transactions/scripts/get_supported_vault_types.cdc")

    var scriptResult = blockchain.executeScript(
        getSupportedVaultsCode,
        [recipient.address, /public/GenericFTReceiver]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    var types = (scriptResult.returnValue as! {Type: Bool}?)!
    Test.assertEqual(0, types.keys.length)

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

    types = (scriptResult.returnValue as! {Type: Bool}?)!
    Test.assertEqual(1, types.keys.length)
}