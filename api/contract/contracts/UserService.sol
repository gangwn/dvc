pragma solidity ^0.4.23;

import "./IService.sol";


contract UserService is IService {
	struct User {
		address addr;
		string displayName;
		Conference[] conferences;           // to do: need to sort
	}

    mapping(address => User) private users;

    function getConferences(uint256 timeAfter) public view returns (string, address, string, uint256, uint, uint){
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
}
