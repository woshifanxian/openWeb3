// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Registry {
    mapping(uint256 => address) public documents;

    function register(uint256 hash) public {
        documents[hash] = msg.sender;
    }
}

contract DocumentRegistry {
    event Registered(uint256 hash, address sender);

    function register(uint256 hash) public {
        emit Registered(hash, msg.sender);
    }
}
