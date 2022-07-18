import FungibleToken from "./FungibleToken-v2.cdc"
import FungibleTokenInterface from "./FungibleToken-v2-ContractInterface.cdc"

pub contract ExampleToken: FungibleTokenInterface {

    /// Total supply of ExampleTokens in existence
    pub var totalSupply: UFix64

    /// Admin Path
    pub let AdminStoragePath: StoragePath

    /// EVENTS

    /// We would like to be able to define events in the resource

    /// TokensWithdrawn
    ///
    /// The event that is emitted when tokens are withdrawn from a Vault
    pub event TokensWithdrawn(amount: UFix64, from: Address?, type: Type)

    /// TokensDeposited
    ///
    /// The event that is emitted when tokens are deposited to a Vault
    pub event TokensDeposited(amount: UFix64, to: Address?, type: Type)

    /// TokensMinted
    ///
    /// The event that is emitted when new tokens are minted
    pub event TokensMinted(amount: UFix64, type: Type)

    /// TokensBurned
    ///
    /// The event that is emitted when tokens are destroyed
    pub event TokensBurned(amount: UFix64, type: Type)

    /// MinterCreated
    ///
    /// The event that is emitted when a new minter resource is created
    pub event MinterCreated(allowedAmount: UFix64, type: Type)

    /// Function to return the types that the contract implements
    pub fun getVaultTypes(): [FungibleToken.VaultInfo] {
        let typeArray: [FungibleToken.VaultInfo] = []

        let vault <- create Vault(balance: 0.0)

        let vaultInfo = vault.getTypeInfo()

        destroy vault

        typeArray.append(vaultInfo)
        return typeArray
    }

    /// Vault
    ///
    /// Each user stores an instance of only the Vault in their storage
    /// The functions in the Vault and governed by the pre and post conditions
    /// in FungibleToken when they are called.
    /// The checks happen at runtime whenever a function is called.
    ///
    /// Resources can only be created in the context of the contract that they
    /// are defined in, so there is no way for a malicious user to create Vaults
    /// out of thin air. A special Minter resource needs to be defined to mint
    /// new tokens.
    ///
    pub resource Vault: FungibleToken.Vault, FungibleToken.Provider, FungibleToken.Receiver, FungibleToken.Balance {

        /// Storage and Public Paths
        pub let VaultStoragePath: StoragePath
        pub let ReceiverPublicPath: PublicPath
        pub let BalancePublicPath: PublicPath

        /// The total balance of this vault
        pub var balance: UFix64

        // initialize the balance at resource creation time
        init(balance: UFix64) {
            self.balance = balance
            self.VaultStoragePath = /storage/exampleTokenVault
            self.ReceiverPublicPath = /public/exampleTokenReceiver
            self.BalancePublicPath = /public/exampleTokenBalance
        }
        
        /// Return information about the vault's type and paths
        pub fun getTypeInfo(): FungibleToken.VaultInfo {
            return FungibleToken.VaultInfo(type: self.getType(), VaultStoragePath: self.VaultStoragePath, ReceiverPublicPath: self.ReceiverPublicPath, BalancePublicPath: self.BalancePublicPath)
        }

        /// Get the balance of the vault
        pub fun getBalance(): UFix64 {
            return self.balance
        }

        /// withdraw
        ///
        /// Function that takes an amount as an argument
        /// and withdraws that amount from the Vault.
        ///
        /// It creates a new temporary Vault that is used to hold
        /// the money that is being transferred. It returns the newly
        /// created Vault to the context that called so it can be deposited
        /// elsewhere.
        ///
        pub fun withdraw(amount: UFix64): @ExampleToken.Vault{FungibleToken.Vault} {
            self.balance = self.balance - amount
            emit TokensWithdrawn(amount: amount, from: self.owner?.address, type: self.getType())
            return <-create Vault(balance: amount)
        }

        /// getAcceptedTypes optionally returns a list of vault types that this receiver accepts
        pub fun getAcceptedTypes(): [Type]? {
            let typeArray: [Type] = []
            typeArray.append(Type<@ExampleToken.Vault>())
            return typeArray
        }

        /// deposit
        ///
        /// Function that takes a Vault object as an argument and adds
        /// its balance to the balance of the owners Vault.
        ///
        /// It is allowed to destroy the sent Vault because the Vault
        /// was a temporary holder of the tokens. The Vault's balance has
        /// been consumed and therefore can be destroyed.
        ///
        pub fun deposit(from: @AnyResource{FungibleToken.Vault}) {
            let vault <- from as! @ExampleToken.Vault
            self.balance = self.balance + vault.balance
            emit TokensDeposited(amount: vault.balance, to: self.owner?.address, type: self.getType())
            vault.balance = 0.0
            destroy vault
        }

        destroy() {
            ExampleToken.totalSupply = ExampleToken.totalSupply - self.balance
        }
    }

    /// createEmptyVault
    ///
    /// Function that creates a new Vault with a balance of zero
    /// and returns it to the calling context. A user must call this function
    /// and store the returned Vault in their storage in order to allow their
    /// account to be able to receive deposits of this token type.
    ///
    pub fun createEmptyVault(): @Vault {
        return <-create Vault(balance: 0.0)
    }

    pub resource Administrator {

        /// createNewMinter
        ///
        /// Function that creates and returns a new minter resource
        ///
        pub fun createNewMinter(allowedAmount: UFix64): @Minter {
            emit MinterCreated(allowedAmount: allowedAmount, type: self.getType())
            return <-create Minter(allowedAmount: allowedAmount)
        }
    }

    /// Minter
    ///
    /// Resource object that token admin accounts can hold to mint new tokens.
    ///
    pub resource Minter {

        /// The amount of tokens that the minter is allowed to mint
        pub var allowedAmount: UFix64

        /// mintTokens
        ///
        /// Function that mints new tokens, adds them to the total supply,
        /// and returns them to the calling context.
        ///
        pub fun mintTokens(amount: UFix64): @ExampleToken.Vault {
            pre {
                amount > 0.0: "Amount minted must be greater than zero"
                amount <= self.allowedAmount: "Amount minted must be less than the allowed amount"
            }
            ExampleToken.totalSupply = ExampleToken.totalSupply + amount
            self.allowedAmount = self.allowedAmount - amount
            emit TokensMinted(amount: amount, type: self.getType())
            return <-create Vault(balance: amount)
        }

        init(allowedAmount: UFix64) {
            self.allowedAmount = allowedAmount
        }
    }

    init() {
        self.totalSupply = 1000.0

        self.AdminStoragePath = /storage/exampleTokenAdmin 

        // Create the Vault with the total supply of tokens and save it in storage
        //
        let vault <- create Vault(balance: self.totalSupply)

        let storagePath = vault.VaultStoragePath
        let receiverPath = vault.ReceiverPublicPath
        let balancePath = vault.BalancePublicPath

        self.account.save(<-vault, to: storagePath)

        // Create a public capability to the stored Vault that only exposes
        // the `deposit` method through the `Receiver` interface
        //
        self.account.link<&{FungibleToken.Receiver}>(
            receiverPath,
            target: storagePath
        )

        // Create a public capability to the stored Vault that only exposes
        // the `balance` field through the `Balance` interface
        //
        self.account.link<&ExampleToken.Vault{FungibleToken.Balance}>(
            balancePath,
            target: storagePath
        )

        let admin <- create Administrator()
        self.account.save(<-admin, to: self.AdminStoragePath)
    }
}
