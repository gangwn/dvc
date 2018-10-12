pragma solidity ^0.4.23;

import "./IService.sol";

import 'LibCLL/LibCLL.sol';

contract CCSService is IService {

    using LibCLLu for LibCLLu.CLL;

    struct CCS {
        string ip;
        int port;
        string peerId;
        string[] servedConfs;
        uint ccsId;
    }

    LibCLLu.CLL private ccsIdList;
    mapping(uint => address) private ccsIdAddressMap;

    uint ccsNum;

    mapping(address => CCS) private ccss;

    struct Job {
            address ccsAddress;
            string confId;
        }

    mapping(string => Job) private jobs;

    event NewJobCreated(string confId, address ccsAddress, string ip, int port, string peerId);

    constructor(address serviceManagerAddr) public IService(serviceManagerAddr) {}

    function registerCCS(string ip, int port, string peerId) external {
        CCS storage ccs = ccss[msg.sender];

        require(ccs.ccsId == 0);

        ccs.ip = ip;
        ccs.port = port;
        ccs.peerId = peerId;

        ccsNum = ccsNum + 1;
        ccs.ccsId = ccsNum;

        ccsIdAddressMap[ccs.ccsId] = msg.sender;

        LibCLLu.push(ccsIdList, ccs.ccsId, true);
    }

    function getFirstCCS() view external returns (address, string, int, string) {
        uint[2] memory first = LibCLLu.getNode(ccsIdList, 0);
        require(first[1] > 0);

        address addr = ccsIdAddressMap[first[1]];
        CCS storage ccs = ccss[addr];

        return (addr, ccs.ip, ccs.port, ccs.peerId);
    }

    function getNextCCS(address addr) view external returns (address, string, int, string) {
        CCS storage ccs = ccss[addr];
        if (ccs.ccsId == 0) {
            return (address(0), "", 0, "");
        }

        uint[2] memory nebor = LibCLLu.getNode(ccsIdList, ccs.ccsId);
        if (nebor[1] == 0) {
            return (address(0), "", 0, "");
        }

        address nextAddr = ccsIdAddressMap[nebor[1]];
        CCS storage ccsNext = ccss[nextAddr];

        return (nextAddr, ccsNext.ip, ccsNext.port, ccs.peerId);
    }

    function newJob(string confId, address ccsAddress) external  {
        Job storage job = jobs[confId];

        require(job.ccsAddress == address(0));

        job.ccsAddress = ccsAddress;

        CCS storage ccs = ccss[job.ccsAddress];
        ccs.servedConfs.push(confId);

        emit NewJobCreated(confId, ccsAddress, ccs.ip, ccs.port, ccs.peerId);
    }

    function getCCS(string confId) view external returns (address, string, int, string) {
        Job storage job = jobs[confId];
        require(job.ccsAddress != address(0));

        CCS storage ccs = ccss[job.ccsAddress];
        return (job.ccsAddress, ccs.ip, ccs.port, ccs.peerId);
    }

}
