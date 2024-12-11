// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract AssertRequireExample {
    address public owner;
    constructor() public {
        owner = msg.sender;
    }
    function transferOwnership(address newOwner) public {
        require(msg.sender == owner, "Only the owner can transfer ownership."); // 检查调用者是否为合约所有者
        owner = newOwner;
    }
    function checkBalance(uint a, uint b) public pure returns (uint) {
        uint result = a + b;
        assert(result >= a); // 检查溢出错误
        return result;
    }
}