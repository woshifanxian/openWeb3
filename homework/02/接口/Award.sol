// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "./IToken.sol";
contract Award {
    IToken immutable token;
    // 构造函数中传入 SimpleToken 合约的地址
    constructor(IToken _token) {
        token = _token;
    }
    // 调用 SimpleToken 合约的 transfer 函数来发送奖励
    function sendBonus(address user) public {
        token.transfer(user, 100);  // 向用户发送100个代币作为奖励
    }
}