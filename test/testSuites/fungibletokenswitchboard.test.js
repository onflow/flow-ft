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
describe("fungibletokenswitchboard", ()=>{

  // Variables for holding the account address
  let serviceAccount;
  let fungibleTokenSwitchboardUser;

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
    serviceAccount = await getAccountAddress("ServiceAccount");

    await deployContract({ to: serviceAccount,    name: "FungibleToken"});
    await deployContract({ to: serviceAccount,    name: "ExampleToken"});
    await deployContract({ to: serviceAccount,    name: "FungibleTokenSwitchboard"});
    
    // Deployed at address which has the alias - fungibletokenswitchboardUser
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
    
  // Second test checks if switchboard user is able to create and remove ft token
  // vault capabilities
  test("should be able to create and remove vault capability", async () => {
    // First step: setup switchboard
    await shallPass(
      sendTransaction({
        name: "switchboard/setup_account",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Second step: create vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    await shallPass(
      sendTransaction({
        name: "switchboard/remove_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
  });

  // Third test checks if switchboard user is able to receive ft through the
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
    //Second step: create vault capability
    await shallPass(
      sendTransaction({
        name: "switchboard/add_vault_capability",
        args: [],
        signers: [fungibleTokenSwitchboardUser]
      })
    );
    //Third step: setup ft vault
    await shallPass(
      sendTransaction({
        name: "setup_account",
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
        args: [50, fungibleTokenSwitchboardUser],
        signers: [serviceAccount]
      })
    );
  });
});