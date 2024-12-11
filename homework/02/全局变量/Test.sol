// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract Test{

    uint256 public unlockWeekTime;
    uint256 public unlockMonthTime;
    address public owner;

    constructor(uint256 _unlockWeekTime,uint256 _unlockMonthTime) {
        owner = msg.sender;
        unlockWeekTime = block.timestamp + _unlockWeekTime * 1 weeks; // 锁定1周
        unlockMonthTime = block.timestamp + _unlockMonthTime * 4 weeks; // 锁定一月
    }


    function getInfo() public view returns(uint,uint,address,uint){
        return (block.number,block.timestamp,block.coinbase,tx.gasprice);
    }

    function encode() public pure returns  (bytes memory){
       return  abi.encode(uint(1), address(0x123));
    }

    function decode(bytes memory encodeData) public pure returns  (uint,address){
       return  abi.decode(encodeData, (uint, address));
    }

    function getReward() public view {
        require(block.timestamp >= unlockWeekTime, "Funds are locked.");
        require(block.timestamp >= unlockMonthTime, "Funds are locked.");
        require(msg.sender == owner, "Only owner can withdraw."); // 执行提款操作
    }
}