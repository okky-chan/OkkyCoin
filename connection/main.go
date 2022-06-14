package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/b475e2c5cc2445569728e45c43a0b5af"
var ganacheURL = "http://localhost:7545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a Ether client: %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println("The block number: ", block.Number())

	addr := "0xa462127735352b1f03da8ab92a87803d05cc6a7b"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get balance: %v", err)
	}
	fmt.Println("The balance: ", balance)

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)
	// 1 Eth = 10^18
	ethBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethBalance)
}
