// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract DataHandling{
    uint[]  public data;
    function updateData(uint[] memory newData)public {
        data = newData;
    }

    function getData()public view  returns(uint[] memory){
        return data;
    }

    function modifyStorageData(uint index, uint value) public{
        require(data.length >= index, unicode"数组越界");
        data[index] = value;
    }

    function modifyMemoryData(uint[] memory memData) public pure returns(uint[] memory){
         require(memData.length > 0,unicode"数组没有元素");
         memData[0] = 1;
         return memData;
    }
}