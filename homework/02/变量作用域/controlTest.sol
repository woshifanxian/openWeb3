// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract controlTest {
    function testWhile() public pure returns (uint256) {
        uint256 i = 0;
        uint256 sumOfOdd = 0;
        while (true) {
            i++;
            if (i % 2 == 0) {
                continue; // 跳过偶数
            }
            if (i > 10) {
                break; // 终止循环
            }
            sumOfOdd += i;
        }
        return sumOfOdd; // 返回奇数和
    }

    function testFor() public pure returns (uint256, uint256) {
        uint256 sumOfOdd = 0;
        uint256 sumofEven = 0;
        for (uint256 i = 0; i < 10; i++) {
            if (i % 2 == 0) {
                sumofEven += i; // 累加偶数
            } else {
                sumOfOdd += i; // 累加奇数
            }
        }
        return (sumOfOdd, sumofEven); // 返回奇数和与偶数和
    }
}
