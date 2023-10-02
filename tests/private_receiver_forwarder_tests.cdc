import Test
import "test_helpers.cdc"

access(all) let sourceAccount = blockchain.createAccount()
access(all) let accounts: {String: Test.Account} = {}

access(all) let exampleToken = "ExampleToken"

access(all) let senderStoragePath = /storage/Sender
access(all) let privateReceiverStoragePath = /storage/PrivateReceiver
access(all) let privateReceiverPublicPath = /public/PrivateReceiver

/* Test Cases */

access(all) fun testSetupForwader() {
    let alice = blockchain.createAccount()
    txExecutor("privateForwarder/setup_and_create_forwarder.cdc", [sourceAccount], [], nil, nil)
}

access(all) fun testTransferPrivateTokens() {
    let sender = getTestAccount(exampleToken)
    let senderBalanceBefore = getExampleTokenBalance(sender)
    assert(senderBalanceBefore == 1000.0, message: "ExampleToken balance should be 1000.0")

    let recipient = blockchain.createAccount()
    let recipientAmount = 300.0

    let pair = {recipient.address: recipientAmount}

    txExecutor("privateForwarder/setup_and_create_forwarder.cdc", [recipient], [], nil, nil)
    txExecutor("privateForwarder/transfer_private_many_accounts.cdc", [sender], [pair], nil, nil)

    let recipientBalance = getExampleTokenBalance(recipient)
    Test.assertEqual(recipientAmount, recipientBalance)

    let senderBalanceAfter = getExampleTokenBalance(sender)
    Test.assertEqual(senderBalanceBefore - recipientAmount, senderBalanceAfter)
}


/* Transaction Helpers */

access(all) fun setupExampleToken(_ acct: Test.Account) {
    txExecutor("setup_account.cdc", [acct], [], nil, nil)
}

access(all) fun mintExampleToken(_ acct: Test.Account, recipient: Address, amount: UFix64) {
    txExecutor("mint_tokens.cdc", [acct], [recipient, amount], nil, nil)
}

access(all) fun setupTokenForwarder(_ acct: Test.Account) {
    txExecutor("privateForwarder/setup_and_create_forwarder.cdc", [acct], [], nil, nil)
}

/* Script Helpers */

access(all) fun getExampleTokenBalance(_ acct: Test.Account): UFix64 {
    let balance: UFix64? = (scriptExecutor("get_balance.cdc", [acct.address])! as! UFix64)
    return balance!
}

/* Test Helper */

access(all) fun getTestAccount(_ name: String): Test.Account {
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
    accounts["MetadataViews"] = sourceAccount
    accounts["FungibleTokenMetadataViews"] = sourceAccount
    accounts["ExampleToken"] = sourceAccount
    accounts["PrivateReceiverForwarder"] = sourceAccount

    blockchain.useConfiguration(Test.Configuration({
        "FungibleToken": sourceAccount.address,
        "NonFungibleToken": sourceAccount.address,
        "MetadataViews": sourceAccount.address,
        "FungibleTokenMetadataViews": sourceAccount.address,
        "ExampleToken": sourceAccount.address,
        "PrivateReceiverForwarder": sourceAccount.address
    }))

    // helper nft contract so we can actually talk to nfts with tests
    deploy("FungibleToken", sourceAccount, "../contracts/FungibleToken.cdc")
    deploy("NonFungibleToken", sourceAccount, "../contracts/utility/NonFungibleToken.cdc")
    deploy("MetadataViews", sourceAccount, "../contracts/utility/MetadataViews.cdc")
    deploy("FungibleTokenMetadataViews", sourceAccount, "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", sourceAccount, "../contracts/ExampleToken.cdc")
    deployWithArgs(
        "PrivateReceiverForwarder",
        sourceAccount,
        "../contracts/utility/PrivateReceiverForwarder.cdc",
        args: [
            senderStoragePath,
            privateReceiverStoragePath,
            privateReceiverPublicPath
        ]
    )
}
