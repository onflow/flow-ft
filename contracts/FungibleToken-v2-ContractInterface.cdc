import FungibleToken from "./FungibleToken-v2.cdc"

pub contract interface FungibleTokenInterface {

    /// Contains the total supply of the fungible token
    pub var totalSupply: {Type: UFix64}

    /// Function to return the types that the contract implements
    pub fun getVaultTypes(): [FungibleToken.VaultInfo] {
        post {
            result.length > 0: "Must indicate what fungible token types this contract defines"
        }
    }

    /// createEmptyVault allows any user to create a new Vault that has a zero balance
    pub fun createEmptyVault(): @AnyResource{FungibleToken.Vault} {
          post {
               result.getBalance() == 0.0: "The newly created Vault must have zero balance"
       }
    }
}
