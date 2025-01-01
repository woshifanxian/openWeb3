package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func main() {
	// Step 1: 生成随机熵
	entropy, err := bip39.NewEntropy(256) // 256 位熵会生成 24 个助记词
	if err != nil {
		log.Fatalf("生成熵失败: %v", err)
	}

	// Step 2: 生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("生成助记词失败: %v", err)
	}
	fmt.Println("助记词:", mnemonic)

	// Step 3: 从助记词生成种子
	password := "" // 可选的密码
	seed := bip39.NewSeed(mnemonic, password)
	fmt.Printf("种子: %x\n", seed)

	// Step 4: 使用种子生成主私钥 (Master Key)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("生成主私钥失败: %v", err)
	}
	fmt.Println("主私钥:", masterKey)

	// Step 5: 派生子私钥 (Child Key)
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild) // 使用硬化路径
	if err != nil {
		log.Fatalf("派生子私钥失败: %v", err)
	}
	fmt.Println("子私钥:", childKey)

	// Step 6: 导出公钥
	fmt.Println("主公钥:", masterKey.PublicKey())
	fmt.Println("子公钥:", childKey.PublicKey())

	// Step 7: 创建派生路径 (如 BIP44)
	// m / purpose' / coin_type' / account' / change / address_index
	accountKey, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 0) // Account 0
	changeKey, _ := accountKey.NewChildKey(0)                            // External chain
	addressKey, _ := changeKey.NewChildKey(0)                            // Address 0

	fmt.Println("派生路径地址私钥:", addressKey)
	fmt.Println("派生路径地址公钥:", addressKey.PublicKey())

}
