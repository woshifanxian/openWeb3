// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract SolidityTest2 {
    uint256 storedData; // 状态变量

    constructor() public {
        storedData = 10;
    }

    function getResult() public view returns (uint256) {
        uint256 a = 1; // 局部变量
        uint256 b = 2;
        uint256 result = a + b;
        return result; // 访问局部变量
    }
}
