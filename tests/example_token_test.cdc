import Test
import BlockchainHelpers
import "test_helpers.cdc"
import "FungibleTokenMetadataViews"
import "ExampleToken"
import "FungibleToken"

access(all) let admin = Test.getAccount(0x0000000000000007)
access(all) let service = Test.getAccount(0x0000000000000001)
access(all) let recipient = Test.createAccount()

access(all) let exampleTokenIdentifier = "A.0000000000000007.ExampleToken.Vault"
access(all) let maliciousTokenIdentifier = "A.0000000000000007.MaliciousToken.Vault"

access(all)
fun setup() {
    deploy("Burner", "../contracts/utility/Burner.cdc")
    deploy("FungibleToken", "../contracts/FungibleToken.cdc")
    deploy("FungibleTokenMetadataViews", "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", "../contracts/ExampleToken.cdc")
    deploy("MaliciousToken", "../contracts/test/MaliciousToken.cdc")
}

access(all)
fun testGetTotalSupply() {
    let scriptResult = executeScript(
        "../transactions/scripts/get_supply.cdc",
        []
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1000.0, totalSupply)
}

access(all)
fun testGetAdminBalance() {
    let scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [admin.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1000.0, balance)
}

access(all)
fun testSetupAccount() {
    let txResult = executeTransaction(
        "../transactions/setup_account.cdc",
        [],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the newly-setup account has a balance of 0.0
    let scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(0.0, balance)
}

access(all)
fun testMintTokens() {
    let txResult = executeTransaction(
        "../transactions/mint_tokens.cdc",
        [recipient.address, 250.0],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    // Test that the proper events were emitted
    var typ = Type<ExampleToken.TokensMinted>()
    var events = Test.eventsOfType(typ)
    Test.assertEqual(2, events.length)

    let tokensMintedEvent = events[1] as! ExampleToken.TokensMinted
    Test.assertEqual(250.0, tokensMintedEvent.amount)

    typ = Type<FungibleToken.Deposited>()
    events = Test.eventsOfType(typ)

    let tokensDepositedEvent = events[events.length - 1] as! FungibleToken.Deposited
    Test.assertEqual(250.0, tokensDepositedEvent.amount)
    Test.assertEqual(recipient.address, tokensDepositedEvent.to!)
    Test.assertEqual(exampleTokenIdentifier, tokensDepositedEvent.type)
    Test.assertEqual(250.0, tokensDepositedEvent.balanceAfter!)

    // Test that the totalSupply increased by the amount of minted tokens
    let scriptResult = executeScript(
        "../transactions/scripts/get_supply.cdc",
        []
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let totalSupply = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1250.0, totalSupply)
}

access(all)
fun testTransferTokens() {
    let txResult = executeTransaction(
        "../transactions/transfer_tokens.cdc",
        [50.0, admin.address],
        recipient
    )
    Test.expect(txResult, Test.beSucceeded())

    var typ = Type<FungibleToken.Deposited>()
    Test.assertEqual(2, Test.eventsOfType(typ).length)

    typ = Type<FungibleToken.Withdrawn>()
    let events = Test.eventsOfType(typ)
    Test.assertEqual(1, events.length)

    let tokensWithdrawnEvent = events[0] as! FungibleToken.Withdrawn
    Test.assertEqual(exampleTokenIdentifier, tokensWithdrawnEvent.type)
    Test.assertEqual(50.0, tokensWithdrawnEvent.amount)
    Test.assertEqual(recipient.address, tokensWithdrawnEvent.from!)
    Test.assertEqual(200.0, tokensWithdrawnEvent.balanceAfter!)

    var scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    var balance = scriptResult.returnValue! as! UFix64
    // 250.0 tokens were previously minted to the recipient
    Test.assertEqual(200.0, balance)

    scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [admin.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    // The admin had initially 1000.0 tokens (initial supply)
    balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1050.0, balance)
}

access(all)
fun testTransferTokenAmountGreaterThanBalance() {
    let txResult = executeTransaction(
        "../transactions/transfer_tokens.cdc",
        [1550.0, admin.address],
        recipient
    )
    Test.expect(txResult, Test.beFailed())
    Test.assertError(
        txResult,
        errorMessage: "FungibleToken.Vault.withdraw: Cannot withdraw tokens! The amount requested to be withdrawn (1550.00000000) is greater than the balance of the Vault (200.00000000)."
    )
}

access(all)
fun testBurnTokens() {
    var txResult = executeTransaction(
        "../transactions/burn_tokens.cdc",
        [exampleTokenIdentifier, 50.0],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    let type = Type<FungibleToken.Burned>()
    let events = Test.eventsOfType(type)
    Test.assertEqual(1, events.length)

    let tokensBurnedEvent = events[0] as! FungibleToken.Burned
    Test.assertEqual(50.0, tokensBurnedEvent.amount)
    Test.assertEqual(exampleTokenIdentifier, tokensBurnedEvent.type)

    var scriptResult = executeScript(
        "../transactions/scripts/get_balance.cdc",
        [admin.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    // The admin should now have the initial supply of 1000.0 tokens
    var balance = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1000.0, balance)

    txResult = executeTransaction(
        "transactions/burn_array.cdc",
        [10.0, 5],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    scriptResult = executeScript(
        "../transactions/scripts/get_supply.cdc",
        []
    )
    Test.expect(scriptResult, Test.beSucceeded())

    var totalSupply = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1150.0, totalSupply)

        txResult = executeTransaction(
        "transactions/burn_dict.cdc",
        [10.0, 5],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    scriptResult = executeScript(
        "../transactions/scripts/get_supply.cdc",
        []
    )
    Test.expect(scriptResult, Test.beSucceeded())

    totalSupply = scriptResult.returnValue! as! UFix64
    Test.assertEqual(1100.0, totalSupply)

    txResult = executeTransaction(
        "transactions/burn_optional.cdc",
        [],
        admin
    )
    Test.expect(txResult, Test.beSucceeded())

    scriptResult = executeScript(
        "../transactions/scripts/get_supply.cdc",
        []
    )
    Test.expect(scriptResult, Test.beSucceeded())

    totalSupply = scriptResult.returnValue! as! UFix64
    Test.assertEqual(200.0, totalSupply)
}

access(all)
fun testVaultTypes() {
    let scriptResult = executeScript(
        "../transactions/metadata/scripts/get_views.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let supportedViews = scriptResult.returnValue! as! [Type]
    let expectedViews = [
        Type<FungibleTokenMetadataViews.FTView>(),
        Type<FungibleTokenMetadataViews.FTDisplay>(),
        Type<FungibleTokenMetadataViews.FTVaultData>(),
        Type<FungibleTokenMetadataViews.TotalSupply>()
    ]
    Test.assertEqual(expectedViews, supportedViews)
}

access(all)
fun testGetVaultDisplay() {
    let scriptResult = executeScript(
        "scripts/get_vault_display.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let ftDisplay = scriptResult.returnValue! as! FungibleTokenMetadataViews.FTDisplay
    Test.assertEqual("Example Fungible Token", ftDisplay.name)
    Test.assertEqual("EFT", ftDisplay.symbol)
    Test.assertEqual(
        "This fungible token is used as an example to help you develop your next FT #onFlow.",
        ftDisplay.description
    )
    Test.assertEqual(
        "https://example-ft.onflow.org",
        ftDisplay.externalURL!.url
    )
    Test.assertEqual(
        "https://twitter.com/flow_blockchain",
        ftDisplay.socials["twitter"]!.url
    )
    Test.assertEqual(
        "https://assets.website-files.com/5f6294c0c7a8cdd643b1c820/5f6294c0c7a8cda55cb1c936_Flow_Wordmark.svg",
        ftDisplay.logos.items[0].file.uri()
    )
}

access(all)
fun testGetVaultData() {
    let scriptResult = executeScript(
        "scripts/get_vault_data.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())
}

access(all)
fun testGetTokenMetadata() {
    let scriptResult = executeScript(
        "scripts/get_token_metadata.cdc",
        [recipient.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())
}

access(all)
fun testGetUnsupportedViewType() {
    let scriptResult = executeScript(
        "scripts/get_unsupported_view.cdc",
        [recipient.address, Type<String>()]
    )
    Test.expect(scriptResult, Test.beSucceeded())

    let view = scriptResult.returnValue
    Test.expect(view, Test.beNil())
}


access(all) fun testTransferTokenFromMaliciousContract() {

    var txResult = executeTransaction(
        "../transactions/generic_transfer_with_address.cdc",
        [0.1, admin.address, maliciousTokenIdentifier],
        recipient
    )
    Test.expect(txResult, Test.beFailed())
    Test.assertError(
        txResult,
        errorMessage: "The Vault that was withdrawn to transfer is not the type that was requested!"
    )
}
