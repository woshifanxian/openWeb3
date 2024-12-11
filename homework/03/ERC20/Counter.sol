// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Counter {
    uint8 public count = 0;

    function increment() external returns (uint8) {
        unchecked {
            count--;
        }
        return count;
    }
}
