// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract VisibilityExample {
    uint256 public data;
    uint256 internal iData = 10;

    function setData(uint256 a) internal {
        data = a;
    }

    function f(uint256 a) private pure returns (uint256) {
        return a + 1;
    }
}
