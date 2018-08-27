pragma solidity ^0.4.23;

contract WebexMeeting {
  //struct User {
  //  address addr;
  //  uint amount;
  //  uint userId;
  //  string token;
  //  string displayName;
  //}

  struct Meeting {
    address addr;  // the meeting host address
    string topic;
    uint256 startTime;
    uint duration;
    string meetingKey;
    uint128 confId;
    string invitees;
  }


  mapping(address => Meeting[]) public meetings;

  function startMeeting() public returns (string) {
    return ("Start Meeting");
  }

  function scheduleMeeting(string topic, uint256 starttime, uint duration, string invites) public returns (string) {
    require (!stringIsNull(topic), "meeting topic is null");

    require (duration > 0, "meeting duration is must > 0");

    require (starttime > 0, "meeting starttime is must > 0");

    gasleft();
    msg.sender;
    msg.data;
    gasleft();

    Meeting memory meeting = Meeting(msg.sender, topic, starttime, duration, "348397514", 1231231232123, "");
    Meeting[] myMeetings = meetings[msg.sender];
    myMeetings[myMeetings.length++] = meeting;
    return meeting.meetingKey;
  }

  function getUpcomingMeeting() public returns (string){
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

  function stringIsNull(string str) private returns (bool) {
      bytes memory strb = bytes(str);
      if(strb.length == 0){
        return true;
      } else {
        return false;
      }
  }
}
