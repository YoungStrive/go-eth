package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e7cbd13b843d4556bb9947cbf12cdbda")
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9543800)
	//根据区块编码获取到交易的信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	for _, tx := range block.Transactions() {
		//交易哈希值
		fmt.Println(tx.Hash())
		//交易的金额
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		} else {
			log.Fatal(err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		break
	}

	txHash := common.HexToHash("0x0021da846b0cff6926893c3a0581f0bebc2739f78951c80d07ab6ec31b668162")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())

}
