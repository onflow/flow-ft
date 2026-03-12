/// Burner is a contract that can facilitate the destruction of any resource on flow.
///
/// Contributors
/// - Austin Kline - https://twitter.com/austin_flowty
/// - Deniz Edincik - https://twitter.com/bluesign
/// - Bastian Müller - https://twitter.com/turbolent
access(all) contract Burner {
    /// Burnable is an interface that replaces the custom destructor feature removed in Cadence 1.0.
    /// It allows resource authors to add a callback that fires when their resource is destroyed,
    /// ensuring they can enforce invariants (e.g. "don't destroy a non-empty vault") or
    /// perform bookkeeping (e.g. updating the total supply of a fungible token).
    ///
    /// NOTE: The only way to see benefit from this interface
    /// is to always use the burn method in this contract. Anyone who owns a resource can always elect **not**
    /// to destroy a resource this way
    access(all) resource interface Burnable {
        access(contract) fun burnCallback()
    }

    /// burn is a global method which will destroy any resource it is given.
    /// If the provided resource implements the Burnable interface,
    /// it will call the burnCallback method and then destroy afterwards.
    access(all) fun burn(_ toBurn: @AnyResource?) {
        if toBurn == nil {
            destroy toBurn
            return
        }
        let r <- toBurn!

        if let s <- r as? @{Burnable} {
            s.burnCallback()
            destroy s
        } else if let arr <- r as? @[AnyResource] {
            while arr.length > 0 {
                let item <- arr.removeFirst()
                self.burn(<-item)
            }
            destroy arr
        } else if let dict <- r as? @{HashableStruct: AnyResource} {
            let keys = dict.keys
            while keys.length > 0 {
                let item <- dict.remove(key: keys.removeFirst())!
                self.burn(<-item)
            }
            destroy dict
        } else {
            destroy r
        }
    }
}