// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract PureFunctionExample {
    function add(uint a, uint b) public pure returns (uint) {
        return a + b;
    }
}