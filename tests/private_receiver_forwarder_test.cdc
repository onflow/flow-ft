import Test
import BlockchainHelpers
import "test_helpers.cdc"
import "ExampleToken"

/* Test Setup */

access(all) let admin = Test.getAccount(0x0000000000000007)
access(all) let senderStoragePath = /storage/Sender
access(all) let privateReceiverStoragePath = /storage/PrivateReceiver
access(all) let privateReceiverPublicPath = /public/PrivateReceiver

access(all) fun setup() {

    // helper nft contract so we can actually talk to nfts with tests
    deploy("ViewResolver", "../contracts/utility/ViewResolver.cdc")
    deploy("Burner", "../contracts/utility/Burner.cdc")
    deploy("FungibleToken", "../contracts/FungibleToken.cdc")
    deploy("NonFungibleToken", "../contracts/utility/NonFungibleToken.cdc")
    deploy("MetadataViews", "../contracts/utility/MetadataViews.cdc")
    deploy("FungibleTokenMetadataViews", "../contracts/FungibleTokenMetadataViews.cdc")
    deploy("ExampleToken", "../contracts/ExampleToken.cdc")
    deployWithArgs(
        "PrivateReceiverForwarder",
        "../contracts/utility/PrivateReceiverForwarder.cdc",
        args: [
            senderStoragePath,
            privateReceiverStoragePath,
            privateReceiverPublicPath
        ]
    )
}

/* Test Cases */

access(all) fun testSetupForwader() {
    let alice = Test.createAccount()
    let txResult = executeTransaction(
        "../transactions/privateForwarder/setup_and_create_forwarder.cdc",
        [],
        alice
    )
    Test.expect(txResult, Test.beSucceeded())
}

access(all) fun testTransferPrivateTokens() {
    let senderBalanceBefore = getExampleTokenBalance(admin)
    assert(senderBalanceBefore == 1000.0, message: "ExampleToken balance should be 1000.0")

    let recipient = Test.createAccount()
    let recipientAmount = 300.0

    let pair = {recipient.address: recipientAmount}

    txExecutor("../transactions/privateForwarder/setup_and_create_forwarder.cdc", [recipient], [], nil, nil)
    
    txExecutor("../transactions/privateForwarder/transfer_private_many_accounts.cdc", [admin], [pair], nil, nil)

    let recipientBalance = getExampleTokenBalance(recipient)
    Test.assertEqual(recipientAmount, recipientBalance)

    let senderBalanceAfter = getExampleTokenBalance(admin)
    Test.assertEqual(senderBalanceBefore - recipientAmount, senderBalanceAfter)
}


/* Transaction Helpers */

access(all) fun setupExampleToken(_ acct: Test.TestAccount) {
    txExecutor("setup_account.cdc", [acct], [], nil, nil)
}

access(all) fun mintExampleToken(_ acct: Test.TestAccount, recipient: Address, amount: UFix64) {
    txExecutor("mint_tokens.cdc", [acct], [recipient, amount], nil, nil)
}

access(all) fun setupTokenForwarder(_ acct: Test.TestAccount) {
    txExecutor("../transactions/privateForwarder/setup_and_create_forwarder.cdc", [acct], [], nil, nil)
}

/* Script Helpers */

access(all) fun getExampleTokenBalance(_ acct: Test.TestAccount): UFix64 {
    let balance: UFix64? = (scriptExecutor("get_balance.cdc", [acct.address])! as! UFix64)
    return balance!
}
