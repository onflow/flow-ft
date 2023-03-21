import TokenForwarding from "../../../contracts/utility/TokenForwarding.cdc"

pub fun main(addr: Address, tokenForwardingPath: PublicPath): Bool {
    let forwarderRef = getAccount(addr)
                       .getCapability<&{TokenForwarding.ForwarderPublic}>(tokenForwardingPath)
                       .borrow()
                       ?? panic("Unable to borrow {TokenForwarding.ForwarderPublic} restrict type from a capability")
    
    return forwarderRef.check()
}