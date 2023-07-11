import FungibleToken from "FungibleToken"
import MetadataViews from "MetadataViews"

/// This contract implements the metadata standard proposed
/// in FLIP-1087.
/// 
/// Ref: https://github.com/onflow/flips/blob/main/application/20220811-fungible-tokens-metadata.md
/// 
/// Structs and resources can implement one or more
/// metadata types, called views. Each view type represents
/// a different kind of metadata.
///
pub contract FungibleTokenMetadataViews {

    /// FTView wraps FTDisplay and FTVaultData, and is used to give a complete 
    /// picture of a Fungible Token. Most Fungible Token contracts should 
    /// implement this view.
    ///
    pub struct FTView {
        pub let ftDisplay: FTDisplay?     
        pub let ftVaultData: FTVaultData?
        init(
            ftDisplay: FTDisplay?,
            ftVaultData: FTVaultData?
        ) {
            self.ftDisplay = ftDisplay
            self.ftVaultData = ftVaultData
        }
    }

    /// Helper to get a FT view.
    ///
    /// @param viewResolver: A reference to the resolver resource
    /// @return A FTView struct
    ///
    pub fun getFTView(viewResolver: &{MetadataViews.Resolver}): FTView {
        let maybeFTView = viewResolver.resolveView(Type<FTView>())
        if let ftView = maybeFTView {
            return ftView as! FTView
        }
        return FTView(
            ftDisplay: self.getFTDisplay(viewResolver),
            ftVaultData: self.getFTVaultData(viewResolver)
        )
    }

    /// View to expose the information needed to showcase this FT. 
    /// This can be used by applications to give an overview and 
    /// graphics of the FT.
    ///
    pub struct FTDisplay {
        /// The display name for this token.
        ///
        /// Example: "Flow"
        ///
        pub let name: String

        /// The abbreviated symbol for this token.
        ///
        /// Example: "FLOW"
        pub let symbol: String

        /// A description the provides an overview of this token.
        ///
        /// Example: "The FLOW token is the native currency of the Flow network."
        pub let description: String

        /// External link to a URL to view more information about the fungible token.
        pub let externalURL: MetadataViews.ExternalURL

        /// One or more versions of the fungible token logo.
        pub let logos: MetadataViews.Medias

        /// Social links to reach the fungible token's social homepages.
        /// Possible keys may be "instagram", "twitter", "discord", etc.
        pub let socials: {String: MetadataViews.ExternalURL}

        init(
            name: String,
            symbol: String,
            description: String,
            externalURL: MetadataViews.ExternalURL,
            logos: MetadataViews.Medias,
            socials: {String: MetadataViews.ExternalURL}
        ) {
            self.name = name
            self.symbol = symbol
            self.description = description
            self.externalURL = externalURL
            self.logos = logos
            self.socials = socials
        }
    }

    /// Helper to get FTDisplay in a way that will return a typed optional.
    /// 
    /// @param viewResolver: A reference to the resolver resource
    /// @return An optional FTDisplay struct
    ///
    pub fun getFTDisplay(_ viewResolver: &{MetadataViews.Resolver}): FTDisplay? {
        if let maybeDisplayView = viewResolver.resolveView(Type<FTDisplay>()) {
            if let displayView = maybeDisplayView as? FTDisplay {
                return displayView
            }
        }
        return nil
    }

    /// View to expose the information needed store and interact with a FT vault.
    /// This can be used by applications to setup a FT vault with proper 
    /// storage and public capabilities.
    ///
    pub struct FTVaultData {
        /// Path in storage where this FT vault is recommended to be stored.
        pub let storagePath: StoragePath

        /// Public path which must be linked to expose the public receiver capability.
        pub let receiverPath: PublicPath

        /// Public path which must be linked to expose the balance and resolver public capabilities.
        pub let metadataPath: PublicPath

        /// Private path which should be linked to expose the provider capability to withdraw funds 
        /// from the vault.
        pub let providerPath: PrivatePath

        /// Type that should be linked at the `receiverPath`. This is a restricted type requiring 
        /// the `FungibleToken.Receiver` interface.
        pub let receiverLinkedType: Type

        /// Type that should be linked at the `receiverPath`. This is a restricted type requiring 
        /// the `FungibleToken.Balance` and `MetadataViews.Resolver` interfaces.
        pub let metadataLinkedType: Type

        /// Type that should be linked at the aforementioned private path. This 
        /// is normally a restricted type with at a minimum the `FungibleToken.Provider` interface.
        pub let providerLinkedType: Type

        /// Function that allows creation of an empty FT vault that is intended
        /// to store the funds.
        pub let createEmptyVault: ((): @FungibleToken.Vault)

        init(
            storagePath: StoragePath,
            receiverPath: PublicPath,
            metadataPath: PublicPath,
            providerPath: PrivatePath,
            receiverLinkedType: Type,
            metadataLinkedType: Type,
            providerLinkedType: Type,
            createEmptyVaultFunction: ((): @FungibleToken.Vault)
        ) {
            pre {
                receiverLinkedType.isSubtype(of: Type<&{FungibleToken.Receiver}>()): "Receiver public type must include FungibleToken.Receiver."
                metadataLinkedType.isSubtype(of: Type<&{FungibleToken.Balance, MetadataViews.Resolver}>()): "Metadata public type must include FungibleToken.Balance and MetadataViews.Resolver interfaces."
                providerLinkedType.isSubtype(of: Type<&{FungibleToken.Provider}>()): "Provider type must include FungibleToken.Provider interface."
            }
            self.storagePath = storagePath
            self.receiverPath = receiverPath
            self.metadataPath = metadataPath
            self.providerPath = providerPath
            self.receiverLinkedType = receiverLinkedType
            self.metadataLinkedType = metadataLinkedType
            self.providerLinkedType = providerLinkedType
            self.createEmptyVault = createEmptyVaultFunction
        }
    }

    /// Helper to get FTVaultData in a way that will return a typed Optional.
    ///
    /// @param viewResolver: A reference to the resolver resource
    /// @return A optional FTVaultData struct
    ///
    pub fun getFTVaultData(_ viewResolver: &{MetadataViews.Resolver}): FTVaultData? {
        if let view = viewResolver.resolveView(Type<FTVaultData>()) {
            if let v = view as? FTVaultData {
                return v
            }
        }
        return nil
    }

    /// View to expose the total supply of the Vault's token
    access(all) struct TotalSupply {
        access(all) let supply: UFix64

        init(totalSupply: UFix64) {
            self.supply = totalSupply
        }
    }
}
 