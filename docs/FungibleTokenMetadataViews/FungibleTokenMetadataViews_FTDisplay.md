# Struct `FTDisplay`

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

### Initializer

```cadence
func init(name String, symbol String, description String, externalURL MetadataViews.ExternalURL, logos MetadataViews.Medias, socials {String: MetadataViews.ExternalURL})
```


