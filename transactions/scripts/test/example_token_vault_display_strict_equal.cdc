import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import MetadataViews from "MetadataViews"
import ViewResolver from "ViewResolver"
import ExampleToken from "ExampleToken"

/// Test helper script to validate ExampleToken serves FTDisplay as expected
///
access(all) fun main(address: Address): Bool {
    let account = getAccount(address)

    let expected = FungibleTokenMetadataViews.FTDisplay(
            name: "Example Fungible Token",
            symbol: "EFT",
            description: "This fungible token is used as an example to help you develop your next FT #onFlow.",
            externalURL: MetadataViews.ExternalURL("https://example-ft.onflow.org"),
            logos: MetadataViews.Medias([
                MetadataViews.Media(
                    file: MetadataViews.HTTPFile(
                        url: "https://assets.website-files.com/5f6294c0c7a8cdd643b1c820/5f6294c0c7a8cda55cb1c936_Flow_Wordmark.svg"
                    ),
                    mediaType: "image/svg+xml"
                )]
            ),
            socials: {
                "twitter": MetadataViews.ExternalURL("https://twitter.com/flow_blockchain")
            }
        )

    let vaultData = ExampleToken.resolveContractView(resourceType: nil, viewType: Type<FungibleTokenMetadataViews.FTVaultData>()) as! FungibleTokenMetadataViews.FTVaultData?
        ?? panic("Could not get vault data view for the contract")

    let vaultRef = account.capabilities.borrow<&{FungibleToken.Balance}>(vaultData.metadataPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let actual = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
        ?? panic("Token does not implement FTDisplay view")

    assert(actual.logos.items.length == expected.logos.items.length, message: "Medias length mismatch")

    let expectedLogoMedia = expected.logos.items[0]
    let actualLogoMedia = actual.logos.items[0]
    return expected.name == actual.name && expected.symbol == actual.symbol &&
        expected.description == actual.description && expected.externalURL.url == actual.externalURL.url && 
        expectedLogoMedia.file.uri() == actualLogoMedia.file.uri() && expectedLogoMedia.mediaType == actualLogoMedia.mediaType
}

