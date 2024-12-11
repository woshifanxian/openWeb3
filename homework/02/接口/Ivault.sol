// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "./IERC20.sol";

interface Ivault {
    function deposit(uint256 money) external;

    function withdraw(uint256 money) external;
}

contract MyToken {
     mapping(address => uint256) public balances;
    constructor() {
        balances[msg.sender] = 1000;  // 初始化代币余额
    }
    // 实现接口中的 transfer 函数
    function transfer(address recipient, uint256 amount) public {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        balances[recipient] += amount;
    }
}

contract Reward is Ivault {
    MyToken immutable token;

    constructor(MyToken _token) {
        token = _token;
    }

    function deposit(uint256 money) public {
        token.transfer(address(this), money); // 向用户发送100个代币作为奖励
    }

    function withdraw(uint256 money) public {
        token.transfer(address(this), money); // 向用户发送100个代币作为奖励
    }
}

contract erc is IERC20 {
    function totalSupply() external pure returns (uint256) {
        return 0;
    }

    function balanceOf(address account) external pure returns (uint256) {
        return 0;
    }

    function transfer(address recipient, uint256 amount)
        external
        pure
        returns (bool)
    {
        return false;
    }

    function allowance(address owner, address spender)
        external
        pure
        returns (uint256)
    {
        return 0;
    }

    function approve(address spender, uint256 amount)
        external
        pure
        returns (bool)
    {
        return false;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external pure returns (bool) {
        return false;
    }
}
