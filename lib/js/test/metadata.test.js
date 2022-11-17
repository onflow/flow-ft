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
  
  // First test uses a vault as a resolver to get its FTView and use it to setup an account for using
  // the ExampleToken without importing the actual ExampleToken contract
  test("should be able to setup an account retrieving the FTView from other account", async () => {
    await shallPass(
      sendTransaction({
        name: "setup_account",
        args: [],
        signers: [exampleTokenUserA]
      })
    );

    await shallPass(
      sendTransaction({
        name: "./metadata/setup_account_from_vault_reference",
        args: [exampleTokenUserA, "/public/exampleTokenMetadata"],
        signers: [exampleTokenUserB]
      })
    );
  })
  
});