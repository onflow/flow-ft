import TokenSwitchboard from "./../../contracts/TokenSwitchboard.cdc"
import Token from "./../../contracts/Token.cdc"

// This transaction is a template for a transaction that could be used by 
// anyone to to add a Switchboard resource to their account so that they can
// receive multiple fungible tokens using a single {Token.Receiver}
transaction {

    prepare(acct: AuthAccount) {

        // Check if the account already has a Switchboard resource
        if acct.borrow<&TokenSwitchboard.Switchboard>
          (from: TokenSwitchboard.StoragePath) == nil {
            
            // Create a new Switchboard resource and put it into storage
            acct.save(
                <- TokenSwitchboard.createSwitchboard(), 
                to: TokenSwitchboard.StoragePath)

            // Create a public capability to the Switchboard exposing the deposit
            // function through the {Token.Receiver} interface
            acct.link<&TokenSwitchboard.Switchboard{Token.Receiver}>(
                TokenSwitchboard.ReceiverPublicPath,
                target: TokenSwitchboard.StoragePath
            )
            
            // Create a public capability to the Switchboard exposing both the
            // {TokenSwitchboard.SwitchboardPublic} and the 
            // {Token.Receiver} interfaces
            acct.link<&TokenSwitchboard.Switchboard{TokenSwitchboard.SwitchboardPublic, Token.Receiver}>(
                TokenSwitchboard.PublicPath,
                target: TokenSwitchboard.StoragePath
            )
        
        }

    }

}
