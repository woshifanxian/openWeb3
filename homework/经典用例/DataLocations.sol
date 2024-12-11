// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
contract DataLocations{

    uint256[] public arr;

    struct MyStruct{
        uint256 foo;
    }

    mapping (uint256 => MyStruct) myStructs;

    mapping (uint256 => address) map;


    function f() public{
        _f(arr,map,myStructs[1]);
    }

    function _f(uint256[] storage _arr,mapping(uint256 => address) storage _map,MyStruct storage _myStruct) internal  {
        
    }

    function g(uint256[] memory _arr) public returns(uint256[] memory){

    }

    function h(uint256[] calldata _arr) public returns(uint256[] memory){
        
    }


}