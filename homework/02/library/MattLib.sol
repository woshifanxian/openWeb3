// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
library MathLib{

    function add(uint a,uint b) external pure returns (uint){
         uint c = a + b;
        require(c >= a, "SafeMath: addition overflow");
        return c;
    }

    function sub(uint a,uint b) external pure returns (int){
        int c = int(a - b);
        require(c <= int(a), "SafeMath: addition overflow");
        return c;
    }

    function multi(uint a,uint b) external pure returns (uint){
        uint c = a * b;
        require(c >= a, "SafeMath: addition overflow");
        return c;
    }
}