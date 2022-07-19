import FungibleToken from "./FungibleToken-v2.cdc"

pub contract interface FungibleTokenInterface {

    /// TokensWithdrawn
    ///
    /// The event that is emitted when tokens are withdrawn from a Vault
    pub event TokensWithdrawn(amount: UFix64, from: Address?, type: Type)

    /// TokensDeposited
    ///
    /// The event that is emitted when tokens are deposited to a Vault
    pub event TokensDeposited(amount: UFix64, to: Address?, type: Type)

    /// TokensTransferred
    ///
    /// The event that is emitted when tokens are transferred from one account to another
    pub event TokensTransferred(amount: UFix64, from: Address?, to: Address?, type: Type)

    /// TokensMinted
    ///
    /// The event that is emitted when new tokens are minted
    pub event TokensMinted(amount: UFix64, type: Type)

    /// Contains the total supply of the fungible token
    pub var totalSupply: {Type: UFix64}

    /// Function to return the types that the contract implements
    pub fun getVaultTypes(): {Type: FungibleToken.VaultInfo} {
        post {
            result.length > 0: "Must indicate what fungible token types this contract defines"
        }
    }
}