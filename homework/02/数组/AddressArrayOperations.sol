// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract AddressArrayOperations {
    uint256[] public dynamicAddressData;


    function getData() public pure returns(uint[3] memory){
        uint[3] memory x = [uint(1000), 1001, 1002]; 
        return x;
    }


    function add(uint256[] memory data) public pure returns (uint256) {
        uint256 x;
        if (data.length > 0) {
            for (uint256 i = 0; i < data.length; i++) {
                x += data[i];
            }
        }
        return x;
    }

    function del(uint256[] memory data) public pure returns (uint256[] memory) {
        uint256[] memory x;
        if (data.length > 0) {
            x = new uint256[](data.length - 1);
            delete data[0];
            x = data;
        }
        return x;
    }
}
