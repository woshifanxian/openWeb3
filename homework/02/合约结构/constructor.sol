// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract ConstructorExample {
    uint256 public x;

    constructor(uint256 _x) public {
        x = _x;
    }
}
