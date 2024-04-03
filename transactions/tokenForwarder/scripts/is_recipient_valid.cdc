import "TokenForwarding"

access(all) fun main(addr: Address, tokenForwardingPath: PublicPath): Bool {
    let forwarderRef = getAccount(addr)
                       .capabilities.borrow<&{TokenForwarding.ForwarderPublic}>(tokenForwardingPath)

    return forwarderRef.check()
}