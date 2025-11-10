package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

// 收据
func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e7cbd13b843d4556bb9947cbf12cdbda")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9543800)
	blockHash := common.HexToHash("0x1a1eaebdeb38584b158780e18972aa2b765930eb43b12606d2c92ce1cad7a766")
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptByHash[0])
	fmt.Println(receiptsByNum[0])
}
