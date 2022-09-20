import FungibleToken from "./FungibleToken.cdc"
import MetadataViews from "./utilityContracts/MetadataViews.cdc"

/// This contract implements the metadata standard proposed
/// in FLIP-1087.
/// 
/// Ref: https://github.com/onflow/flow/blob/master/flips/20220811-fungible-tokens-metadata.md
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

    /// Helper to get a FT view 
    ///
    /// @param viewResolver: A reference to the resolver resource
    /// @return A FTView struct
    ///
    pub fun getFTView(viewResolver: &{MetadataViews.Resolver}): FTView {
        let ftView = viewResolver.resolveView(Type<FTView>())
        if ftView != nil {
            return ftView! as! FTView
        }
        return FTView(
            ftDisplay: self.getFTDisplay(viewResolver),
            ftVaultData : self.getFTVaultData(viewResolver)
        )
    }

    /// View to expose the information needed to showcase this FT. 
    /// This can be used by applications to give an overview and 
    /// graphics of the FT.
    ///
    pub struct FTDisplay {
        /// Name that should be used when displaying this FT.
        pub let name: String

        /// Symbol that could be used as a shorter name for the FT.
        pub let symbol: String

        /// Description that should be used to give an overview of this FT.
        pub let description: String

        /// External link to a URL to view more information about the fungible token.
        pub let externalURL: MetadataViews.ExternalURL

        /// Image to represent the fungible token logo.
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

    /// Helper to get FTDisplay in a way that will return a typed optional
    /// 
    /// @param viewResolver: A reference to the resolver resource
    /// @return A optional FTDisplay struct
    ///
    pub fun getFTDisplay(_ viewResolver: &{MetadataViews.Resolver}): FTDisplay? {
        if let view = viewResolver.resolveView(Type<FTDisplay>()) {
            if let v = view as? FTDisplay {
                return v
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

        /// Public path which must be linked to expose public capabilities of
        /// this FT, including standard FT interfaces and metadataviews 
        /// interfaces
        pub let publicPath: PublicPath

        /// Private path which should be linked to expose the provider
        /// capability to withdraw funds from the vault
        pub let providerPath: PrivatePath

        /// Type that should be linked at the aforementioned public path. This 
        /// is normally a restricted type with many interfaces. Notably the 
        /// `FungibleToken.Balance`, `FungibleToken.Receiver`, and  
        /// `MetadataViews.Resolver` interfaces are required.
        pub let publicLinkedType: Type

        /// Type that should be linked at the aforementioned private path. This 
        /// is normally a restricted type with at a minimum the  
        /// `FungibleToken.Provider` interface
        pub let providerLinkedType: Type

        /// Function that allows creation of an empty FT vault that is intended
        /// to store the funds.
        pub let createEmptyVault: ((): @FungibleToken.Vault)

        init(
            storagePath: StoragePath,
            publicPath: PublicPath,
            providerPath: PrivatePath,
            publicLinkedType: Type,
            providerLinkedType: Type,
            createEmptyVaultFunction: ((): @FungibleToken.Vault)
        ) {
            pre {
                publicLinkedType.isSubtype(of: Type<&{FungibleToken.Receiver, FungibleToken.Balance, MetadataViews.Resolver}>()): "Public type must include FungibleToken.Receiver, FungibleToken.Balance and MetadataViews.Resolver interfaces."
                providerLinkedType.isSubtype(of: Type<&{FungibleToken.Provider, MetadataViews.Resolver}>()): "Provider type must include FungibleToken.Provider and MetadataViews.Resolver interface."
            }
            self.storagePath = storagePath
            self.publicPath = publicPath
            self.providerPath = providerPath
            self.publicLinkedType = publicLinkedType
            self.providerLinkedType = providerLinkedType
            self.createEmptyVault = createEmptyVaultFunction
        }
    }

    /// Helper to get FTVaultData in a way that will return a typed Optional
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

}
 