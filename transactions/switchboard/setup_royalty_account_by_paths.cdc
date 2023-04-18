import "FungibleToken"
import "FungibleTokenSwitchboard"
import "FlowToken"
import "FiatToken"


// This transaction is a template for a transaction that could be used by 
// anyone fully setup an account for receiving both Flow and USDC tokens at the same
// public path (for instance the royalties /public/GenericFTReceiver path)
// Using the addNewVaultWrappersByPath switchboard method allows anyone to use
// capability wrappers such as TokenForwarders instead of the actual token vault.
transaction (address: Address) {

    let vaultPaths: [PublicPath]
    let vaultTypes: [Type]    
    let switchboardRef:  &FungibleTokenSwitchboard.Switchboard

    prepare(acct: AuthAccount) {

        // Prepare the paths and types arrays with the Flow and USDC tokens data
        self.vaultPaths = []
        self.vaultPaths.append(/public/flowTokenReceiver)
        self.vaultPaths.append(FiatToken.VaultReceiverPubPath)
        self.vaultTypes = []
        self.vaultTypes.append(Type<@FlowToken.Vault>())
        self.vaultTypes.append(Type<@FiatToken.Vault>())

        // Check if the account already has a Flow Vault
        if acct.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault) == nil {
            // If not, create a new Flow Vault resource and put it into storage
            acct.save(<- FlowToken.createEmptyVault(), to: /storage/flowTokenVault)
        }
        // Check if the receiver capability is linked on the flow receiver path
        if !acct.getCapability<&{FungibleToken.Receiver}>(/public/flowTokenReceiver)
                                                                           .check() {
            // if it's not, create a public capability to the flow vault exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            acct.unlink(/public/flowTokenReceiver)
            acct.link<&{FungibleToken.Receiver}>(/public/flowTokenReceiver, 
                                                     target: /storage/flowTokenVault)
        }

        // Check if the account already has a USDC Vault
        if acct.borrow<&FiatToken.Vault>(from: FiatToken.VaultStoragePath) == nil {
            // If not, create a new USDC Vault resource and put it into storage
            acct.save(<- FiatToken.createEmptyVault(), 
                                                      to: FiatToken.VaultStoragePath)
        }
        // Check if the receiver capability is linked on the USDC receiver path
        if !acct.getCapability<&{FungibleToken.Receiver}>
                                           (FiatToken.VaultReceiverPubPath).check() {
            // if it's not, create a public capability to the USDC vault exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            acct.unlink(FiatToken.VaultReceiverPubPath)
            acct.link<&{FungibleToken.Receiver}>(FiatToken.VaultReceiverPubPath, 
                                                  target: FiatToken.VaultStoragePath)
        }

        // Check if the account already has a Switchboard resource
        if acct.borrow<&FungibleTokenSwitchboard.Switchboard>
                                (from: FungibleTokenSwitchboard.StoragePath) == nil {
            // If not, create a new Switchboard resource and put it into storage
            acct.save(<- FungibleTokenSwitchboard.createSwitchboard(), 
                                            to: FungibleTokenSwitchboard.StoragePath)
        }
        // Check if the receiver capability is linked on the receiver path
        if !acct.getCapability
                      <&FungibleTokenSwitchboard.Switchboard{FungibleToken.Receiver}>
                              (FungibleTokenSwitchboard.ReceiverPublicPath).check() {
            // if it's not, create a public capability to the Switchboard exposing 
            // the deposit function through the {FungibleToken.Receiver} interface
            acct.unlink(FungibleTokenSwitchboard.ReceiverPublicPath)
            acct.link<&FungibleTokenSwitchboard.Switchboard{FungibleToken.Receiver}>(
                                         FungibleTokenSwitchboard.ReceiverPublicPath,
                                        target: FungibleTokenSwitchboard.StoragePath)
        }
        // Check if the SwitchboardPublic and ft receiver capabilities are linked on
        // the switchboard public path
        if !acct.getCapability<
        &FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic, FungibleToken.Receiver}
                                       >(FungibleTokenSwitchboard.ReceiverPublicPath)
                                                                           .check() {
            // if it's not, create a public capability to the Switchboard exposing 
            // both the {FungibleTokenSwitchboard.SwitchboardPublic} and the 
            // {FungibleToken.Receiver} interfaces
            acct.unlink(FungibleTokenSwitchboard.PublicPath)
            acct.link<
            &FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic, FungibleToken.Receiver}
                                               >(FungibleTokenSwitchboard.PublicPath,
                                        target: FungibleTokenSwitchboard.StoragePath)
        }
        // Get a reference to the switchboard
        self.switchboardRef = acct.borrow<&FungibleTokenSwitchboard.Switchboard>
                                         (from: FungibleTokenSwitchboard.StoragePath) 
                                ?? panic("Could not borrow reference to switchboard")

    }

    execute {

      // Add the capabilities to the switchboard using addNewVaultWrappersByPath
      self.switchboardRef.addNewVaultWrappersByPath(paths: self.vaultPaths, 
                                            types: self.vaultTypes, address: address)

    }

}
 