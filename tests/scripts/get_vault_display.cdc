// This script checks the FTDisplay view from ExampleToken
// is the expected one. This is merely used in testing,
// since we cannot return on-chain types to the test
// files yet.

import ExampleToken from "ExampleToken"
import FungibleTokenMetadataViews from "FungibleTokenMetadataViews"
import ViewResolver from "ViewResolver"

access(all) fun main(address: Address): Bool {
    let account = getAccount(address)

    let vaultRef = account.capabilities.borrow<&{ViewResolver.Resolver}>(ExampleToken.VaultPublicPath)
        ?? panic("Could not borrow a reference to the vault resolver")

    let ftDisplay = FungibleTokenMetadataViews.getFTDisplay(vaultRef)
        ?? panic("Token does not implement FTDisplay view")

    assert("Example Fungible Token" == ftDisplay.name)
    assert("EFT" == ftDisplay.symbol)
    assert("This fungible token is used as an example to help you develop your next FT #onFlow." == ftDisplay.description)
    assert("https://example-ft.onflow.org" == ftDisplay.externalURL!.url)
    assert("https://twitter.com/flow_blockchain" == ftDisplay.socials["twitter"]!.url)
    assert("https://assets.website-files.com/5f6294c0c7a8cdd643b1c820/5f6294c0c7a8cda55cb1c936_Flow_Wordmark.svg" == ftDisplay.logos.items[0].file.uri())

    return true
}
