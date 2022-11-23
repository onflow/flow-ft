# Contract `FungibleTokenMetadataViews`

```cadence
contract FungibleTokenMetadataViews {
}
```

This contract implements the metadata standard proposed
in FLIP-1087.

Ref: https://github.com/onflow/flow/blob/master/flips/20220811-fungible-tokens-metadata.md

Structs and resources can implement one or more
metadata types, called views. Each view type represents
a different kind of metadata.
## Structs & Resources

### struct `FTView`

```cadence
struct FTView {

    ftDisplay:  FTDisplay?

    ftVaultData:  FTVaultData?
}
```
FTView wraps FTDisplay and FTVaultData, and is used to give a complete
picture of a Fungible Token. Most Fungible Token contracts should
implement this view.

[More...](FungibleTokenMetadataViews_FTView.md)

---

### struct `FTDisplay`

```cadence
struct FTDisplay {

    name:  String

    symbol:  String

    description:  String

    externalURL:  MetadataViews.ExternalURL

    logos:  MetadataViews.Medias

    socials:  {String: MetadataViews.ExternalURL}
}
```
View to expose the information needed to showcase this FT.
This can be used by applications to give an overview and
graphics of the FT.

[More...](FungibleTokenMetadataViews_FTDisplay.md)

---

### struct `FTVaultData`

```cadence
struct FTVaultData {

    storagePath:  StoragePath

    receiverPath:  PublicPath

    metadataPath:  PublicPath

    providerPath:  PrivatePath

    receiverLinkedType:  Type

    metadataLinkedType:  Type

    providerLinkedType:  Type

    createEmptyVault:  ((): @FungibleToken.Vault)
}
```
View to expose the information needed store and interact with a FT vault.
This can be used by applications to setup a FT vault with proper
storage and public capabilities.

[More...](FungibleTokenMetadataViews_FTVaultData.md)

---
## Functions

### fun `getFTView()`

```cadence
func getFTView(viewResolver &{MetadataViews.Resolver}): FTView
```
Helper to get a FT view.

Parameters:
  - viewResolver : _A reference to the resolver resource_

Returns: A FTView struct

---

### fun `getFTDisplay()`

```cadence
func getFTDisplay(_ &{MetadataViews.Resolver}): FTDisplay?
```
Helper to get FTDisplay in a way that will return a typed optional.

Parameters:
  - viewResolver : _A reference to the resolver resource_

Returns: An optional FTDisplay struct

---

### fun `getFTVaultData()`

```cadence
func getFTVaultData(_ &{MetadataViews.Resolver}): FTVaultData?
```
Helper to get FTVaultData in a way that will return a typed Optional.

Parameters:
  - viewResolver : _A reference to the resolver resource_

Returns: A optional FTVaultData struct

---
