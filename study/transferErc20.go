package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

// 代币转账
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

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0) // in wei (0 eth)
	//燃气价格 用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(10000000000))
	fmt.Println("gasPrice", gasPrice)
	//转给谁 转入的 地址
	toAddress := common.HexToAddress("0xf9626ac88c765de3961047f5706f7798756462fb")

	//这个是代币的地址
	tokenAddress := common.HexToAddress("0xE8FF1c356205c7d188F057BFF84889C62b4Cd59E")

	transferFnSignature := []byte("transfer(address,uint256)")

	//生成函数签名的 Keccak256 哈希
	hash := sha3.NewLegacyKeccak256()
	//写入方法签名
	hash.Write(transferFnSignature)
	//获取到方法ID
	methodID := hash.Sum(nil)[:4]

	//发送代币的地址左填充到 32 字节。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	///一个代币
	amount.SetString("1000000000000000000", 10)
	//代币量也需要左填充到 32 个字节。
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	//只需将方法 ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片。
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//燃气上限制将取决于交易数据的大小和智能合约必须执行的计算步骤
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	gasLimit := uint64(60000)
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
