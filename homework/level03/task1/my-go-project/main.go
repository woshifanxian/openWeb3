package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"my-go-project/solidity"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 521386980 进入技术交流群 请先star，谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

const (
	keyStorePath = `your keyStorePath`
	fromKeyStore = `your from keyStore`
	toAddress    = "0x477257735cCEF9C7d42856649c7149865D04eDeb"
	chainId      = 7851254
	chtAddress   = "0xb9bD7405797CfFBc7e57309444b4af89c39cA92c" //合约地址
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/5907491931db4d6f939b3332ceb2ec2f")
	if err != nil {
		log.Fatalf("Failed to connect to sepolia Ethereum client: %v", err)
	}

	// 私钥（请使用您自己的私钥）
	privateKey, err := crypto.HexToECDSA("8ed72b66cec3121c88416e127457aff65556fb86b28978a93b949512b32adad3")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// 创建 Auth 事务
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Invalid public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}
	chainID := big.NewInt(11155111) //  测试网

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // 部署合约无需初始 ETH
	auth.GasLimit = uint64(300000) // 设置 Gas 限制
	auth.GasPrice = gasPrice

	// 部署智能合约
	address, tx, instance, err := solidity.DeploySimpleStorage(auth, client, big.NewInt(22))
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	_storeData, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	// 使用合约实例
	fmt.Printf("Instance: %v\n", instance)
	fmt.Printf("storeData: %v\n", _storeData)

	// 智能合约地址
	//address := common.HexToAddress("0xf64cD0ae1722781F260DD77566Ee795C8dfe41aa")
	//
	//// 创建实例
	//instance, err := solidity.NewSolidity(address, client)
	//if err != nil {
	//	log.Fatalf("Failed to instantiate contract: %v", err)
	//}
	//// 调用 set 方法
	//data, err := instance.SetMessage(&bind.TransactOpts{}, "aaa")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err != nil {
	//	log.Fatalf("Failed to call contract method: %v", err)
	//}
	//fmt.Printf("Transaction: %d\n", data)
	//
	//// 调用 set 方法
	//data2, err := instance.Message(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err != nil {
	//	log.Fatalf("Failed to call contract method: %v", err)
	//}
	//log.Fatalf("Stored data:" + data2)
}
