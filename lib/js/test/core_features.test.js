import path from "path";
import { 
  emulator, 
  init, 
  getAccountAddress, 
  deployContract, 
  deployContractByName, 
  sendTransaction, 
  shallPass, 
  executeScript,
  mintFlow
} from "@onflow/flow-js-testing";
  import fs from "fs";


async function deployContractByContractCode(param) {
  const [result, error] = await deployContract(param);
  if (error != null) {
    console.log(`Error in deployment - ${error}`);
    emulator.stop();
    process.exit(1);
  }
}

// Auxiliary function for deploying the cadence contracts
async function deployContractByContractName(param) {
  const [result, error] = await deployContractByName(param);
  if (error != null) {
    console.log(`Error in deployment - ${error}`);
    emulator.stop();
    process.exit(1);
  }
}

const get_vault_types = fs.readFileSync(path.resolve(__dirname, "./../../../transactions/scripts/switchboard/get_vault_types.cdc"), {encoding:'utf8', flag:'r'});
const get_balance = fs.readFileSync(path.resolve(__dirname, "./../../../transactions/scripts/get_balance.cdc"), {encoding:'utf8', flag:'r'});
const get_supported_vault_types = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/scripts/get_supported_vault_types.cdc"), {encoding:'utf8', flag:'r'});
const setup_account_tx = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/setup_account.cdc"), {encoding:'utf8', flag:'r'});
const setup_token_switchboard_tx = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/setup_token_switchboard.cdc"), {encoding:'utf8', flag:'r'});
const setup_token_forwarding_tx = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/tokenForwarder/setup_token_forwarder.cdc"), {encoding:'utf8', flag:'r'});
const setup_demo_token_tx = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/setup_account_demo.cdc"), {encoding:'utf8', flag:'r'});
const safe_generic_transfer_tx = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/safe_generic_transfer.cdc"), {encoding:'utf8', flag:'r'});
const get_balance_read = fs.readFileSync(path.resolve(__dirname, "./mocks/transactions/scripts/get_balance.cdc"), {encoding:'utf8', flag:'r'});
const is_recipient_valid = fs.readFileSync(path.resolve(__dirname, "./../../../transactions/scripts/tokenForwarder/is_recipient_valid.cdc"), {encoding:'utf8', flag:'r'});



const token_contract_code = fs.readFileSync(path.resolve(__dirname, "./mocks/contracts/Token.cdc"), {encoding:'utf8', flag:'r'});
const test_token_contract_code = fs.readFileSync(path.resolve(__dirname, "./mocks/contracts/TestToken.cdc"), {encoding:'utf8', flag:'r'});
const demo_token_contract_code = fs.readFileSync(path.resolve(__dirname, "./mocks/contracts/DemoToken.cdc"), {encoding:'utf8', flag:'r'});
const token_switchboard_contract_code = fs.readFileSync(path.resolve(__dirname, "./mocks/contracts/TokenSwitchboard.cdc"), {encoding:'utf8', flag:'r'});
const token_forwarding_contract_code = fs.readFileSync(path.resolve(__dirname, "./mocks/contracts/TokenForwarding.cdc"), {encoding:'utf8', flag:'r'});

