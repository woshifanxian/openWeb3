// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.17;
contract Owner{

    Identity private owner;
    State private state;

    struct Identity{
        address addr;
        string name;
    }

    enum State{
        HasOwner,
        NoOwner
    }

    event OwnerSet(address indexed oldOwnerAddr,address indexed newOwnerAddr);

    event OwnerRemoved(address indexed oldOwner);

    modifier isOwner{
        require(msg.sender == owner.addr,"Caller is not owner");
        _;
    }

    constructor (string memory name){
        owner.addr = msg.sender;
        owner.name = name;
        state = State.HasOwner;
        emit OwnerSet(address(0), owner.addr);
    }

    function changeOwner(address addr,string calldata name) public isOwner{
        owner.addr = msg.sender;
        owner.name = name;
        emit OwnerSet(owner.addr, addr);
    }

    function removeOwner() public isOwner{
        emit OwnerRemoved(owner.addr);
        delete owner;
        state = State.NoOwner;
    }

     // 普通函数
    function getOwner() external view returns (address, string memory) {
        return (owner.addr, owner.name);
    }

    // 普通函数
    function getState() external view returns (State) {
        return state;
    }
}