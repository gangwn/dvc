pragma solidity ^0.4.23;

import "./Utils.sol";

contract ConferenceManager {
    struct Conference {
        string confId;
        address creatorAddr;  // the meeting creator address
        string topic;
        uint256 startTime;
        uint duration;        // seconds
        address[] invitees;
    }

    enum ConferenceStatus {Inactive, Active};

    event ConferenceScheduled(address creatorAddr, string confId, string topic, uint256 startTime, uint duration, address[] invitees);

    mapping(string => Conference[]) public conferences;

    function scheduleConference(string confId, string topic, uint256 startTime, uint duration, address[] invitees) public {
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

        ConferenceScheduled(msg.sender, confId, topic, startTime, duration, invitees);
    }


    function startConference() public returns (string) {
        return ("Start Meeting");
    }


  function getConferences() public view returns (string){
    address addr = msg.sender;
    Meeting[] storage myMeetings = meetings[addr];

    if(myMeetings.length == 0){
      return "";
    }

    uint256 nearByStartTime = 999999999999999999999999;
    uint upcomingIndex = 0;
    for (uint i = 0; i < myMeetings.length; i++) {
      if(nearByStartTime > myMeetings[i].startTime){
         upcomingIndex = i;
         nearByStartTime = myMeetings[i].startTime;
      }

      return myMeetings[upcomingIndex].meetingKey;
    }
  }
}
