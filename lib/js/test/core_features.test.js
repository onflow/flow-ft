import path from "path";
import { 
  emulator, 
  init, 
  getAccountAddress, 
  deployContractByName, 
  sendTransaction, 
  shallPass, 
  executeScript 
} from "@onflow/flow-js-testing";
  import fs from "fs";


// Auxiliary function for deploying the cadence contracts
async function deployContract(param) {
  const [result, error] = await deployContractByName(param);
  if (error != null) {
    console.log(`Error in deployment - ${error}`);
    emulator.stop();
    process.exit(1);
  }
}

const get_vault_types = fs.readFileSync(path.resolve(__dirname, "./../../../transactions/scripts/switchboard/get_vault_types.cdc"), {encoding:'utf8', flag:'r'});
const get_balance = fs.readFileSync(path.resolve(__dirname, "./../../../transactions/scripts/get_balance.cdc"), {encoding:'utf8', flag:'r'});


// Defining the test suite for the example token
describe("exampletoken", ()=>{

  // Variables for holding the account address
  let serviceAccount;
  let exampleTokenUserA;
  let exampleTokenUserB;

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

    await deployContract({ to: serviceAccount,    name: "FungibleToken"});
    await deployContract({ to: serviceAccount,    name: "utility/NonFungibleToken"});
    await deployContract({ to: serviceAccount,    name: "utility/MetadataViews"});
    await deployContract({ to: serviceAccount,    name: "FungibleTokenMetadataViews"});
    await deployContract({ to: serviceAccount,    name: "ExampleToken"});

    // ...and finally we get the address for a couple of regular accounts
    exampleTokenUserA  = await getAccountAddress("exampleTokenUserA");
    exampleTokenUserB = await getAccountAddress("exampleTokenUserB");
  
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