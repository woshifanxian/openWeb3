// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract ReceiveExample {
    event Received(address sender, uint256 amount);

    //仅用于接收以太币
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }
}
