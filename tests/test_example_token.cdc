import Test

pub let blockchain = Test.newEmulatorBlockchain()
pub let admin = blockchain.createAccount()
pub let recipient = blockchain.createAccount()

pub fun setup() {
    blockchain.useConfiguration(Test.Configuration({
        "FungibleTokenMetadataViews": admin.address,
        "ExampleToken": admin.address
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
}

pub fun testTokensInitializedEventEmitted() {
    let typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensInitialized")!

    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

pub fun testGetTotalSupply() {
    let code = Test.readFile("../transactions/scripts/get_supply.cdc")
    let scriptResult = blockchain.executeScript(code, [])

    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1000.0, totalSupply)
}

pub fun testGetAdminBalance() {
    let code = Test.readFile("../transactions/scripts/get_balance.cdc")
    let scriptResult = blockchain.executeScript(
        code,
        [admin.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1000.0, balance)
}

pub fun testSetupAccount() {
    var code = Test.readFile("../transactions/setup_account.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: []
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the newly-setup account has a balance of 0.0
    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    let scriptResult = blockchain.executeScript(
        code,
        [recipient.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(0.0, balance)
}

pub fun testMintTokens() {
    var code = Test.readFile("../transactions/mint_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 250.0]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // Test that the proper events were emitted
    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensMinted")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.MinterCreated")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensDeposited")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    // Test that the totalSupply increased by the amount of minted tokens
    code = Test.readFile("../transactions/scripts/get_supply.cdc")
    let scriptResult = blockchain.executeScript(code, [])

    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1250.0, totalSupply)
}

pub fun testTransferTokens() {
    var code = Test.readFile("../transactions/transfer_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [50.0, admin.address]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensDeposited")!
    Test.assertEqual(2, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensWithdrawn")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    var scriptResult = blockchain.executeScript(
        code,
        [recipient.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    var balance = (scriptResult.returnValue as! UFix64?)!
    // 250.0 tokens were previously minted to the recipient
    Test.assertEqual(200.0, balance)

    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    scriptResult = blockchain.executeScript(
        code,
        [admin.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    // The admin had initially 1000.0 tokens (initial supply)
    balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1050.0, balance)
}

pub fun testTransferTokenAmountGreaterThanBalance() {
    var code = Test.readFile("../transactions/transfer_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [1550.0, admin.address]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beFailed())
    Test.assertEqual(
        "Amount withdrawn must be less than or equal than the balance of the Vault",
        txResult.error!.message.slice(from: 174, upTo: 247)
    )
}

pub fun testBurnTokens() {
    var code = Test.readFile("../transactions/burn_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [50.0]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.BurnerCreated")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensBurned")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    code = Test.readFile("../transactions/scripts/get_balance.cdc")
    let scriptResult = blockchain.executeScript(
        code,
        [admin.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    // The admin should now have the initial supply of 1000.0 tokens
    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1000.0, balance)
}

pub fun testVaultTypes() {
    let code = Test.readFile("./scripts/get_views.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

pub fun testGetVaultDisplay() {
    let code = Test.readFile("./scripts/get_vault_display.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

pub fun testGetVaultData() {
    let code = Test.readFile("./scripts/get_vault_data.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

pub fun testGetTokenMetadata() {
    let code = Test.readFile("./scripts/get_token_metadata.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

pub fun testGetUnsupportedViewType() {
    let code = Test.readFile("./scripts/get_unsupported_view.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}
