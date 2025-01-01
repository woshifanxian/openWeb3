// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract StorageExample {
    address public user;  // 20 字节
    uint96 public balance; // 12 字节，与 `user` 共享一个存储槽

    function setValues(address _user, uint96 _balance) public {
        user = _user;
        balance = _balance;
    }
}
