pragma solidity ^0.4.23;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract ServiceManager is Ownable {

    struct ServiceInfo {
        address contractAddress;
    }

    mapping (string => ServiceInfo) private services;

    function registerService(string serviceName, address contractAddress) external onlyOwner {
        ServiceInfo storage service = services[serviceName];
        service.contractAddress = contractAddress;
    }

    function getService(string serviceName) external view returns (address) {
        return services[serviceName].contractAddress;
    }
}