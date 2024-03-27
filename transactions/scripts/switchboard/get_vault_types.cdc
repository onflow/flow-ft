import "FungibleToken"
import "FungibleTokenSwitchboard"

// This script reads the stored vault capabilities from a switchboard on the
// passed account
access(all) fun main(account: Address): [Type] {

    let acct = getAccount(account)

    // Get a reference to the switchboard conforming to SwitchboardPublic
    let switchboardRef = acct.getCapability(FungibleTokenSwitchboard.PublicPath)
        .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
        ?? panic("Could not borrow reference to switchboard")

    // Return the result of `getVaultTypes()`
    return switchboardRef.getVaultTypes()

}
