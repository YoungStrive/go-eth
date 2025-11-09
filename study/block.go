package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e7cbd13b843d4556bb9947cbf12cdbda")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9566002)
	block, err := client.HeaderByNumber(context.Background(), blockNumber)
	//获取区块头编码 区块号
	fmt.Println(block.Number.Uint64())
	//获取区块的时间
	fmt.Println(block.Time)
}
