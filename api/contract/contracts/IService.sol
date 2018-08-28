pragma solidity ^0.4.23;

contract IService {

    struct Conference {
        string confId;
        address creatorAddr;  // the meeting creator address
        string topic;
        uint256 startTime;
        uint duration;        // seconds
        address[] invitees;
    }

}