package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	// 读取 Keystore 文件
	fileContent, err := os.ReadFile("./tmp/UTC--2024-12-27T08-03-42.982271500Z--c8fff46774f101c9a0a092b8b43d073b8e129143")
	if err != nil {
		log.Fatalf("Failed to read keystore file: %v", err)
	}

	// 解密 Keystore 文件，提取私钥
	key, err := keystore.DecryptKey(fileContent, password)
	if err != nil {
		log.Fatalf("Failed to decrypt keystore: %v", err)
	}

	// 打印私钥和地址
	fmt.Printf("Private Key: %x\n", key.PrivateKey.D.Bytes())
	fmt.Printf("Account Address: %s\n", key.Address.Hex())
}

func importKs() {
	file := "./tmp/UTC--2024-12-27T07-29-54.511301500Z--72382322e69489af09f73f9d7a8ff6c6b7dcaacf"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	createKs()
	//importKs()
}
