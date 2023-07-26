import FungibleToken from "FungibleToken-v2"
import MetadataViews from "MetadataViews"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MultipleVaults from "MultipleVaults.cdc"
import ViewResolver from "ViewResolver.cdc"

access(all) contract ExampleToken: ViewResolver, MultipleVaults {

    /// The event that is emitted when new tokens are minted
    access(all) event TokensMinted(amount: UFix64, type: String)

    /// Total supply of ExampleTokens in existence
    access(contract) var totalSupply: {Type: UFix64}

    /// Admin Path
    access(all) let AdminStoragePath: StoragePath

    /// Function to return the types that the contract implements
    access(all) view fun getVaultTypes(): [Type] {
        let typeArray: [Type] = [Type<@ExampleToken.Vault>()]
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
    access(all) resource Vault: FungibleToken.Vault {

        /// The total balance of this vault
        access(all) var balance: UFix64

        access(self) var storagePath: StoragePath
        access(self) var publicPath: PublicPath

        /// Returns the storage path where the vault should typically be stored
        access(all) view fun getDefaultStoragePath(): StoragePath? {
            return self.storagePath
        }

        /// Returns the public path where this vault should have a public capability
        access(all) view fun getDefaultPublicPath(): PublicPath? {
            return self.publicPath
        }

        access(all) view fun getViews(): [Type] {
            return [Type<FungibleTokenMetadataViews.FTView>(),
                    Type<FungibleTokenMetadataViews.FTDisplay>(),
                    Type<FungibleTokenMetadataViews.FTVaultData>()]
        }

        access(all) fun resolveView(_ view: Type): AnyStruct? {
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
                        logos: medias,
                        socials: {
                            "twitter": MetadataViews.ExternalURL("https://twitter.com/flow_blockchain")
                        }
                    )
                case Type<FungibleTokenMetadataViews.FTVaultData>():
                    let vaultRef = ExampleToken.account.borrow<&ExampleToken.Vault>(from: self.storagePath)
                        ?? panic("Could not borrow a reference to the stored vault")
                    return FungibleTokenMetadataViews.FTVaultData(
                        storagePath: self.storagePath,
                        receiverPath: self.publicPath,
                        metadataPath: self.publicPath,
                        providerPath: /private/exampleTokenVault,
                        receiverLinkedType: Type<&ExampleToken.Vault{FungibleToken.Receiver}>(),
                        metadataLinkedType: Type<&ExampleToken.Vault{FungibleToken.Balance, ViewResolver.Resolver}>(),
                        providerLinkedType: Type<&ExampleToken.Vault{FungibleToken.Provider}>(),
                        createEmptyVaultFunction: (fun(): @ExampleToken.Vault{FungibleToken.Vault} {
                            return <-vaultRef.createEmptyVault()
                        })
                    )
            }
            return nil
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

        // initialize the balance at resource creation time
        init(balance: UFix64) {
            self.balance = balance
            let identifier = "exampleTokenVault"
            self.storagePath = StoragePath(identifier: identifier)!
            self.publicPath = PublicPath(identifier: identifier)!
        }

        /// Get the balance of the vault
        access(all) view fun getBalance(): UFix64 {
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
        access(FungibleToken.Withdrawable) fun withdraw(amount: UFix64): @ExampleToken.Vault{FungibleToken.Vault} {
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

        access(all) fun transfer(amount: UFix64, receiver: Capability<&{FungibleToken.Receiver}>) {
            let transferVault <- self.withdraw(amount: amount)

            // Get a reference to the recipient's Receiver
            let receiverRef = receiver.borrow()
                ?? panic("Could not borrow receiver reference to the recipient's Vault")

            // Deposit the withdrawn tokens in the recipient's receiver
            receiverRef.deposit(from: <-transferVault)
        }

        /// createEmptyVault
        ///
        /// Function that creates a new Vault with a balance of zero
        /// and returns it to the calling context. A user must call this function
        /// and store the returned Vault in their storage in order to allow their
        /// account to be able to receive deposits of this token type.
        ///
        access(all) fun createEmptyVault(): @ExampleToken.Vault{FungibleToken.Vault} {
            return <-create Vault(balance: 0.0)
        }

        destroy() {
            if self.balance > 0.0 {
                ExampleToken.totalSupply[self.getType()] = ExampleToken.totalSupply[self.getType()]! - self.balance
            }
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
            ExampleToken.totalSupply[self.getType()] = ExampleToken.totalSupply[self.getType()]! + amount
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
    access(all) fun createEmptyVault(vaultType: Type): @{FungibleToken.Vault} {
        switch vaultType {
            case Type<@ExampleToken.Vault>():
                return <- create Vault(balance: 0.0)
            default:
                return <- create Vault(balance: 0.0)
        }
    }

    init() {
        self.totalSupply = {}
        self.totalSupply[Type<@ExampleToken.Vault>()] = 1000.0

        self.AdminStoragePath = /storage/exampleTokenAdmin 

        // Create the Vault with the total supply of tokens and save it in storage
        //
        let vault <- create Vault(balance: self.totalSupply[Type<@ExampleToken.Vault>()]!)
        self.account.save(<-vault, to: /storage/exampleTokenVault)

        // Create a public capability to the stored Vault that exposes
        // the `deposit` method and getAcceptedTypes method through the `Receiver` interface
        // and the `getBalance()` method through the `Balance` interface
        //
        self.account.link<&{FungibleToken.Receiver, FungibleToken.Balance}>(
            /public/exampleTokenVault,
            target: /storage/exampleTokenVault
        )

        let admin <- create Minter()
        self.account.save(<-admin, to: self.AdminStoragePath)
    }
}
 