pragma solidity ^0.4.23;

import "./ServiceManager.sol";

contract IService {

    ServiceManager public serviceManager;

    struct Conference {
        string confId;
        address creatorAddr;  // the meeting creator address
        string topic;
        uint256 startTime;
        uint duration;        // seconds
        address[] invitees;
    }

    constructor(address serviceManagerAddr) public {
        serviceManager = ServiceManager(serviceManagerAddr);
    }

}