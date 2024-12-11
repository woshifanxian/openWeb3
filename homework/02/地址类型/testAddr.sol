// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "hardhat/console.sol";
contract testAddr {
    // 如果合约的余额大于等于 10，而 x 小于 10,则给 x 转 10 wei
    function testTrasfer(address payable x) public {
        address myAddress = address(this);
        console.log(myAddress.balance);
        if (x.balance < 10 && myAddress.balance >= 10) {
            x.transfer(10);
        }
    }
}
