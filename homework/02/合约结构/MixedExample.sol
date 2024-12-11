// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract MixedExample {
    event Received(address sender, uint256 amount);
    event FallbackCalled(address sender, uint256 amount);

    // 当纯以太币转账时触发
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    // 当调用不存在的函数或附加了数据时触发
    fallback() external payable {
        emit FallbackCalled(msg.sender, msg.value);
    }
}
