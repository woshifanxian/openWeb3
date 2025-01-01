// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

abstract contract erc20 {
    string public constant name = "aaa";
    string public constant symbol = "bbbb";
    uint8 public constant decimals = 18;

    // 将所有的函数标记为 virtual
    function totalSupply() public view virtual returns (uint);
    function balanceOf(address tokenOwner) public view virtual returns (uint balance);
    function allowance(address tokenOwner, address spender) public view virtual returns (uint remaining);
    function transfer(address to, uint tokens) public virtual returns (bool success);
    function approve(address spender, uint tokens) public virtual returns (bool success);
    function transferFrom(address from, address to, uint tokens) public virtual returns (bool success);

    // 定义事件
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
