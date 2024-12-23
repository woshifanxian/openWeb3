// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract EtherStore {
    uint256 public withdrawalLimit = 1 ether;
    mapping(address => uint256) public lastWithdrawTime;
    mapping(address => uint256) public balances;
// 存储所有键（地址）的数组
    address[] public addressList;

        // 初始化互斥锁
    bool reEntrancyMutex = false;

    function depositFunds() public payable {
        balances[msg.sender] += msg.value;
        addressList.push(msg.sender);
    }

    function getAllBalances() public view returns (address[] memory, uint256[] memory) {
        uint256 length = addressList.length;
        uint256[] memory balanceValues = new uint256[](length);
        for (uint256 i = 0; i < length; i++) {
            balanceValues[i] = balances[addressList[i]];
        }
        return (addressList, balanceValues);
    }

    function withdrawFunds(uint256 _weiToWithdraw) public {
        require(balances[msg.sender] >= _weiToWithdraw, "Insufficient balance");
        // 限制提现金额
        require(_weiToWithdraw <= withdrawalLimit, "Exceeds withdrawal limit");
        // 限制允许提现的时间
        require(
            block.timestamp >= lastWithdrawTime[msg.sender] + 1 weeks,
            "Withdrawal not allowed yet"
        );

        //防止重入攻击
        balances[msg.sender] -= _weiToWithdraw;
        lastWithdrawTime[msg.sender] = block.timestamp;
        reEntrancyMutex = true;
        // 使用新的 call 语法，并传递空的 calldata
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Transfer failed");
        reEntrancyMutex = false;
    }

    function getSelfAndSenderBalance() public view returns(uint256,uint256,uint256){
        return (address(this).balance,balances[msg.sender],msg.sender.balance);
    }
}
