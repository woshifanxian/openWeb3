// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract MultiStageTimeLock {
    uint256 public stage1UnlockTime;
    uint256 public stage2UnlockTime;
    address public owner;

    constructor(uint256 _stage1Lock, uint256 _stage2Lock) {
        owner = msg.sender;
        stage1UnlockTime = block.timestamp + _stage1Lock * 1 days; // 第一阶段锁定天数
        stage2UnlockTime = block.timestamp + _stage2Lock * 1 days; // 第二阶段锁定天数
    }

    function withdrawStage1() public {
        require(
            block.timestamp >= stage1UnlockTime,
            "Stage 1 funds are locked."
        );
        require(msg.sender == owner, "Only owner can withdraw.");
        // 提取第一阶段资金操作
    }

    function withdrawStage2() public {
        require(
            block.timestamp >= stage2UnlockTime,
            "Stage 2 funds are locked."
        );
        require(msg.sender == owner, "Only owner can withdraw.");
        // 提取第二阶段资金操作
    }
}
