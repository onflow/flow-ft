import FungibleTokens from "./FungibleTokens.cdc" // 0xFungibleTokensADDRESS

pub contract ExampleTokens: FungibleTokens {

    /// Total supply of ExampleTokenSets in existence
    access(contract) var totalSupplyByID: {UInt64: UFix64}

    pub var CollectionStoragePath: StoragePath

    pub event ContractInitialized()

    /// TokensInitialized
    ///
    /// The event that is emitted when a new token is created
    pub event TokensInitialized(initialSupply: UFix64)

    /// TokensWithdrawn
    ///
    /// The event that is emitted when tokens are withdrawn from a Vault
    pub event TokensWithdrawn(tokenID: UInt64, amount: UFix64, from: Address?)

    /// TokensDeposited
    ///
    /// The event that is emitted when tokens are deposited to a Vault
    pub event TokensDeposited(tokenID: UInt64, amount: UFix64, to: Address?)

    /// TokensMinted
    ///
    /// The event that is emitted when new tokens are minted
    pub event TokensMinted(tokenID: UInt64, amount: UFix64)

    /// TokensBurned
    ///
    /// The event that is emitted when tokens are destroyed
    pub event TokensBurned(tokenID: UInt64, amount: UFix64)

    /// MinterCreated
    ///
    /// The event that is emitted when a new minter resource is created
    pub event MinterCreated(tokenID: UInt64, allowedAmount: UFix64)

    /// BurnerCreated
    ///
    /// The event that is emitted when a new burner resource is created
    pub event BurnerCreated()

    /// Vault
    ///
    /// Each user stores an instance of only the Vault in their storage
    /// The functions in the Vault and governed by the pre and post conditions
    /// in FungibleTokens when they are called.
    /// The checks happen at runtime whenever a function is called.
    ///
    /// Resources can only be created in the context of the contract that they
    /// are defined in, so there is no way for a malicious user to create Vaults
    /// out of thin air. A special Minter resource needs to be defined to mint
    /// new tokens.
    ///
    pub resource TokenVault: FungibleTokens.Provider, FungibleTokens.Receiver, FungibleTokens.Balance {

        /// The total balance of this vault
        pub var balance: UFix64
        pub let tokenID: UInt64

        // initialize the balance at resource creation time
        init(tokenID: UInt64, balance: UFix64) {
            self.balance = balance
            self.tokenID = tokenID
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
        pub fun withdraw(amount: UFix64): @FungibleTokens.TokenVault {
            self.balance = self.balance - amount
            emit TokensWithdrawn(tokenID: self.tokenID, amount: amount, from: self.owner?.address)
            return <-create TokenVault(tokenID: self.tokenID, balance: amount)
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
        pub fun deposit(from: @FungibleTokens.TokenVault) {
            let vault <- from as! @ExampleTokens.TokenVault
            self.balance = self.balance + vault.balance
            emit TokensDeposited(tokenID: self.tokenID, amount: vault.balance, to: self.owner?.address)
            vault.balance = 0.0
            destroy vault
        }

        destroy() {
            ExampleTokens.totalSupplyByID[self.tokenID] = ExampleTokens.totalSupplyByID[self.tokenID]! - self.balance
        }
    }

    /// createEmptyVault
    ///
    /// Function that creates a new Vault with a balance of zero
    /// and returns it to the calling context. A user must call this function
    /// and store the returned Vault in their storage in order to allow their
    /// account to be able to receive deposits of this token type.
    ///
    pub fun createEmptyTokenVault(tokenID: UInt64): @TokenVault {
        return <-create TokenVault(tokenID: tokenID, balance: 0.0)
    }

    pub resource Administrator {

        /// createNewMinter
        ///
        /// Function that creates and returns a new minter resource
        /// Minter can mint an allowance of tokens of the specified ID
        ///
        pub fun createNewMinter(tokenID: UInt64, allowedAmount: UFix64): @Minter {
            emit MinterCreated(tokenID: tokenID, allowedAmount: allowedAmount)
            return <-create Minter(tokenID: tokenID, allowedAmount: allowedAmount)
        }

        /// createNewBurner
        ///
        /// Function that creates and returns a new burner resource
        ///
        pub fun createNewBurner(): @Burner {
            emit BurnerCreated()
            return <-create Burner()
        }
    }

    /// Minter
    ///
    /// Resource object that token admin accounts can hold to mint new tokens.
    ///
    pub resource Minter {

        /// The amount of tokens that the minter is allowed to mint
        pub var allowedAmount: UFix64
        pub var tokenID: UInt64

        /// mintTokens
        ///
        /// Function that mints new tokens, adds them to the total supply,
        /// and returns them to the calling context.
        ///
        pub fun mintTokens(amount: UFix64): @ExampleTokens.TokenVault {
            pre {
                amount > 0.0: "Amount minted must be greater than zero"
                amount <= self.allowedAmount: "Amount minted must be less than the allowed amount"
            }
            ExampleTokens.totalSupplyByID[self.tokenID] = ExampleTokens.totalSupplyByID[self.tokenID]! + amount
            self.allowedAmount = self.allowedAmount - amount
            emit TokensMinted(tokenID: self.tokenID, amount: amount)
            return <-create TokenVault(tokenID: self.tokenID, balance: amount)
        }

        init(tokenID: UInt64, allowedAmount: UFix64) {
            self.allowedAmount = allowedAmount
            self.tokenID = tokenID
        }
    }

    /// Burner
    ///
    /// Resource object that token admin accounts can hold to burn tokens.
    ///
    pub resource Burner {

        /// burnTokens
        ///
        /// Function that destroys a Vault instance, effectively burning the tokens.
        ///
        /// Note: the burned tokens are automatically subtracted from the
        /// total supply in the Vault destructor.
        ///
        pub fun burnTokens(from: @FungibleTokens.TokenVault) {
            let vault <- from as! @ExampleTokens.TokenVault
            let amount = vault.balance
            let tokenID = vault.tokenID
            destroy vault
            emit TokensBurned(tokenID: tokenID, amount: amount)
        }
    }

    pub resource Collection: FungibleTokens.CollectionPublic, FungibleTokens.CollectionPrivate {
        pub var ownedTokenVaults: @{UInt64: FungibleTokens.TokenVault}

        pub fun deposit(token: @FungibleTokens.TokenVault) {
            if self.ownedTokenVaults[token.tokenID] == nil {
                self.ownedTokenVaults[token.tokenID] <-! token
            } else {
                self.ownedTokenVaults[token.tokenID]?.deposit!(from: <- token)
            }
        }
        
        pub fun getIDs(): [UInt64] {
            return self.ownedTokenVaults.keys
        }

        pub fun borrowTokenVault(id: UInt64): &FungibleTokens.TokenVault {
            let vaultRef = &self.ownedTokenVaults[id] as! &FungibleTokens.TokenVault
            return vaultRef
        }

        init() {
            self.ownedTokenVaults <- {}
        }

        destroy () {
            destroy self.ownedTokenVaults
        }

    }

    pub fun createEmptyCollection(): @FungibleTokens.Collection {
        return <- create Collection()
    }

    init() {
        self.totalSupplyByID = {0: 1000.0}
        self.CollectionStoragePath = /storage/ExampleTokens

        // Create the Vault with the total supply of tokens and save it in storage
        //
        let vault <- create TokenVault(tokenID: 0, balance: self.totalSupplyByID[0]!)
        
        // create collection to store the vaults
        let collection <- ExampleTokens.createEmptyCollection()
        
        // deposit vault into collection
        collection.deposit(token: <- vault)
        
        // save collection to storage
        self.account.save(<-collection, to: self.CollectionStoragePath)

        // create admin and save to storage
        let admin <- create Administrator()
        self.account.save(<-admin, to: /storage/ExampleTokenSetAdmin)

        // Emit an event that shows that the contract was initialized
        //
        emit TokensInitialized(initialSupply: self.totalSupplyByID[0]!)
    }
}
