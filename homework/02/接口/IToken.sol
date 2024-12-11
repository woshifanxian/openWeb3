// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

// 定义一个接口 IToken
interface IToken {
    function transfer(address recipient, uint256 amount) external;
}

