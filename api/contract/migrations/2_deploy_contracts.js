const DeployHelper = require('./deploy_helper.js');

var ServiceManager = artifacts.require("./ServiceManager.sol");
var Utils = artifacts.require("./Utils.sol");
var ConferenceService = artifacts.require("./ConferenceService.sol");

module.exports = function(deployer) {
    deployer.logger.log("Start to deploy contracts")

    const deployHelper = new DeployHelper(deployer, ServiceManager)

    const serviceManager =  deployHelper.deployManager()

    deployer.deploy(Utils);
    deployer.link(Utils, ConferenceService);

    deployHelper.deployContract(ConferenceService, "ConferenceService", serviceManager.address)
};