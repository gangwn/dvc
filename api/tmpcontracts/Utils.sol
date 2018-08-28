pragma solidity ^0.4.23;

library Utils {
    function isNull(string str) private pure returns (bool) {
        bytes memory b = bytes(str);
        if(b.length == 0){
            return true;
        } else {
            return false;
        }
    }
}