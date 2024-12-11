// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract SolidityTest {
    uint256 storedData; // 状态变量

    constructor() public {
        storedData = 10; // 使用状态变量
    }
}
