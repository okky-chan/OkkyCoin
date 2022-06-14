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
	/*nonce, err := client.PendingNonceAt(context.Background(), addr1)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(100000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, addr2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("wallet/UTC--2022-06-12T14-54-48.670521200Z--7f0cc8f62f89ad5a899192d3922eb55815649c36")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction : %s", tx.Hash().Hex())*/
}
