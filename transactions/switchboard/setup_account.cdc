// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the exampleToken

import FungibleTokenSwitchboard from "./../../contracts/fungibleTokenSwitchBoard/FungibleTokenSwitchboard.cdc"

transaction {

    prepare(acct: AuthAccount) {

        if acct.borrow<&FungibleTokenSwitchboard.Switchboard>(from: FungibleTokenSwitchboard.SwitchboardStoragePath) == nil {
            
            let switchboard <- FungibleTokenSwitchboard.createSwitchboard() as! @FungibleTokenSwitchboard.Switchboard

            acct.save(<- switchboard, to: FungibleTokenSwitchboard.SwitchboardStoragePath)

            acct.link<&FungibleTokenSwitchboard.Switchboard{FungibleTokenSwitchboard.SwitchboardPublic}>(
                FungibleTokenSwitchboard.SwitchboardPublicPath,
                target: FungibleTokenSwitchboard.SwitchboardStoragePath
            )
        }

    }
}









