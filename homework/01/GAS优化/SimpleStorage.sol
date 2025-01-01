// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private storedData;

    event DataStored(uint256 indexed data);

    constructor(uint256 _storedData)   {
        storedData = _storedData;
    }

    function set(uint256 data) public {
        storedData = data;
        emit DataStored(data);
    }

    function get() public view returns (uint256) {
        return storedData;
    }
}
