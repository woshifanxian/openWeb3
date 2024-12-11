// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Hello {
    function sayHi() public {}

    // 可支付回退函数
    receive() external payable {}

    function isContract(address addr) public view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }

    function getAddress()public view returns(address){
        return address(this);
    }

    function getContractInfo() internal pure  returns(string memory,bytes memory,bytes memory){
        return (type(Hello).name,type(Hello).creationCode,type(Hello).runtimeCode);
    }
}
