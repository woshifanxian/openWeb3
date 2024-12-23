// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract OverflowExample {
    function add(uint8 a, uint8 b) public pure returns (uint8) {
        return a + b; // 如果发生溢出，会抛出异常
    }

    function subtract(uint8 a, uint8 b) public pure returns (uint8) {
        return a - b; // 如果发生下溢，会抛出异常
    }
}
