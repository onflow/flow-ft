import "FungibleToken"

/// This scripts returns the supported FungibleToken's type by the provided `target` address.
/// `target` address should hold the capability which conforms with FungibleToken.Receiver restricted type
/// while it doesn't matter whether capability refers to fungible token or a custom receiver like 
/// `FungibleTokenSwitchboard` or `TokenReceiver`. However `targetPath` tells where the capability stores
access(all) fun main(target: Address, targetPath: PublicPath): {Type: Bool} {

    // Access the capability for the provided target address
    let capabilityRef = getAccount(target).capabilities.borrow<&{FungibleToken.Receiver}>(targetPath)
        ?? panic("Unable to borrow capability with restricted sub type {FungibleToken.Receiver} from path".concat(targetPath.toString()))
    // Return the supported vault types.
    return capabilityRef.getSupportedVaultTypes()
}
