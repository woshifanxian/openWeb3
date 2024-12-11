// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
import "homework/library/IterableMapping.sol";
contract UserBalanceManagerment{

    IterableMapping.Map private balances;

    function addBalances(uint256 newBalance) public{
        IterableMapping.set(balances,msg.sender,newBalance);
    }

    function queryBalances() public view returns(uint256){
        return IterableMapping.get(balances,msg.sender);
    }

    function queryUserBalances(address userAddress ) public view returns(uint256){
        return IterableMapping.get(balances,userAddress);
    }

    function addUserBalances(address userAddress,uint256 newBalance) public{
         IterableMapping.set(balances,userAddress,newBalance);
    }

    function sortUserBalances() public{
        if (IterableMapping.size(balances)<=0) {
            return;
        }
         /**
         * 冒泡排序
         * */
        require(IterableMapping.size(balances) > 0,"balances need > 0");
        uint size =  IterableMapping.size(balances);
        for (uint i = 0; i < size-1; i++) {
            for (uint j = 0; j < size-1-i; j++) {
                if(IterableMapping.getKeyAtIndex(balances,j)>IterableMapping.getKeyAtIndex(balances,j+1)){
                    uint256 temp= IterableMapping.getValueAtIndex(balances,j);
                    IterableMapping.set(balances,j,IterableMapping.getValueAtIndex(balances,j+1));
                    IterableMapping.set(balances,j+1,temp);
                }
            }
        }
    }
}