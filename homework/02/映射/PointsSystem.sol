// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract PointsSystem {
    mapping(address => uint) public points;
    address[] public users;

    function addUser(address _user) public {
        require(points[_user] == 0, "User already exists.");
        users.push(_user);
        points[_user] = 100; // 初始化用户积分为100
    }

    function transferPoints(address _to, uint _amount) public {
        require(points[msg.sender] >= _amount, "Insufficient points.");
        points[msg.sender] -= _amount;
        points[_to] += _amount;
    }

    function getAllUsers() public view returns (address[] memory) {
        return users;
    }
}