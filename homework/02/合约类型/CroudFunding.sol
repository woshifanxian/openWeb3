// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Crowdfunding {
    // 合约所有者
    address public owner;

    // 捐款记录
    mapping(address => uint256) public contributions;

    // 事件：当有人捐款时触发
    event DonationReceived(address indexed donor, uint256 amount);

    // 事件：当资金被提取时触发
    event FundsWithdrawn(address indexed owner, uint256 amount);

    // 构造函数：设置合约所有者
    constructor() {
        owner = msg.sender;
    }

    // 捐款函数：带有payable关键字，允许接收以太币
    function donate() external payable {
        // 捐款的金额必须大于0，否则会抛出异常
        require(msg.value > 0, "Donation amount must be greater than zero");
        contributions[msg.sender] += msg.value;
        emit DonationReceived(msg.sender, msg.value);
    }

    // 提取资金函数：仅合约所有者可以调用
    function withdraw(uint256 _amount) external onlyOwner {
        require(address(this).balance >= _amount, "Insufficient contract balance");
        payable(owner).transfer(_amount);
        emit FundsWithdrawn(owner, _amount);
    }

    // 获取合约当前余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // 修饰符：限制只有合约所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the contract owner can call this function");
        _;
    }
}

