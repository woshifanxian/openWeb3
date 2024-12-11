// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract C {
    uint256 public data = 30; // 公共状态变量
    uint256 internal iData = 10; // 内部状态变量

    function x() public returns (uint256) {
        data = 3; // 内部访问公共变量
        return data;
    }
}

contract Caller {
    C c = new C();

    function f() public view returns (uint256) {
        return c.data(); // 外部访问公共变量
    }
}

contract D is C {
    uint256 storedData;

    function y() public returns (uint256) {
        iData = 3; // 派生合约内部访问内部变量
        return iData;
    }

    function getResult() public view returns (uint256) {
        uint256 a = 1; // 局部变量
        uint256 b = 2;
        uint256 result = a + b;
        return storedData; // 访问状态变量
    }
}
