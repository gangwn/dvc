pragma solidity ^0.4.23;

import "./Utils.sol";
import "./IService.sol";

contract ConferenceService is IService {

    enum ConferenceStatus {Inactive, Active}

    event ConferenceScheduled(address creatorAddr, string confId, string topic, uint256 startTime, uint duration, address[] invitees);

    mapping(string => Conference) private conferences;

    constructor(address serviceManagerAddr) public IService(serviceManagerAddr) {}

    struct User {
        address addr;
        string displayName;
        Conference[] conferences;           // to do: need to sort
    }

    mapping(address => User) private users;

    function scheduleConference(string confId, string topic, uint256 startTime, uint duration, address[] invitees) external {
        //require (!Utils.isNull(topic), "meeting topic is null");
        //require (duration > 0, "meeting duration must > 0");
        //require (startTime > 0, "meeting start time must > 0");

        Conference storage conf = conferences[confId];
        conf.confId = confId;
        conf.creatorAddr = msg.sender;
        conf.topic = topic;
        conf.startTime = startTime;
        conf.duration = duration;
        conf.invitees = invitees;

        //insertConferenceToUser(conf);

        emit ConferenceScheduled(msg.sender, confId, topic, startTime, duration, invitees);
    }

    function getConferences(uint256 timeAfter) external view returns (string, address, string, uint256, uint, uint) {
        User storage user = users[msg.sender];

        for (uint i = 0; i < user.conferences.length; i++) {
            if (timeAfter > user.conferences[i].startTime) {
                continue;
            }

            Conference storage conference = user.conferences[i];
            return (conference.confId, conference.creatorAddr, conference.topic, conference.startTime,
                conference.duration, conference.invitees.length);
        }

        return ("", address(0), "", 0, 0, 0);
    }

    function insertConferenceToUser(Conference conf) internal {
        User storage user = users[msg.sender];
        user.conferences.push(conf);

        for (uint i = 0; i < conf.invitees.length; i++) {
            User storage invitee = users[conf.invitees[i]];
            invitee.conferences.push(conf);
        }
    }
}
