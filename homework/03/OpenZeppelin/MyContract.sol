// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "@openzeppelin/contracts/access/Ownable.sol";
contract MyContract is Ownable {
    mapping(address => bool) blacklist; // 记录账户是否在黑名单中

    constructor() Ownable(msg.sender) {
        if (msg.sender == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }

        _transferOwnership(msg.sender);
    }

    function _transferOwnership(address newOwner) internal override {
        super._transferOwnership(newOwner);
    }

    // 设置账户进出黑名单
    function setBlacklist(address account, bool flag) public onlyOwner {
        // 设置账户在黑名单中的状态为 flag
        blacklist[account] = flag;
    }

    // 检查账户是否在黑名单中
    function isBlaclisted(address account) public view returns (bool) {
        // 返回账户在黑名单中的状态
        return blacklist[account];
    }
}
