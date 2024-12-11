// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "./IToken.sol";
// 实现接口的合约
contract SimpleToken is IToken {
    mapping(address => uint256) public balances;

    constructor() {
        balances[msg.sender] = 1000; // 初始化代币余额
    }

    // 实现接口中的 transfer 函数
    function transfer(address recipient, uint256 amount) public override {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        balances[recipient] += amount;
    }
}