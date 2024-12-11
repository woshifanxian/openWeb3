// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract OverloadingExample{
    function f(uint _in) public pure returns(uint){
        return _in;
    }

     function f(uint _in, bool _flag) public pure returns (uint) {
        return _flag ? _in : 0;
    }
}