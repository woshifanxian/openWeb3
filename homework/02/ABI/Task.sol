// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Task {
    function execute(  address target, uint256 value, bytes memory data) public payable returns (bytes memory) {
        (bool success, bytes memory returnData) = target.call{value: value}(
            data
        );
        require(success, "execute: Transaction execution reverted.");
        return returnData;
    }
}
