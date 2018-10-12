const DeployHelper = require('./deploy_helper.js');

var ServiceManager = artifacts.require("./ServiceManager.sol");
var Utils = artifacts.require("./Utils.sol");
var ConferenceService = artifacts.require("./ConferenceService.sol");
var CCSService = artifacts.require("./CCSService.sol");
var Token = artifacts.require("./FixedSupplyToken.sol");
module.exports = function(deployer) {
    deployer.then(async () => {
        deployer.logger.log("Start to deploy contracts")

        const deployHelper = new DeployHelper(deployer, ServiceManager)

        const serviceManager = await  deployHelper.deployManager()

        await deployer.deploy(Utils);
        await deployer.link(Utils, ConferenceService);
        await deployer.deploy(Token);

        await deployHelper.deployContract(Token, "FixedSupplyToken")
        await deployHelper.deployContract(ConferenceService, "ConferenceService", serviceManager.address)

        await deployHelper.deployContract(CCSService, "CCSService", serviceManager.address)
    })
};
