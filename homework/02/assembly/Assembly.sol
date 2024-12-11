// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Assembly {
    function addAssembly(uint256 x, uint256 y) public pure returns (uint256) {
        assembly {
            let result := add(x, y) // x+y
            mstore(0x0, result) // 在内存中保存结果
            return(0x0, 32) // 从内存中返回 32 字节
        }
    }

    function addSolidity(uint256 x, uint256 y) public pure returns (uint256) {
        return x + y;
    }

    function localvar() public pure {
        uint256 b = 5;
        assembly {
            let x := add(2, 3)
            let y := mul(x, b) // 使用了外面的 b
            let z := add(x, y) // 访问了内部定义的 x,y
        }
    }

    function forloop(uint256 n, uint256 value) public pure returns (uint256) {
        assembly {
            for {
                let i := 0
            } lt(i, n) {
                i := add(i, 1)
            } {
                value := add(2, value)
            }
            mstore(0x0, value)
            return(0x0, 32)
        }
    }
}
