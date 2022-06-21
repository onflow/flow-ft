import path from "path";
import { emulator, init, getAccountAddress, deployContractByName, sendTransaction, shallPass, 
  executeScript } from "flow-js-testing";
  import fs from "fs";

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
    // We do some scafolding...

    // Getting the base path of the project
    const basePath = path.resolve(__dirname, "./../../../"); 
		// You can specify different port to parallelize execution of describe blocks
    const port = 8080; 
		// Setting logging flag to true will pipe emulator output to console
    const logging = false;

    await init(basePath, { port });
    await emulator.start(port, { logging });

    // ...then we deploy the ft and example token contracts using the getAccountAddress function
    // from the flow-js-testing library...

    // Create a service account and deploy contracts to it
    serviceAccount = await getAccountAddress("ServiceAccount");

    await deployContract({ to: serviceAccount,    name: "FungibleToken"});
    await deployContract({ to: serviceAccount,       name: "ExampleToken"});

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

// Defining the test suite for the fungible token switchboard
describe("fungibletokenswitchboard", ()=>{

  // Variables for holding the account address
  let serviceAccount;
  let fungibleTokenSwitchboardUser;

  // Before each test...
  beforeEach(async () => {
    // We do some scafolding...

    // Getting the base path of the project
    const basePath = path.resolve(__dirname, "./../../../"); 
		// You can specify different port to parallelize execution of describe blocks
    const port = 8080; 
		// Setting logging flag to true will pipe emulator output to console
    const logging = false;

    await init(basePath, { port });
    await emulator.start(port, { logging });

    // ...then we deploy the ft and example token contracts using the getAccountAddress function
    // from the flow-js-testing library...

    // Create a service account and deploy contracts to it
    serviceAccount = await getAccountAddress("ServiceAccount");

    await deployContract({ to: serviceAccount,    name: "FungibleToken"});
    await deployContract({ to: serviceAccount,    name: "ExampleToken"});
    await deployContract({ to: serviceAccount,    name: "FungibleTokenSwitchboard"});
    
    // Deployed at address which has the alias - SwitchboardUser
    fungibleTokenSwitchboardUser = await getAccountAddress("SwitchboardUser");
  
  });

  // After each test we stop the emulator, so it could be restarted
  afterEach(async () => {
    return emulator.stop();
  });

  // First test checks if switchboard can be installed
  test("should be able to setup switchboard", async () => {
    // Only step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
  });

  // Second test checks if switchboard user is able to add ft token vault capabilities
  test("should be able to add a vault capability", async () => {
    // First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: add example token vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
  });

  // Third test checks if switchboard user is able to remove ft token vault capabilities
  test("should be able to create and remove vault capability", async () => {
    // First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );   
    //Third step: add vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: remove vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/remove_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    // Fourth step: verify that no capabilities are returned
    const [result, e] = await executeScript({
      code: get_vault_types,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(result).toStrictEqual([]);
  });

  // Fourth test checks if switchboard user is able to add capabilities using
  // paths 
  test("should be able to add a capability using the token vault public path and delete it after", async () => {
    // First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );   
    //Third step: add vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/batch_add_vault_capabilities",
        args: [fungibleTokenSwitchboardUser],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    // Fourth step: verify that vault types are returned
    const [result, e] = await executeScript({
      code: get_vault_types,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(result).not.toBe(null);
    // Fifth step: remove vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/remove_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    // Sixth step: verify that no capabilities are returned
    const [result2, e2] = await executeScript({
      code: get_vault_types,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(result2).toStrictEqual([]);
  });

  // Fifth test checks if switchboard user is able to receive ft through the
  // switchboard deposit function
  test("should be able to receive tokens through switchboard", async () => {
    //First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: add vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Fourth step: mint tokens into service account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [serviceAccount, 100],
        signers: [serviceAccount]
      })
    );
    //Fifth step: transfer tokens to switchboard user
    await shallPass(
      sendTransaction({
        name: "switchboard/transfer_tokens",
        args: [fungibleTokenSwitchboardUser, 50],
        signers: [serviceAccount]
      })
    );
    //Sixth step: verify that switchboard user has 50 tokens
    const [result, e] = await executeScript({
      code: get_balance,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(parseFloat(result)).toBeCloseTo(50);
  });

  // Sixth test checks if switchboard user is able to receive ft through the
  // switchboard safeDeposit function
  test("should be able to receive tokens through safeDeposit", async () => {
    //First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: add vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Fourth step: mint tokens into service account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [serviceAccount, 100],
        signers: [serviceAccount]
      })
    );
    //Fifth step: transfer tokens to switchboard user
    await shallPass(
      sendTransaction({
        name: "switchboard/safe_transfer_tokens",
        args: [fungibleTokenSwitchboardUser, 50],
        signers: [serviceAccount]
      })
    );
    //Sixth step: verify that switchboard user has 50 tokens
    const [result, e] = await executeScript({
      code: get_balance,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(parseFloat(result)).toBeCloseTo(50);
  });

  // Seventh test checks if transaction is not rejected if switchboard the vault
  // is not added to the switchboard using the safeDeposit function
  test("should not reject transaction if vault is not added to switchboard", async () => {
    //First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: mint tokens into service account
    await shallPass(
      sendTransaction({
        name: "mint_tokens",
        args: [serviceAccount, 100],
        signers: [serviceAccount]
      })
    );
    //Save the balance after minting new tokens
    const [mintBalance, e] = await executeScript({
      code: get_balance,
      args: [serviceAccount]
    });
    //Third step: transfer tokens to switchboard user
    await shallPass(
      sendTransaction({
        name: "switchboard/safe_transfer_tokens",
        args: [fungibleTokenSwitchboardUser, 50],
        signers: [serviceAccount]
      })
    );
    //Fourth step: verify that the serviceAccount balance has not changed
    const [result, e2] = await executeScript({
      code: get_balance,
      args: [serviceAccount]
    });
    expect(parseFloat(result)).toBe(parseFloat(mintBalance));
  });

  // Eighth test checks if vault capabilities could be retrieved from a switchboard
  test("should be able to retrieve vault capabilities", async () => {
    //First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: setup example token vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: add vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    // Fourth step: verify that vault types are returned
    const [result, e] = await executeScript({
      code: get_vault_types,
      args: [fungibleTokenSwitchboardUser]
    });
    expect(result).not.toBe(null);
  });
  
});
