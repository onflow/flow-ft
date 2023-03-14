# Struct `FTView`

```cadence
struct FTView {

    ftDisplay:  FTDisplay?

    ftVaultData:  FTVaultData?
}
```

FTView wraps FTDisplay and FTVaultData, and is used to give a complete
picture of a Fungible Token. Most Fungible Token contracts should
implement this view.

### Initializer

```cadence
func init(ftDisplay FTDisplay?, ftVaultData FTVaultData?)
```


