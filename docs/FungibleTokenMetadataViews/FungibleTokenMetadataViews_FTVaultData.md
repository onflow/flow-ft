# Struct `FTVaultData`

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

### Initializer

```cadence
func init(storagePath StoragePath, receiverPath PublicPath, metadataPath PublicPath, providerPath PrivatePath, receiverLinkedType Type, metadataLinkedType Type, providerLinkedType Type, createEmptyVaultFunction ((): @FungibleToken.Vault))
```


