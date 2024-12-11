// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;
contract StructTest{

    struct Product{
        uint id;
        string name;
        uint price;
        uint store;
    }

    mapping(uint => Product) public products;
    uint public productCount;

    function createProduct(string memory _name, uint _price,uint _store) public {
        productCount = productCount++ ;
        products[productCount] = Product(productCount,_name,_price,_store);
    }

    function updateProduct(uint _id, string memory _name, uint _price,uint _store) public {
        Product storage product = products[_id];
        product.id = _id;
        product.name = _name;
        product.price = _price;
        product.store = _store;
    }

    function getProduct(uint _id) public view returns(string memory,uint price,uint store){
        Product storage product = products[_id];
        return(product.name,product.price,product.store);
    }
}