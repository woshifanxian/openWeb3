// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract ArrayOperations {
    uint256[] public dynamicArray;

    function addElement(uint256 _element) public {
        dynamicArray.push(_element); // 向数组添加元素
    }

    function removeLastElement() public {
        dynamicArray.pop(); // 删除数组最后一个元素
    }

    function getLength() public view returns (uint256) {
        return dynamicArray.length; // 获取数组长度
    }
}
