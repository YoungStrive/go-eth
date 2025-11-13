package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	//获取到客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/pNX-SE87t8JfMGoicYCwy")
	if err != nil {
		log.Fatal(err)
	}
	var privateStr = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

	//加载自己的私钥
	privateKey, err := crypto.HexToECDSA(privateStr)
	if err != nil {
		log.Fatal(err)
	}
	//根据私钥获取到公钥
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//根据公钥 获取到 帐户的公共地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取到随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//单位wei 0.01eth
	value := big.NewInt(1000000000000000)
	//ETH 转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(21000)

	//燃气价格 用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//转给谁 转入的 地址
	toAddress := common.HexToAddress("0xf9626ac88c765de3961047f5706f7798756462fb")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//是使用发件人的私钥对事务进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//将已签名的事务广播到整个网络。
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
