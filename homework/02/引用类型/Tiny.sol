// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract Tiny {
    uint[] x; // x 的数据存储位置是 storage
    function f(uint[] memory memoryArray) public {
        x = memoryArray; // 将整个数组拷贝到 storage 中
        uint[] storage y = x; // 分配一个指针，y 的数据存储位置是 storage
        y[0] = 100; // 修改 y 的值，同时 x 的值也会改变
        delete x; // 清除 x，同时影响 y
    }
    function g(uint[] storage) internal pure {}
    function h(uint[] memory) public pure {}
}