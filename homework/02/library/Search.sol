// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
library Search {
    function indexOf(uint[] storage self, uint value) public view returns (uint) {
        for (uint i = 0; i < self.length; i++) {
            if (self[i] == value) return i;
        }
        return uint(1);  // 如果找不到返回 -1
    }
}
contract C {
    using Search for uint[];  // 将库函数绑定到 uint[] 类型
    uint[] data;
    function append(uint value) public {
        data.push(value);
    }
    function replace(uint _old, uint _new) public {
        uint index = data.indexOf(_old);  // 调用库函数
        if (index == uint(1)) {
            data.push(_new);
        } else {
            data[index] = _new;
        }
    }
}