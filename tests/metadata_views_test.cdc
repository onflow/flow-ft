import Test
import BlockchainHelpers
import "test_helpers.cdc"
import "ViewResolver"
import "FungibleTokenMetadataViews"
import "ExampleToken"
import "FungibleToken"

access(all) let admin = Test.getAccount(0x0000000000000007)

/* Test Setup */

access(all) fun setup() {
    deploy("ViewResolver", "../contracts/utility/ViewResolver.cdc")
    deploy("Burner", "../contracts/utility/Burner.cdc")
    deploy("FungibleToken", "../contracts/FungibleToken.cdc")
    deploy("NonFungibleToken", "../contracts/utility/NonFungibleToken.cdc")
    deploy("MetadataViews", "../contracts/utility/MetadataViews.cdc")
    deploy("FungibleTokenMetadataViews", "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", "../contracts/ExampleToken.cdc")
}

/* Test Cases */

access(all) fun testSetupAccountUsingFTView() {
    let alice = Test.createAccount()
    let bob = Test.createAccount()

    setupExampleToken(alice)
    let aliceBalance = getExampleTokenBalance(alice)
    txExecutor("metadata/setup_account_from_vault_reference.cdc", [bob], [alice.address, /public/exampleTokenVault], nil, nil)
    let bobBalance = getExampleTokenBalance(bob)

    Test.assertEqual(0.0, aliceBalance)
    Test.assertEqual(0.0, bobBalance)
}

access(all) fun testSetupAccountUsingContractAddressAndName() {
    let bob = Test.createAccount()

    txExecutor("metadata/setup_account_from_address.cdc", [bob], [admin.address, "ExampleToken"], nil, nil)
    let bobBalance = getExampleTokenBalance(bob)

    Test.assertEqual(0.0, bobBalance)
}

access(all) fun testRetrieveVaultDisplayInfo() {
    let alice = Test.createAccount()

    setupExampleToken(alice)
    let scriptResult = executeScript(
        "scripts/example_token_vault_display_strict_equal.cdc",
        [alice.address]
    )
    Test.expect(scriptResult, Test.beSucceeded())
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