var Utils = artifacts.require("./Utils.sol");
var IService = artifacts.require("./IService.sol");
var ConferenceService = artifacts.require("./ConferenceService.sol");
var UserService = artifacts.require("./UserService.sol");


module.exports = function(deployer) {
    deployer.deploy(Utils);
    deployer.link(Utils, ConferenceService);
    deployer.deploy(IService);
    deployer.deploy(ConferenceService);
    deployer.deploy(UserService);
};