import "ExampleToken"
import "Burner"

/// Test-only transaction that removes and destroys the signer's ExampleToken.Vault,
/// simulating a scenario where an account's vault is gone and any previously issued
/// receiver capabilities pointing to it are now stale.

transaction {
    prepare(signer: auth(LoadValue) &Account) {
        let vault <- signer.storage.load<@ExampleToken.Vault>(from: ExampleToken.VaultStoragePath)
            ?? panic("No ExampleToken.Vault found at VaultStoragePath")
        Burner.burn(<-vault)
    }
}
