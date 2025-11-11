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
	//根据区块hash获取收据
	blockHash := common.HexToHash("0x1a1eaebdeb38584b158780e18972aa2b765930eb43b12606d2c92ce1cad7a766")
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	//根据区块编号获取收据
	blockNumber := big.NewInt(9543800)
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&receiptByHash[0] == &receiptsByNum[0])

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}
	//===================================================================================
	txHash := common.HexToHash("0x0021da846b0cff6926893c3a0581f0bebc2739f78951c80d07ab6ec31b668162")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	//收据的交易哈希
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())
}
