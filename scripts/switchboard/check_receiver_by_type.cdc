import "FungibleTokenSwitchboard"
import "ExampleToken"

pub fun main(switchboard: Address): Bool {
let switchboardRef = getAccount(switchboard)
    .getCapability<&{FungibleTokenSwitchboard.SwitchboardPublic}>(FungibleTokenSwitchboard.PublicPath)
    .borrow() 
    ?? panic("Unable to borrow capability with restricted type of {FungibleTokenSwitchboard.SwitchboardPublic} from ".concat(switchboard.toString()).concat( "account"))
    return switchboardRef.checkReceiverByType(type: Type<@ExampleToken.Vault>())
}