// Defining the test suite for the example token
describe("CoreFeatures", ()=>{

  // Variables for holding the account address
  let serviceAccount;
  let exampleTokenUserA;
  let exampleTokenUserB;
  let testTokenUserA;
  let testTokenUserB;

  // Before each test...
  beforeEach(async () => {
    // We do some scaffolding...

    // Getting the base path of the project
    const basePath = path.resolve(__dirname, "./../../../"); 
		// Setting logging flag to true will pipe emulator output to console
    const logging = false;

    await init(basePath);
    await emulator.start({ logging });

    // ...then we deploy the ft and example token contracts using the getAccountAddress function
    // from the flow-js-testing library...

    // Create a service account and deploy contracts to it
    serviceAccount = await getAccountAddress("ServiceAccount");
    await mintFlow(serviceAccount, "100000000")

    await deployContractByContractName({ to: serviceAccount,    name: "FungibleToken"});
    await deployContractByContractName({ to: serviceAccount,    name: "utility/NonFungibleToken"});
    await deployContractByContractName({ to: serviceAccount,    name: "utility/MetadataViews"});
    await deployContractByContractName({ to: serviceAccount,    name: "FungibleTokenMetadataViews"});
    await deployContractByContractName({ to: serviceAccount,    name: "ExampleToken"});
    await deployContractByContractCode({ to: serviceAccount,    name: "Token", code: token_contract_code});
    await deployContractByContractCode({ to: serviceAccount,    name: "TestToken", code: test_token_contract_code});
    await deployContractByContractCode({ to: serviceAccount,    name: "DemoToken", code: demo_token_contract_code});
    await deployContractByContractCode({ to: serviceAccount,    name: "TokenSwitchboard", code: token_switchboard_contract_code});
    await deployContractByContractCode({ to: serviceAccount,    name: "TokenForwarding", code: token_forwarding_contract_code}); 

    // ...and finally we get the address for a couple of regular accounts
    exampleTokenUserA  = await getAccountAddress("exampleTokenUserA");
    exampleTokenUserB = await getAccountAddress("exampleTokenUserB");
    testTokenUserA = await getAccountAddress("testTokenUserA");
    testTokenUserB = await getAccountAddress("testTokenUserB");
  
  });

  // After each test we stop the emulator, so it could be restarted
  afterEach(async () => {
    return emulator.stop();
  });
  
  // First test is to check if the example token contract is deployed
  // just sending the tx defined in the transactions folder signed by a regular account
  test("should be able to setup account", async () => {
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserA]
      })
    );
  });

  test("should be able to return the correct supported vault types", async () => {
    await shallPass(
      sendTransaction({
        code: setup_account_tx,
        args: [],
        signers: [exampleTokenUserA]
      })
    );
  test("should able to setup token forwarding and check whether recipient is valid or not", async() => {
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserA]
      })
    );
    const [resultA, e] = await executeScript({
      code: get_supported_vault_types,
      args: [exampleTokenUserA, "/public/testTokenReceiver"]
    });
    expect(resultA[0].typeID).toStrictEqual('A.01cf0e2f2f715450.TestToken.Vault');
  });


  test("should return an empty type dictionary for custom receivers", async () => {
    await shallPass(
      sendTransaction({
        code: setup_token_switchboard_tx,
        args: [],
        signers: [exampleTokenUserA]
      })
    );
    const [resultA, e] = await executeScript({
      code: get_supported_vault_types,
      args: [exampleTokenUserA, "/public/tokenSwitchboardPublic"]
    });
    expect(resultA).toStrictEqual([]);
  });

  test("should successfully setup token forwarding and check whether it returns the correct supported vault types", async () => {
    await shallPass(
      sendTransaction({
        code: setup_account_tx,
        args: [],
        signers: [testTokenUserA]
      })
    );
    await shallPass(
      sendTransaction({
        code: setup_account_tx,
        args: [],
        signers: [testTokenUserB]
      })
    );
    await shallPass(
      sendTransaction({
        code: setup_token_forwarding_tx,
        args: [testTokenUserB],
        signers: [testTokenUserA]
      })
    );
    const [resultA, e] = await executeScript({
      code: get_supported_vault_types,
      args: [testTokenUserA, "/public/testTokenReceiver"]
    });
    expect(resultA[0].typeID).toStrictEqual("A.01cf0e2f2f715450.TestToken.Vault");
  });

  test("should successfully test the safe_generic_transfer function", async() => {
    await shallPass(
      sendTransaction({
        code: setup_account_tx,
        args: [],
        signers: [testTokenUserA]
      })
    );

    await shallPass(
      sendTransaction({
        code: setup_demo_token_tx,
        args: [],
        signers: [testTokenUserB]
      })
    );

    // Perform generic transfer to Account A
    await shallPass(
      sendTransaction({
        code: safe_generic_transfer_tx,
        args: [100, testTokenUserA, "/storage/testTokenVault", "/public/testTokenReceiver"],
        signers: [serviceAccount]
      })
    );
    
    const [balanceA, eA] = await executeScript({
      code: get_balance_read,
      args: [testTokenUserA]
    });

    expect(balanceA).toEqual("100.00000000");

    // Should fail to transfer and move funds back to the user
    await shallPass(
      sendTransaction({
        code: safe_generic_transfer_tx,
        args: [10, testTokenUserB, "/storage/testTokenVault", "/public/demoTokenReceiver"],
        signers: [testTokenUserA]
      })
    );

    const [balanceB, eB] = await executeScript({
      code: get_balance_read,
      args: [testTokenUserA]
    });

    expect(balanceB).toEqual("100.00000000");
  });

  // Second test mint tokens from the contract account to a regular account
  test("should be able to mint tokens", async () => {
    // Step 1: Setup a regular account
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserA]
      })
    );
    // Step 2: Mint tokens signing as example token admin, depositing to a regular account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [exampleTokenUserA, 100],
        signers: [serviceAccount]
      })
    );
  });

  // Third test transfer tokens between two regular accounts
  test("should be able to transfer tokens", async () => {
    // Step 1: Setup a regular account
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserA]
      })
    );
    // Step 2: Setup another regular account
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserB]
      })
    );
    // Step 3: Mint tokens signing as example token admin, depositing to the account A
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [exampleTokenUserA, 100],
        signers: [serviceAccount]
      })
    );
    // Step 4: Transfer 50 tokens from account A to account B
    await shallPass(
      sendTransaction({
        name: "transfer_tokens",
        args: [50, exampleTokenUserB],
        signers: [exampleTokenUserA]
      })
    );
    // Step 5: Check if account A has 50 tokens
    const [resultA, e] = await executeScript({
      code: get_balance,
      args: [exampleTokenUserA]
    });
    expect(parseFloat(resultA)).toBe(50);
    // Step 6: Check if account B has 50 tokens
    const [resultB, e2] = await executeScript({
      code: get_balance,
      args: [exampleTokenUserB]
    });
    expect(parseFloat(resultB)).toBe(50);
  })

  // Fourth test burns tokens
  test("should be able to burn tokens", async () => {
    // Step 1: Mint tokens to the very example token admin account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [serviceAccount, 100],
        signers: [serviceAccount]
      })
    );
    // Get the account balance to compare after burning the tokens
    const [accountBalance, e] = await executeScript({
      code: get_balance,
      args: [serviceAccount]
    });
    // Step 2: Burn 50 tokens from the example token admin account
    await shallPass(
      sendTransaction({
        name: "burn_tokens",
        args: [50],
        signers: [serviceAccount]
      })
    );
    // Step 3: Check if the example token admin account has burnt 50 tokens
    const [result, e2] = await executeScript({
      code: get_balance,
      args: [serviceAccount]
    });
    expect(parseFloat(accountBalance) - parseFloat(result)).toBe(50);
  })
});
