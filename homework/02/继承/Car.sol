// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract Car{
    uint internal speed;

    function drive() virtual public{}
}
contract ElectricCar is Car{
     uint internal batteryLevel;
     function drive() public override {}
}
