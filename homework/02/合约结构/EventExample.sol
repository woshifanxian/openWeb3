// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract EventExample{
    event DataChanged(uint newValue);

    uint public data;

    function setData(uint _data)public{
        data = _data;
        emit DataChanged(_data);
    }
}