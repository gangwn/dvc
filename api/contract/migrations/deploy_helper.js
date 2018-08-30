class DeployHelper {
    constructor(deployer, serviceManagerArtifact) {
        this.deployer = deployer
        this.serviceManagerArtifact = serviceManagerArtifact
    }

    async deployManager() {
        try {
            this.serviceManager = await this.serviceManagerArtifact.deployed()
        } catch (e) {
            this.serviceManager = await this.deployer.deploy(this.serviceManagerArtifact)
        }

        return this.serviceManager
    }

    async deployContract(artifact, contractName, ...args) {
        await this.deployer.deploy(artifact, ...args)
        const contract = await artifact.deployed()

        await this.serviceManager.registerService(contractName, contract.address)

        return contract;
    }
}

module.exports = DeployHelper