pragma solidity ^0.4.23;

import "./IService.sol";

contract CCSService is IService {

    struct CCS {
        string ip;
        int port;
        string[] servedConfs;
    }

    struct Job {
        address ccsAddress;
        string confId;
    }

    mapping(address => CCS) private ccss;

    mapping(string => Job) private jobs;

    event NewJobCreated(string confId, address ccsAddress, string ip, int port);

    constructor(address serviceManagerAddr) public IService(serviceManagerAddr) {}

    function registerCCS(string ip, int port) external {
        CCS storage ccs = ccss[msg.sender];
        ccs.ip = ip;
        ccs.port = port;
    }

    function newJob(string confId, address ccsAddress) external  {
        Job storage job = jobs[confId];

        require(job.ccsAddress == address(0));

        CCS storage ccs = ccss[job.ccsAddress];
        ccs.servedConfs.push(confId);

        emit NewJobCreated(confId, ccsAddress, ccs.ip, ccs.port);
    }

    function getCCS(string confId) view external returns (address, string, int) {
        Job storage job = jobs[confId];
        require(job.ccsAddress != address(0));

        CCS storage ccs = ccss[job.ccsAddress];
        return (job.ccsAddress, ccs.ip, ccs.port);
    }
}
