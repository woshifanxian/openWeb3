// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "hardhat/console.sol";
contract C1 {
 function tryCatchDemo(address _contractAddress,uint n) public pure returns(uint result2) {
        // 尝试调用外部合约的函数
        try C2(_contractAddress).sum(n) returns (uint result2){
            // 处理成功
             console.log(result2);
        } catch {
            // 处理失败
              console.log(result2);
        }
    }
}

contract C2{
  function sum(uint n) public pure returns (uint256) {
        uint256 sum = 0;
        for (uint256 i = 1; i <= n; i++) {
            if (i % 2 != 0) {
                sum += i;
            }
        }
        return sum;
    }

    function getAddress()public view  returns(address){
        return address(this);
    }
}
