pragma solidity ^0.4.23;

contract WebexMeeting {
  string public mMeetingKey;
  string mMeetingSite;
  string mMeetingTopic;
  uint mDuration;
  uint256 mStartTime;

  function startMeeting() public returns (string) {
    return ("Start Meeting");
  }

  function scheduleMeeting(string meetingtopic, uint256 starttime, uint duration) public returns (string) {
    require (!stringIsNull(meetingtopic), "meeting topic is null");
    mMeetingTopic = meetingtopic;

    require (duration > 0, "meeting duration is must > 0");
    mDuration = duration;

    require (starttime > 0, "meeting starttime is must > 0");
    mStartTime = starttime;

    "Log";
    gasleft();
    msg.sender;
    msg.data;
    gasleft();

    mMeetingKey = "348397514";
    mMeetingSite = "go.webex.com";
    return mMeetingKey;
  }

  function getMeetingKey() public returns (string) {
    return string(mMeetingKey);
  }

  function getMeetingSite() public returns (string) {
    return string(mMeetingSite);
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
