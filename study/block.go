package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e7cbd13b843d4556bb9947cbf12cdbda")
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	fmt.Println(header.Number.String())

}
