// This script reads the stored vault capabilities from a switchboard on the
// passed account

import FungibleToken from "../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "../../contracts/FungibleTokenSwitchboard.cdc"

pub fun main(account: Address): {Type: Capability<&{FungibleToken.Receiver}>} {
    let acct = getAccount(account)
    // Get a reference to the switchboard conforming to SwitchboardPublic
    let switchboardRef = acct.getCapability(FungibleTokenSwitchboard.SwitchboardPublicPath)
        .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>()
        ?? panic("Could not borrow reference to switchboard")

    return switchboardRef.getVaultCapabilities()
}
 