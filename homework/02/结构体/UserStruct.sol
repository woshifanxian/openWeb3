// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract UserStruct {
    struct User {
        string name;
        uint256 age;
        mapping(string => uint256) score;
    }

    mapping(address => User) public users;

    function createUser(string memory _name, uint256 _age) public {
        User storage user = users[msg.sender];
        user.name = _name;
        user.age = _age;
    }

    function updateUserScore(string memory _subject, uint256 _score) public {
        users[msg.sender].score[_subject] = _score;
    }

    function getUserScore(string memory _subject) public view returns (uint256)
    {
        return users[msg.sender].score[_subject];
    }
}
