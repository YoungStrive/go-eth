package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

// 查询ETH余额
func main() {
	//获取到客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/pNX-SE87t8JfMGoicYCwy")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x48f99e89590e7c2e9f9198ffe6257c6be41e7a1c")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	//blockNumber := big.NewInt(5532993)
	//balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balanceAt)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

}
