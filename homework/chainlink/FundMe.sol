// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;
contract FundMe{

    uint256 public minumumUsd = 1 * 1e18;

    address[] public funders;

    mapping(address=>uint256) public addressToAmountFund;

    address public owner;

    constructor(){
        owner = msg.sender;
    }

    function fund() public payable {
        require(msg.value >= minumumUsd,"didnt enough fund");
        funders.push(msg.sender);
        addressToAmountFund[msg.sender] += msg.value;
    }

    function withdraw() public onlyOwner{
        for(uint funderIndex = 0;0<funders.length;funderIndex++){
            address funder = funders[funderIndex];
            addressToAmountFund[funder] = 0;
        }

        funders = new address[](0);
        (bool success,) = payable (msg.sender).call{value:address(this).balance}("");
        require(success,"call failed");

    }

    modifier onlyOwner{
        require(owner ==  msg.sender,"not owner");
        _;
    }

    function getBalance()public view returns(uint256){
        return address(this).balance;
    }
}