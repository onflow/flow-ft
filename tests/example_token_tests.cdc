import Test
import "test_helpers.cdc"

// access(all) let blockchain = Test.newEmulatorBlockchain()
access(all) let admin = blockchain.createAccount()
access(all) let recipient = blockchain.createAccount()

access(all) fun setup() {
    blockchain.useConfiguration(
        Test.Configuration(
            addresses: {
                "ViewResolver": admin.address,
                "FungibleToken": admin.address,
                "NonFungibleToken": admin.address,
                "MetadataViews": admin.address,
                "FungibleTokenMetadataViews": admin.address,
                "ExampleToken": admin.address
            }
        )
    )

    deploy("ViewResolver", admin, "../contracts/utility/ViewResolver.cdc")
    deploy("FungibleToken", admin, "../contracts/FungibleToken-v2.cdc")
    deploy("NonFungibleToken", admin, "../contracts/utility/NonFungibleToken.cdc")
    deploy("MetadataViews", admin, "../contracts/utility/MetadataViews.cdc")
    deploy("FungibleTokenMetadataViews", admin, "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", admin, "../contracts/ExampleToken-v2.cdc")
}

access(all) fun testTokensInitializedEventEmitted() {
    let addrString = admin.address.toString()
    let identifier = "A.".concat(addrString.slice(from: 2, upTo: addrString.length)).concat(".").concat("ExampleToken").concat(".").concat("TokensInitialized")
    let typ = CompositeType(identifier)
        ?? panic("Problem constructing CompositeType")
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)
}

access(all) fun testGetTotalSupply() {
    let code = Test.readFile("../transactions/scripts/get_supply.cdc")
    let scriptResult = blockchain.executeScript(code, [])

    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1000.0, totalSupply)
}

access(all) fun testGetAdminBalance() {
    let code = Test.readFile("../transactions/scripts/get_balance.cdc")
    let scriptResult = blockchain.executeScript(
        code,
        [admin.address]
    )

    Test.expect(scriptResult, Test.beSucceeded())

    let balance = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1000.0, balance)
}

access(all) fun testSetupAccount() {
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

access(all) fun testMintTokens() {
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
    var typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "TokensMinted"))!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "MinterCreated"))!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "TokensDeposited"))!
    Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    // Test that the totalSupply increased by the amount of minted tokens
    code = Test.readFile("../transactions/scripts/get_supply.cdc")
    let scriptResult = blockchain.executeScript(code, [])

    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = (scriptResult.returnValue as! UFix64?)!
    Test.assertEqual(1250.0, totalSupply)
}

access(all) fun testTransferTokens() {
    var code = Test.readFile("../transactions/transfer_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [recipient.address],
        signers: [recipient],
        arguments: [50.0, admin.address]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    var typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "TokensDeposited"))!
    Test.assertEqual(2, blockchain.eventsOfType(typ).length)

    typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "TokensWithdrawn"))!
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

access(all) fun testTransferTokenAmountGreaterThanBalance() {
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

access(all) fun testBurnTokens() {
    var code = Test.readFile("../transactions/burn_tokens.cdc")
    let tx = Test.Transaction(
        code: code,
        authorizers: [admin.address],
        signers: [admin],
        arguments: [50.0]
    )
    let txResult = blockchain.executeTransaction(tx)

    Test.expect(txResult, Test.beSucceeded())

    // var typ = CompositeType(buildTypeIdentifier(admin, "ExampleToken", "BurnerCreated"))!
    // Test.assertEqual(1, blockchain.eventsOfType(typ).length)

    var typ = CompositeType(buildTypeIdentifier(admin, "FungibleToken", "Burrn"))!
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

access(all) fun testVaultTypes() {
    let code = Test.readFile("./scripts/get_views.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    log(scriptResult.error)

    Test.expect(scriptResult, Test.beSucceeded())
}

access(all) fun testGetVaultDisplay() {
    let code = Test.readFile("./scripts/get_vault_display.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

access(all) fun testGetVaultData() {
    let code = Test.readFile("./scripts/get_vault_data.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

access(all) fun testGetTokenMetadata() {
    let code = Test.readFile("./scripts/get_token_metadata.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}

access(all) fun testGetUnsupportedViewType() {
    let code = Test.readFile("./scripts/get_unsupported_view.cdc")
    let scriptResult = blockchain.executeScript(code, [recipient.address])

    Test.expect(scriptResult, Test.beSucceeded())
}
