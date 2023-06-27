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

pub fun testTokensInitialized() {
    let typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensInitialized")!

    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

pub fun testSetupRecipientAccount() {
    let code = Test.readFile("../transactions/setup_account.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: []
    )

    let result = blockchain.executeTransaction(tx)

    Test.expect(result, Test.beSucceeded())
}

pub fun testMintTokens() {
    let code = Test.readFile("../transactions/mint_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [recipient.address, 250.0]
    )

    let result = blockchain.executeTransaction(tx)

    Test.expect(result, Test.beSucceeded())

    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensMinted")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.MinterCreated")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensDeposited")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

pub fun testTransferTokens() {
    let code = Test.readFile("../transactions/transfer_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [50.0, admin.address]
    )

    let result = blockchain.executeTransaction(tx)

    Test.expect(result, Test.beSucceeded())

    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensDeposited")!
    Test.assertEqual(2, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensWithdrawn")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

pub fun testBurnTokens() {
    let code = Test.readFile("../transactions/burn_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [5.0]
    )

    let result = blockchain.executeTransaction(tx)

    Test.expect(result, Test.beSucceeded())

    var typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.BurnerCreated")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType("A.01cf0e2f2f715450.ExampleToken.TokensBurned")!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

pub fun testVaultTypes() {
    let code = Test.readFile("./scripts/get_views.cdc")
    let result = blockchain.executeScript(code, [recipient.address])
    let viewIDs = (result.returnValue as! [String]?)!

    Test.expect(result, Test.beSucceeded())

    let expected = [
        "A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTView",
        "A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTDisplay",
        "A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTVaultData"
    ]
    Test.assertEqual(expected, viewIDs)
}

pub fun testGetVaultDisplay() {
    let code = Test.readFile("./scripts/get_vault_display.cdc")
    let result = blockchain.executeScript(code, [recipient.address])

    Test.expect(result, Test.beSucceeded())
}

pub fun testGetVaultData() {
    let code = Test.readFile("./scripts/get_vault_data.cdc")
    let result = blockchain.executeScript(code, [recipient.address])

    Test.expect(result, Test.beSucceeded())
}

pub fun testGetTokenMetadata() {
    let code = Test.readFile("./scripts/get_token_metadata.cdc")
    let result = blockchain.executeScript(code, [recipient.address])

    Test.expect(result, Test.beSucceeded())
}

pub fun testGetUnsupportedViewType() {
    let code = Test.readFile("./scripts/get_unsupported_view.cdc")
    let result = blockchain.executeScript(code, [recipient.address])

    Test.expect(result, Test.beSucceeded())
}
