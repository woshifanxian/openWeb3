// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
 import "@openzeppelin/contracts/utils/Address.sol";
import "./Address2.sol";

contract   MyToken {
    using Address2 for address;

    function send(address recipient, uint256 amount) external view {
        if (recipient.isContract()) {
            // do something
        }
    }

    function getAddress() public view returns(address){
       address addr = 0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990;
            address payable addr_pay = payable(0x8306300ffd616049FD7e4b0354a64Da835c1A81C);
            return addr_pay;
    }
}
