import FungibleToken from "FungibleToken"
import FungibleTokenSwitchboard from "FungibleTokenSwitchboard"
import FlowToken from "FlowToken"
import FiatToken from "FiatToken"


// This transaction is a template for a transaction that could be used by 
// anyone fully setup an account for receiving both Flow and USDC tokens at the same
// public path (for instance the royalties /public/GenericFTReceiver path)
// Using the addNewVaultWrappersByPath switchboard method allows anyone to use
// capability wrappers such as TokenForwarders instead of the actual token vault.
transaction () {

    let flowTokenVaultCapability: Capability<&{FungibleToken.Receiver}>
    let fiatTokenVaultCapability: Capability<&{FungibleToken.Receiver}>   
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(signer: AuthAccount) {

        // Check if the account already has a Flow Vault
        if signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault) == nil {
            // If not, create a new Flow Vault resource and put it into storage
            signer.save(<- FlowToken.createEmptyVault(vaultType: Type<@FlowToken.Vault>()), to: /storage/flowTokenVault)
        }
        // Check if the receiver capability is linked on the flow receiver path
        if !signer.getCapability<&{FungibleToken.Receiver}>(/public/flowTokenReceiver)
                                                                           .check() {
            // if it's not, create a public capability to the flow vault exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            signer.unlink(/public/flowTokenReceiver)
            signer.link<&{FungibleToken.Receiver}>(/public/flowTokenReceiver, 
                                                     target: /storage/flowTokenVault)
        }
        self.flowTokenVaultCapability = signer.getCapability<&{FungibleToken.Receiver}>(/public/flowTokenReceiver)

        // Check if the account already has a USDC Vault
        if signer.borrow<&FiatToken.Vault>(from: FiatToken.VaultStoragePath) == nil {
            // If not, create a new USDC Vault resource and put it into storage
            signer.save(<- FiatToken.createEmptyVault(vaultType: Type<@FiatToken.Vault>()), 
                                                      to: FiatToken.VaultStoragePath)
        }
        // Check if the receiver capability is linked on the USDC receiver path
        if !signer.getCapability<&{FungibleToken.Receiver}>
                                           (FiatToken.VaultReceiverPubPath).check() {
            // if it's not, create a public capability to the USDC vault exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            signer.unlink(FiatToken.VaultReceiverPubPath)
            signer.link<&{FungibleToken.Receiver}>(FiatToken.VaultReceiverPubPath, 
                                                  target: FiatToken.VaultStoragePath)
        }
        self.fiatTokenVaultCapability = signer.getCapability<&{FungibleToken.Receiver}>(FiatToken.VaultReceiverPubPath)
        
        // Check if the account already has a Switchboard resource
        if signer.borrow<&FungibleTokenSwitchboard.Switchboard>
                                (from: FungibleTokenSwitchboard.StoragePath) == nil {
            // If not, create a new Switchboard resource and put it into storage
            signer.save(<- FungibleTokenSwitchboard.createSwitchboard(), 
                                            to: FungibleTokenSwitchboard.StoragePath)
        }
        // Check if the receiver capability is linked on the receiver path
        if !signer.getCapability
                      <&FungibleTokenSwitchboard.Switchboard{FungibleToken.Receiver}>
                              (FungibleTokenSwitchboard.ReceiverPublicPath).check() {
            // if it's not, create a public capability to the Switchboard exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            signer.unlink(FungibleTokenSwitchboard.ReceiverPublicPath)
            signer.link<&FungibleTokenSwitchboard.Switchboard{FungibleToken.Receiver}>(
                                         FungibleTokenSwitchboard.ReceiverPublicPath,
                                        target: FungibleTokenSwitchboard.StoragePath)
        }
        // Check if the SwitchboardPublic and ft receiver capabilities are linked on
        // the switchboard public path
        if !signer.getCapability<
        &FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic, FungibleToken.Receiver}
                                       >(FungibleTokenSwitchboard.ReceiverPublicPath)
                                                                           .check() {
            // if it's not, create a public capability to the Switchboard exposing 
            // both the {FungibleTokenSwitchboard.SwitchboardPublic} and the 
            // {FungibleToken.Receiver} interfaces
            signer.unlink(FungibleTokenSwitchboard.PublicPath)
            signer.link<
            &FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic, FungibleToken.Receiver}
                                               >(FungibleTokenSwitchboard.PublicPath,
                                        target: FungibleTokenSwitchboard.StoragePath)
        }
        // Get a reference to the switchboard
        self.switchboardRef = signer.borrow<&FungibleTokenSwitchboard.Switchboard>
                                         (from: FungibleTokenSwitchboard.StoragePath) 
                                ?? panic("Could not borrow reference to switchboard")

    }

    execute {
        // Add the capability to the switchboard using addNewVault method
        self.switchboardRef.addNewVault(capability: self.flowTokenVaultCapability)
        self.switchboardRef.addNewVault(capability: self.fiatTokenVaultCapability)
    }

}
 