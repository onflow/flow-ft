import FungibleToken from "./FungibleToken-v2.cdc"
import FungibleTokenMetadataViews from "./FungibleTokenMetadataViews.cdc"

pub contract ExampleToken: FungibleToken {

    /// Total supply of ExampleTokens in existence
    pub var totalSupply: {Type: UFix64}

    /// Admin Path
    pub let AdminStoragePath: StoragePath

    /// EVENTS
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

    /// TokensBurned
    ///
    /// The event that is emitted when tokens are destroyed
    pub event TokensBurned(amount: UFix64, type: Type)

    /// Function to return the types that the contract implements
    pub fun getVaultTypes(): {Type: FungibleTokenMetadataViews.FTView} {
        let typeDictionary: {Type: FungibleTokenMetadataViews.FTView} = {}

        let vault <- create Vault(balance: 0.0)

        typeDictionary[vault.getType()] = vault.resolveView(Type<FungibleTokenMetadataViews.FTView>())

        destroy vault
        
        return typeDictionary
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
    pub resource Vault: FungibleToken.Vault, FungibleToken.Provider, FungibleToken.Transferable, FungibleToken.Receiver, FungibleToken.Balance, MetadataViews.Resolver {

        /// The total balance of this vault
        pub var balance: UFix64

        access(self) var storagePath: StoragePath
        access(self) var publicPath: PublicPath

        /// Returns the standard storage path for the Vault
        pub fun getDefaultStoragePath(): StoragePath? {
            return self.storagePath
        }

        /// Returns the standard public path for the Vault
        pub fun getPublicReceiverBalancePath(): PublicPath? {
            return self.publicPath
        }

        // initialize the balance at resource creation time
        init(balance: UFix64) {
            self.balance = balance
            let identifier = "exampleTokenVault"
            self.storagePath = StoragePath(identifier: identifier)
            self.publicPath = PublicPath(identifier: identifier)
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
        /// the tokens that are being transferred. It returns the newly
        /// created Vault to the context that called so it can be deposited
        /// elsewhere.
        ///
        pub fun withdraw(amount: UFix64): @ExampleToken.Vault{FungibleToken.Vault} {
            self.balance = self.balance - amount
            emit TokensWithdrawn(amount: amount, from: self.owner?.address, type: self.getType())
            return <-create Vault(balance: amount)
        }

        /// getAcceptedTypes optionally returns a list of vault types that this receiver accepts
        pub fun getAcceptedTypes(): {Type: Bool} {
            let types: {Type: Bool} = {}
            types[Type<@ExampleToken.Vault>()] = true
            return types
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

        pub fun transfer(amount: UFix64, recipient: Capability<&{FungibleToken.Receiver}>) {
            let transferVault <- self.withdraw(amount: amount)

            // Get a reference to the recipient's Receiver
            let receiverRef = recipient.borrow()
                ?? panic("Could not borrow receiver reference to the recipient's Vault")

            // Deposit the withdrawn tokens in the recipient's receiver
            receiverRef.deposit(from: <-transferVault)

            emit TokensTransferred(amount: amount, from: self.owner?.address, to: receiverRef.owner?.address, type: self.getType())
        }

        /// createEmptyVault
        ///
        /// Function that creates a new Vault with a balance of zero
        /// and returns it to the calling context. A user must call this function
        /// and store the returned Vault in their storage in order to allow their
        /// account to be able to receive deposits of this token type.
        ///
        pub fun createEmptyVault(): @ExampleToken.Vault{FungibleToken.Vault} {
            return <-create Vault(balance: 0.0)
        }

        destroy() {
            if self.balance > 0.0 {
                ExampleToken.totalSupply[self.getType()] = ExampleToken.totalSupply[self.getType()]! - self.balance
            }
        }

        /// The way of getting all the Metadata Views implemented by ExampleToken
        ///
        /// @return An array of Types defining the implemented views. This value will be used by
        ///         developers to know which parameter to pass to the resolveView() method.
        ///
        pub fun getViews(): [Type]{
            return [Type<FungibleTokenMetadataViews.FTView>(),
                    Type<FungibleTokenMetadataViews.FTDisplay>(),
                    Type<FungibleTokenMetadataViews.FTVaultData>()]
        }

        /// The way of getting a Metadata View out of the ExampleToken
        ///
        /// @param view: The Type of the desired view.
        /// @return A structure representing the requested view.
        ///
        pub fun resolveView(_ view: Type): AnyStruct? {
            switch view {
                case Type<FungibleTokenMetadataViews.FTView>():
                    return FungibleTokenMetadataViews.FTView(
                        ftDisplay: self.resolveView(Type<FungibleTokenMetadataViews.FTDisplay>()) as! FungibleTokenMetadataViews.FTDisplay?,
                        ftVaultData: self.resolveView(Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
                    )
                case Type<FungibleTokenMetadataViews.FTDisplay>():
                    let media = MetadataViews.Media(
                            file: MetadataViews.HTTPFile(
                            url: "https://assets.website-files.com/5f6294c0c7a8cdd643b1c820/5f6294c0c7a8cda55cb1c936_Flow_Wordmark.svg"
                        ),
                        mediaType: "image/svg+xml"
                    )
                    let medias = MetadataViews.Medias([media])
                    return FungibleTokenMetadataViews.FTDisplay(
                        name: "Example Fungible Token",
                        symbol: "EFT",
                        description: "This fungible token is used as an example to help you develop your next FT #onFlow.",
                        externalURL: MetadataViews.ExternalURL("https://example-ft.onflow.org"),
                        logo: medias,
                        socials: {
                            "twitter": MetadataViews.ExternalURL("https://twitter.com/flow_blockchain")
                        }
                    )
                case Type<FungibleTokenMetadataViews.FTVaultData>():
                    return FungibleTokenMetadataViews.FTVaultData(
                        storagePath: self.getStoragePath(),
                        receiverPath: self.getPublicReceiverBalancePath(),
                        metadataPath: self.getPublicReceiverBalancePath(),
                        providerPath: /private/exampleTokenVault,
                        receiverLinkedType: Type<&{FungibleToken.Receiver}>(),
                        metadataLinkedType: Type<&{FungibleToken.Balance, MetadataViews.Resolver}>(),
                        providerLinkedType: Type<&ExampleToken.Vault{FungibleToken.Provider, MetadataViews.Resolver}>(),
                        createEmptyVaultFunction: (fun (): @ExampleToken.Vault{FungibleToken.Vault} {
                            return <-self.createEmptyVault()
                        })
                    )
            }
            return nil
        }
    }

    pub resource Administrator {

        /// createNewMinter
        ///
        /// Function that creates and returns a new minter resource
        ///
        pub fun createNewMinter(allowedAmount: UFix64): @Minter {
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
            ExampleToken.totalSupply[self.getType()] = ExampleToken.totalSupply[self.getType()]! + amount
            self.allowedAmount = self.allowedAmount - amount
            emit TokensMinted(amount: amount, type: self.getType())
            return <-create Vault(balance: amount)
        }

        init(allowedAmount: UFix64) {
            self.allowedAmount = allowedAmount
        }
    }

    /// createEmptyVault
    ///
    /// Function that creates a new Vault with a balance of zero
    /// and returns it to the calling context. A user must call this function
    /// and store the returned Vault in their storage in order to allow their
    /// account to be able to receive deposits of this token type.
    ///
    pub fun createEmptyVault(vaultType: Type): @{FungibleToken.Vault}? {
        switch vaultType {
            case Type<@ExampleToken.Vault>():
                return <- create Vault(balance: 0.0)
            default:
                return nil
        }
    }

    init() {
        self.totalSupply = {}
        self.totalSupply[Type<@ExampleToken.Vault>()] = 1000.0

        self.AdminStoragePath = /storage/exampleTokenAdmin 

        // Create the Vault with the total supply of tokens and save it in storage
        //
        let vault <- create Vault(balance: self.totalSupply[Type<@ExampleToken.Vault>()]!)

        let storagePath = vault.StoragePath
        let receiverBalancePath = vault.PublicReceiverBalancePath

        self.account.save(<-vault, to: storagePath)

        // Create a public capability to the stored Vault that exposes
        // the `deposit` method and getAcceptedTypes method through the `Receiver` interface
        // and the `getBalance()` method through the `Balance` interface
        //
        self.account.link<&{FungibleToken.Receiver, FungibleToken.Balance}>(
            receiverBalancePath,
            target: storagePath
        )

        let admin <- create Administrator()
        self.account.save(<-admin, to: self.AdminStoragePath)
    }
}
