pragma solidity ^0.4.23;

import "./Utils.sol";
import "./IService.sol";

contract ConferenceService is IService {

    enum ConferenceStatus {Inactive, Active}

    event ConferenceScheduled(address creatorAddr, string confId, string topic, uint256 startTime, uint duration, address[] invitees);

    mapping(string => Conference) private conferences;

    function scheduleConference(string confId, string topic, uint256 startTime, uint duration, address[] invitees) external {
        require (!Utils.isNull(topic), "meeting topic is null");
        require (duration > 0, "meeting duration must > 0");
        require (startTime > 0, "meeting start time must > 0");

        Conference storage conf = conferences[confId];
        conf.confId = confId;
        conf.creatorAddr = msg.sender;
        conf.topic = topic;
        conf.startTime = startTime;
        conf.duration = duration;
        conf.invitees = invitees;

        // to do: store this conference to userservice contract


        emit ConferenceScheduled(msg.sender, confId, topic, startTime, duration, invitees);
    }


}
