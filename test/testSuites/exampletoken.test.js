import path from "path";
import { emulator, init, getAccountAddress, deployContractByName, sendTransaction, shallPass, 
  executeScript, shallRevert, getFlowBalance, mintFlow } from "flow-js-testing";

// Increase timeout if your tests failing due to timeout
jest.setTimeout(10000);

// Auxiliary function for deploying the cadence contracts
async function deployContract(param) {
  const [result, error] = await deployContractByName(param);
  if (error != null) {
    console.log(`Error in deployment - ${error}`);
    emulator.stop();
    process.exit(1);
  }
}


// Defining the test suite for the example token
describe("exampletoken", ()=>{

  // Variables for holding the account address
  let serviceAccountAddress;
  let exampleTokenUserA;
  let exampleTokenUserB;

  // Before each test...
  beforeEach(async () => {
    // We do some scafolding...

    // Getting the base path of the project
    const basePath = path.resolve(__dirname, "../../"); 
		// You can specify different port to parallelize execution of describe blocks
    const port = 8080; 
		// Setting logging flag to true will pipe emulator output to console
    const logging = false;

    await init(basePath, { port });
    await emulator.start(port, {logging});

    // ...then we deploy the ft and example token contracts using the getAccountAddress function
    // from the flow-js-testing library...

    // Create a service account and deploy contracts to it
    serviceAccountAddress = await getAccountAddress("ServiceAccount");

    await deployContract({ to: serviceAccountAddress,    name: "FungibleToken"});
    await deployContract({ to: serviceAccountAddress,       name: "ExampleToken"});

    // ...and finally we get the address for a copuple of regular accounts
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
  })

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
        signers: [serviceAccountAddress]
      })
    );

  })

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
        signers: [serviceAccountAddress]
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
  })

  // Fourth test burns tokens
  test("should be able to burn tokens", async () => {
    // Step 1: Mint tokens to the very example token admin account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [serviceAccountAddress, 100],
        signers: [serviceAccountAddress]
      })
    );
    // Step 2: Burn 50 tokens from the example token admin account
    await shallPass(
      sendTransaction({
        name: "burn_tokens",
        args: [50],
        signers: [serviceAccountAddress]
      })
    );
  })

})
