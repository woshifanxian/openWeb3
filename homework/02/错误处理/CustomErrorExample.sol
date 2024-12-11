// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract CustomErrorExample {
    error Unauthorized(address caller);  // 自定义错误
    address public owner;
    constructor() {
        owner = msg.sender;
    }
    function restrictedFunction() public {
        if (msg.sender != owner) {
            revert Unauthorized(msg.sender);  // 使用自定义错误
        }
    }
}