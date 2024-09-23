import "FungibleToken"
import "FungibleTokenSwitchboard"

/// This script reads the stored vault capabilities from a switchboard on the passed account
///
access(all) fun main(account: Address): {Type: Address} {

    // Get a reference to the switchboard conforming to SwitchboardPublic
    let switchboardRef = getAccount(account).capabilities.borrow<&{FungibleTokenSwitchboard.SwitchboardPublic}>(
            FungibleTokenSwitchboard.PublicPath)
        ?? panic("The signer does not store a FungibleToken Switchboard capability at the path "
            .concat(FungibleTokenSwitchboard.PublicPath.toString())
            .concat(". The signer must initialize their account with this object first!"))

    // Return the result of `getVaultTypesWithAddress()`
    return switchboardRef.getVaultTypesWithAddress()

}
