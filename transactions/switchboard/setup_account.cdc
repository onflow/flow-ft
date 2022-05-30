import FungibleTokenSwitchboard from "./../../contracts/FungibleTokenSwitchboard.cdc"


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









