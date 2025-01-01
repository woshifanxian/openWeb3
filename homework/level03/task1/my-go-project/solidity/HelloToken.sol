// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 导入 OpenZeppelin 的 ERC20 标准合约
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract HelloToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("HelloToken", "HTN") {
        // 初始化代币总量，分配给合约创建者
        _mint(msg.sender, initialSupply * (10 ** decimals()));
    }
}
