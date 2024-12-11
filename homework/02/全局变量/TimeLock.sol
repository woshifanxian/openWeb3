// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract TimeLock {
    uint256 public unlockTime;
    address public owner;

    constructor(uint256 _lockTime) {
        owner = msg.sender;
        unlockTime = block.timestamp + _lockTime * 1 days; // 锁定指定天数
    }

    function withdraw() public {
        require(block.timestamp >= unlockTime, "Funds are locked.");
        require(msg.sender == owner, "Only owner can withdraw."); // 执行提款操作
    }
}
