// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "./SafeMath.sol";
contract TestLib {
    using SafeMath for uint;
    function add(uint x, uint y) public pure returns (uint) {
        return x.add(y);  // 直接在 uint 类型的变量上调用库函数
    }
}