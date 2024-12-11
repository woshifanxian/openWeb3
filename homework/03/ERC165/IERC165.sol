// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
interface IERC165 {
    // @param interfaceID 参数：接口 ID
    function supportsInterface(bytes4 interfaceID) external view returns(bool);
}
