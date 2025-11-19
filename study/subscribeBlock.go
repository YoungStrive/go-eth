package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 订阅区块信息
func main() {
	//订阅区块需要 websocket RPC URL。
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/pNX-SE87t8JfMGoicYCwy")
	if err != nil {
		log.Fatal(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		/**
		 * select语句是用于多通道（channel）操作的控制结构，其核心功能是‌随机选择一个可运行的case执行‌
		 * ， 若无可用case则阻塞
		 *  随机性：select会随机选择一个就绪的case执行，而非顺序匹配
		 * 阻塞机制：若所有case均未就绪，程序会阻塞直至有case可执行
		 * 支持default：可添加default分支处理无数据可接收的情况
		 */
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}
