/// Deploys the TokenForwarding contract with the specified init parameters

transaction(contractName: String,
            code: [UInt8],
            senderStoragePath: StoragePath,
            storagePath: StoragePath,
            publicPath: PublicPath) {

  prepare(signer: AuthAccount) {

    signer.contracts.add(name: contractName, code: code, senderStoragePath, storagePath, publicPath)

  }
}
 
