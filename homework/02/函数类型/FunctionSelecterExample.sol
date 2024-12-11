// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract FunctionSelecterExample {
    bytes4 public storedSelector;

    function square(uint256 x) public pure returns (uint256) {
        return x * x;
    }

    function double(uint256 x) public pure returns (uint256) {
        return x * 2;
    }

    function storeSelector(bytes4 selector) public {
        storedSelector = selector;
    }

    function getDoubleSelector() external pure returns (bytes4) {
        return this.double.selector; // 返回函数double的选择器
    }

    function getSquareSelector() external pure returns (bytes4) {
        return this.square.selector; // 返回函数a的选择器
    }

    function executeStoredFunction(uint256 x) public returns (uint256) {
         require(storedSelector != bytes4(0), "Selector not set");
        (bool success, bytes memory data) = address(this).call(
            abi.encodeWithSelector(storedSelector, x)
        );
        require(success, "Call Failed");
        return abi.decode(data, (uint256));
    }

    function executeFunction(bytes4 selector, uint256 x) public returns (uint256)
    {
        require(selector != bytes4(0), "Selector not set");
        (bool success, bytes memory data) = address(this).call(
            abi.encodeWithSelector(selector, x)
        );
        require(success, "Call Failed");
        return abi.decode(data, (uint256));
    }
}
