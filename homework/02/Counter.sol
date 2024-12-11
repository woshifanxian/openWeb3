// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract Counter{
    uint256 public count;

    //function to get current count
    function get() public view returns(uint256){
        return count;
    } 

    //function to increment count by 1
    function inc() public {
        count +=1;
    }

    //function to decrement count by 1
    function dec() public{
        count -= 1;
    }
}