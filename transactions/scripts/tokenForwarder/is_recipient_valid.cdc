import TokenForwarding from "TokenForwarding"

access(all) fun main(addr: Address, tokenForwardingPath: PublicPath): Bool {
    let forwarderRef = getAccount(addr)
                       .getCapability<&{TokenForwarding.ForwarderPublic}>(tokenForwardingPath)
                       .borrow()
                       ?? panic("Unable to borrow {TokenForwarding.ForwarderPublic} restrict type from a capability")

    return forwarderRef.check()
}