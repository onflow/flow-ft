import Test
import "test_helpers.cdc"

// access(all) let blockchain = Test.newEmulatorBlockchain()
access(all) let sourceAccount = blockchain.createAccount()
access(all) let accounts: {String: Test.TestAccount} = {}

access(all) let exampleToken = "ExampleToken"

/* Test Cases */

access(all) fun testSetupAccountUsingFTView() {
    let alice = blockchain.createAccount()
    let bob = blockchain.createAccount()

    setupExampleToken(alice)
    let aliceBalance = getExampleTokenBalance(alice)
    txExecutor("metadata/setup_account_from_vault_reference.cdc", [bob], [alice.address, /public/exampleTokenVault], nil, nil)
    let bobBalance = getExampleTokenBalance(alice)

    Test.assertEqual(0.0, aliceBalance)
    Test.assertEqual(0.0, bobBalance)
}

access(all) fun testRetrieveVaultDisplayInfo() {
    let alice = blockchain.createAccount()

    setupExampleToken(alice)
    let result = scriptExecutor("test/example_token_vault_display_strict_equal.cdc", [alice.address])! as! Bool

    Test.assertEqual(true, result)
}


/* Transaction Helpers */

access(all) fun setupExampleToken(_ acct: Test.TestAccount) {
    txExecutor("setup_account.cdc", [acct], [], nil, nil)
}

/* Script Helpers */

access(all) fun getExampleTokenBalance(_ acct: Test.TestAccount): UFix64 {
    let balance: UFix64? = (scriptExecutor("get_balance.cdc", [acct.address])! as! UFix64)
    return balance!
}

/* Test Helper */

access(all) fun getTestAccount(_ name: String): Test.TestAccount {
    if accounts[name] == nil {
        accounts[name] = blockchain.createAccount()
    }

    return accounts[name]!
}

/* Test Setup */

access(all) fun setup() {

    let sourceAccount = blockchain.createAccount()

    accounts["FungibleToken"] = sourceAccount
    accounts["NonFungibleToken"] = sourceAccount
    accounts["ViewResolver"] = sourceAccount
    accounts["MetadataViews"] = sourceAccount
    accounts["FungibleTokenMetadataViews"] = sourceAccount
    accounts["ExampleToken"] = sourceAccount

    blockchain.useConfiguration(
        Test.Configuration(
            addresses: {
                "FungibleToken": sourceAccount.address,
                "NonFungibleToken": sourceAccount.address,
                "ViewResolver": sourceAccount.address,
                "MetadataViews": sourceAccount.address,
                "FungibleTokenMetadataViews": sourceAccount.address,
                "ExampleToken": sourceAccount.address
            }
        )
    )

    deploy("ViewResolver", sourceAccount, "../contracts/utility/ViewResolver.cdc")
    deploy("FungibleToken", sourceAccount, "../contracts/FungibleToken-v2.cdc")
    deploy("NonFungibleToken", sourceAccount, "../contracts/utility/NonFungibleToken.cdc")
    deploy("MetadataViews", sourceAccount, "../contracts/utility/MetadataViews.cdc")
    deploy("FungibleTokenMetadataViews", sourceAccount, "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", sourceAccount, "../contracts/ExampleToken-v2.cdc")
}
