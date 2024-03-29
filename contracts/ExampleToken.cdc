import "FungibleToken"
import "MetadataViews"
import "FungibleTokenMetadataViews"

access(all) contract ExampleToken: FungibleToken {

    /// The event that is emitted when new tokens are minted
    access(all) event TokensMinted(amount: UFix64, type: String)

    /// Total supply of ExampleTokens in existence
    access(all) var totalSupply: UFix64

    /// Storage and Public Paths
    access(all) let VaultStoragePath: StoragePath
    access(all) let VaultPublicPath: PublicPath
    access(all) let ReceiverPublicPath: PublicPath
    access(all) let AdminStoragePath: StoragePath

    access(all) view fun getContractViews(resourceType: Type?): [Type] {
        return [
            Type<FungibleTokenMetadataViews.FTView>(),
            Type<FungibleTokenMetadataViews.FTDisplay>(),
            Type<FungibleTokenMetadataViews.FTVaultData>(),
            Type<FungibleTokenMetadataViews.TotalSupply>()
        ]
    }

    access(all) fun resolveContractView(resourceType: Type?, viewType: Type): AnyStruct? {
        switch viewType {
            case Type<FungibleTokenMetadataViews.FTView>():
                return FungibleTokenMetadataViews.FTView(
                    ftDisplay: self.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTDisplay>()) as! FungibleTokenMetadataViews.FTDisplay?,
                    ftVaultData: self.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
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
                    logos: medias,
                    socials: {
                        "twitter": MetadataViews.ExternalURL("https://twitter.com/flow_blockchain")
                    }
                )
            case Type<FungibleTokenMetadataViews.FTVaultData>():
                return FungibleTokenMetadataViews.FTVaultData(
                    storagePath: /storage/exampleTokenVault,
                    receiverPath: /public/exampleTokenReceiver,
                    metadataPath: /public/exampleTokenVault,
                    receiverLinkedType: Type<&ExampleToken.Vault>(),
                    metadataLinkedType: Type<&ExampleToken.Vault>(),
                    createEmptyVaultFunction: (fun(): @{FungibleToken.Vault} {
                        return <-ExampleToken.createEmptyVault(vaultType: Type<@ExampleToken.Vault>())
                    })
                )
            case Type<FungibleTokenMetadataViews.TotalSupply>():
                return FungibleTokenMetadataViews.TotalSupply(
                    totalSupply: ExampleToken.totalSupply
                )
        }
        return nil
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
    access(all) resource Vault: FungibleToken.Vault {

        /// The total balance of this vault
        access(all) var balance: UFix64

        // initialize the balance at resource creation time
        init(balance: UFix64) {
            self.balance = balance
        }

        /// Called when a fungible token is burned via the `Burner.burn()` method
        access(contract) fun burnCallback() {
            if self.balance > 0.0 {
                ExampleToken.totalSupply = ExampleToken.totalSupply - self.balance
            }
            self.balance = 0.0
        }

        access(all) view fun getViews(): [Type] {
            return ExampleToken.getContractViews(resourceType: nil)
        }

        access(all) fun resolveView(_ view: Type): AnyStruct? {
            return ExampleToken.resolveContractView(resourceType: nil, viewType: view)
        }

        /// getSupportedVaultTypes optionally returns a list of vault types that this receiver accepts
        access(all) view fun getSupportedVaultTypes(): {Type: Bool} {
            let supportedTypes: {Type: Bool} = {}
            supportedTypes[self.getType()] = true
            return supportedTypes
        }

        access(all) view fun isSupportedVaultType(type: Type): Bool {
            return self.getSupportedVaultTypes()[type] ?? false
        }

        /// Asks if the amount can be withdrawn from this vault
        access(all) view fun isAvailableToWithdraw(amount: UFix64): Bool {
            return amount <= self.balance
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
        access(FungibleToken.Withdraw) fun withdraw(amount: UFix64): @ExampleToken.Vault {
            self.balance = self.balance - amount
            return <-create Vault(balance: amount)
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
        access(all) fun deposit(from: @{FungibleToken.Vault}) {
            let vault <- from as! @ExampleToken.Vault
            self.balance = self.balance + vault.balance
            vault.balance = 0.0
            destroy vault
        }

        /// createEmptyVault
        ///
        /// Function that creates a new Vault with a balance of zero
        /// and returns it to the calling context. A user must call this function
        /// and store the returned Vault in their storage in order to allow their
        /// account to be able to receive deposits of this token type.
        ///
        access(all) fun createEmptyVault(): @ExampleToken.Vault {
            return <-create Vault(balance: 0.0)
        }
    }

    /// Minter
    ///
    /// Resource object that token admin accounts can hold to mint new tokens.
    ///
    access(all) resource Minter {
        /// mintTokens
        ///
        /// Function that mints new tokens, adds them to the total supply,
        /// and returns them to the calling context.
        ///
        access(all) fun mintTokens(amount: UFix64): @ExampleToken.Vault {
            ExampleToken.totalSupply = ExampleToken.totalSupply + amount
            emit TokensMinted(amount: amount, type: self.getType().identifier)
            return <-create Vault(balance: amount)
        }
    }

    /// createEmptyVault
    ///
    /// Function that creates a new Vault with a balance of zero
    /// and returns it to the calling context. A user must call this function
    /// and store the returned Vault in their storage in order to allow their
    /// account to be able to receive deposits of this token type.
    ///
    access(all) fun createEmptyVault(vaultType: Type): @ExampleToken.Vault {
        return <- create Vault(balance: 0.0)
    }

    init() {
        self.totalSupply = 1000.0

        self.VaultStoragePath = /storage/exampleTokenVault
        self.VaultPublicPath = /public/exampleTokenVault
        self.ReceiverPublicPath = /public/exampleTokenReceiver
        self.AdminStoragePath = /storage/exampleTokenAdmin 

        // Create the Vault with the total supply of tokens and save it in storage
        //
        let vault <- create Vault(balance: self.totalSupply)

        // Create a public capability to the stored Vault that exposes
        // the `deposit` method and getAcceptedTypes method through the `Receiver` interface
        // and the `balance` method through the `Balance` interface
        //
        let exampleTokenCap = self.account.capabilities.storage.issue<&ExampleToken.Vault>(self.VaultStoragePath)
        self.account.capabilities.publish(exampleTokenCap, at: self.VaultPublicPath)
        let receiverCap = self.account.capabilities.storage.issue<&ExampleToken.Vault>(self.VaultStoragePath)
        self.account.capabilities.publish(receiverCap, at: self.ReceiverPublicPath)

        self.account.storage.save(<-vault, to: /storage/exampleTokenVault)

        let admin <- create Minter()
        self.account.storage.save(<-admin, to: self.AdminStoragePath)
    }
}
