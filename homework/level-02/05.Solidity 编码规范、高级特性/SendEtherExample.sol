// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SendEtherExample {
    // 事件：记录存款操作
    event Deposit(address indexed sender, uint256 amount);

    // 存款函数：接收 ETH 并触发事件
    function deposit() public payable {
        require(msg.value > 0, "Must send some Ether");
        emit Deposit(msg.sender, msg.value);
    }

    // 获取合约的当前余额
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function getAddress() public view returns (address) {
        return msg.sender;
    }
    
    // 提款函数：发送合约内的所有 ETH 给调用者
    function withdraw() public {
        payable(msg.sender).transfer(address(this).balance);
    }

    // 接收 ETH 的特殊函数
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    function payToContract() public payable {
        require(msg.value > 0, "Must send some Ether");
        emit Deposit(msg.sender, msg.value);
    }
}
