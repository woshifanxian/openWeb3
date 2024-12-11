// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract DailyAction {
    uint256 public lastActionTime;
    address public owner;

    constructor() {
        owner = msg.sender;
        lastActionTime = block.timestamp - 1 days; // 初始化为可立即执行
    }

    function performAction() public {
        require(msg.sender == owner, "Only owner can perform this action.");
        require(
            block.timestamp >= lastActionTime + 1 days,
            "You can only perform this action once per day."
        );
        lastActionTime = block.timestamp; // 更新最后一次执行时间
        // 执行具体操作
    }
}
