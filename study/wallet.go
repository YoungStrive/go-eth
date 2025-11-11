package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func main() {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//将私钥转换成字节数组
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//去掉前两位
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	//返回公钥
	publicKey := privateKey.Public()
	//将其转换为十六进制的过程与我们使用转化私钥的过程类似
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	//利用公钥，就可以轻松生成你经常看到的公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	/**
	*公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀。
	以下是使用 golang.org/x/crypto/sha3 的 Keccak256 函数手动完成的方法。
	*/
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
