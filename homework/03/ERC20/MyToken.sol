// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

import "homework/library/openzeppelin-contracts-5.1.0/openzeppelin-contracts-5.1.0/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20("aaa", "bb") {
    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    constructor() public  {}

    function totalSupply() public view override  returns (uint256) {
        return address(this).balance;
    }

    receive() external payable {
        deposit();
    }

    function deposit() public payable {
        _mint(msg.sender, msg.value);
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 wad) public {
        _burn(msg.sender, wad);
        payable (msg.sender).transfer(wad);
        emit Withdrawal(msg.sender, wad);
    }

    function getBytes4()public pure returns (bytes4){
        return bytes4(keccak256('supportsInterface(bytes4)'));
    }
}
