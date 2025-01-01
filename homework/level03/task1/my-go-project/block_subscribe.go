package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/5907491931db4d6f939b3332ceb2ec2f")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("区块hash是:" + header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("block.Hash()是:" + block.Hash().Hex())                                // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println("block.Number()是:" + strconv.FormatUint(block.Number().Uint64(), 10)) // 3477413
			fmt.Println(int64(block.Time()))                                                  // 1529525947
			fmt.Println(block.Nonce())
			fmt.Println(block.Transactions())      // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
