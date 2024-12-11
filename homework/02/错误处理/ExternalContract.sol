// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract ExternalContract {
    function getValue() public pure returns (uint256) {
        return 42;
    }

    function willRevert() public pure {
        revert("This function always fails");
    }
}

contract TryCatchExample {
    ExternalContract externalContract;

    constructor() {
        externalContract = new ExternalContract();
    }

    function tryCatchTest() public returns (uint256, string memory) {
        try externalContract.getValue() returns (uint256 value) {
            return (value, "Success");
        } catch {
            return (0, "Failed");
        }
    }

    function tryCatchWithRevert() public returns (string memory) {
        try externalContract.willRevert() {
            return "This will not execute";
        } catch Error(string memory reason) {
            return reason; // 捕获错误信息
        } catch {
            return "Unknown error";
        }
    }
}
