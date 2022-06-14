package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url       = "https://kovan.infura.io/v3/b475e2c5cc2445569728e45c43a0b5af"
	infuraURL = "https://mainnet.infura.io/v3/b475e2c5cc2445569728e45c43a0b5af"
)

func main() {
	/*ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount("password")
	if err != nil {
		log.Fatal(err)
	}

	_, err = ks.NewAccount("password")
	if err != nil {
		log.Fatal(err)
	}*/

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	addr1 := common.HexToAddress("7f0cc8f62f89ad5a899192d3922eb55815649c36")
	addr2 := common.HexToAddress("96119df6a3f7bcaf1054cfae63e86eec436b1aba")

	bAddr1, err := client.BalanceAt(context.Background(), addr1, nil)
	if err != nil {
		log.Fatal(err)
	}

	bAddr2, err := client.BalanceAt(context.Background(), addr2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance Addr 1: ", bAddr1)
	fmt.Println("Balance Addr 2: ", bAddr2)
}
