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
	auth.Value = big.NewInt(0)      // 部署合约无需初始 ETH
	auth.GasLimit = uint64(6000000) // 设置 Gas 限制
	auth.GasPrice = gasPrice

	// 部署智能合约
	address, tx, _, err := solidity.DeployHelloToken(auth, client, big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
}
