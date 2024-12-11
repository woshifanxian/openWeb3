// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "hardhat/console.sol";
contract TryCatchExample {
    function tryCatchDemo(address _contractAddress) public {
        // 尝试调用外部合约的函数
        try ExternalContract(_contractAddress).someFunction() returns (
            bool result
        ) {
            // 处理成功
            console.log(1);
            
        } catch {
            // 处理失败
            console.log(2);
        }
    }
}

contract ExternalContract {
    function someFunction() public returns (bool) {
       return false;
    }
}
