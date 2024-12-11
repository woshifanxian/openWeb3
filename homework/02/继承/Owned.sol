// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract Owned {
    address payable public owner;


    function setOwner(address _owner) public virtual {
        owner = payable(_owner);
    }
}

abstract  contract Owned2 {
    address payable public owner;


    function setOwner(address _owner) public virtual {
        owner = payable(_owner);
    }

    function setOwner2(address _owner) public virtual {
        owner = payable(_owner);
    }

    function setOwner3(address _owner) public virtual{}
}

contract Mortal is Owned2 {
    event SetOwner(address indexed owner);


    // function setOwner3(address _owner) public override {
    //     super.setOwner(_owner); // 调用父合约的 setOwner
    //     emit SetOwner(_owner);
    // }
}