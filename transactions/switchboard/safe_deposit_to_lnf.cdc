import FungibleToken from "./../../contracts/FungibleToken.cdc"
import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"
import FiatToken from "./../../contracts/utility/USDC/FiatToken.cdc"
import LostAndFound from "./../../contracts/utility/L&F/LostAndFound.cdc"
import FlowToken from "./../../contracts/utility/FlowToken.cdc"

// This transaction templates how to send USDC funds `to` any `Address` without knowing 
// if it holds a vault of that specific token type. The transaction will attempt to 
// deposit the funds into the receiver's switchboard, if no capability for routing 
// that specific FT is found on the switchboard, then it will deposit the funds 
// into the LostAndFound contract.
//
transaction(to: Address, amount: UFix64) {
    
    // The reference to the vault from the payer's account
    let vaultRef: &FiatToken.Vault
    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault
    // A reference to the signer's flow vault, that will be used for paying the L&F fees 
    let flowProviderRef: &{FungibleToken.Provider}
    // A capability to the signer's flow token receiver for recovering the L&F fees
    let flowReceiver: Capability<&FlowToken.Vault{FungibleToken.Receiver}>?

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        self.vaultRef = signer.borrow<&FiatToken.Vault>(from: FiatToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")
        // Withdraw tokens from the signer's stored vault
        self.sentVault <-self.vaultRef.withdraw(amount: amount)

        // Borrow a reference to the signer's flow token provider
        self.flowProviderRef = signer.borrow<&{FungibleToken.Provider}>(from: /storage/flowTokenVault)
            ?? panic("Could not borrow signer's flow vault provider")

        // Get the capability of the signer's flow token receiver
        self.flowReceiver = signer.getCapability<&FlowToken.Vault{FungibleToken.Receiver}>(/public/flowTokenReceiver)
    
    }

    execute {

        // Get the recipient's public account object
        let recipient = getAccount(to)

        if let switchboardRef = recipient.getCapability(FungibleTokenSwitchboard.PublicPath)
            .borrow<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>() {
            // Attempt to deposit the USDC funds into the receiver's switchboard 
            if let notDepositedVault <-switchboardRef.safeDeposit(from: <- self.sentVault.withdraw(amount: amount)) {
                // If a vault is returned, then their deposit didn't succeed, so we put
                // the funds into Lost And Found
                let memo = "Due royalties"
                let depositEstimate <- LostAndFound.estimateDeposit(redeemer: to, item: <-notDepositedVault, memo: memo, display: nil)
                let storageFee <- self.flowProviderRef.withdraw(amount: depositEstimate.storageFee)
                let resource <- depositEstimate.withdraw()

                LostAndFound.deposit(
                    redeemer: to,
                    item: <-resource,
                    memo: memo,
                    display: nil,
                    storagePayment: &storageFee as &FungibleToken.Vault,
                    flowTokenRepayment: self.flowReceiver
                )

                // Return any remaining storage fees in this vault to the configured
                // flow receiver
                self.flowReceiver!.borrow()!.deposit(from: <-storageFee)
                destroy depositEstimate
            }
            destroy self.sentVault
        } else {
            // If the user did not had a switchboard we deposit the funds into L&F
            let memo = "Due royalties"
            let depositEstimate <- LostAndFound.estimateDeposit(redeemer: to, item: <-self.sentVault, memo: memo, display: nil)
            let storageFee <- self.flowProviderRef.withdraw(amount: depositEstimate.storageFee)
            let resource <- depositEstimate.withdraw()
            LostAndFound.deposit(
                redeemer: to,
                item: <-resource,
                memo: memo,
                display: nil,
                storagePayment: &storageFee as &FungibleToken.Vault,
                flowTokenRepayment: self.flowReceiver
            )
            // Return any remaining storage fees in this vault to the configured
            // flow receiver
            self.flowReceiver!.borrow()!.deposit(from: <-storageFee)
            destroy depositEstimate


        }
    }

}
 