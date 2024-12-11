// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract testInt {
    int8 a = -1;
    int16 b = 2;
    uint32 c = 10;
    uint8 d = 16;

    function add(uint256 x, uint256 y) public pure returns (uint256 z) {
        z = x + y;
    }

    function divide(uint256 x, uint256 y) public pure returns (uint256 z) {
        z = x / y;
    }

    function leftshift(int256 x, uint256 y) public pure returns (int256 z) {
        z = x << y;
    }

    function rightshift(int256 x, uint256 y) public pure returns (int256 z) {
        z = x >> y;
    }

    function testPlusPlus() public pure returns (uint256) {
        uint256 x = 1;
        uint256 y = ++x; // c = ++a;
        return y;
    }
}
