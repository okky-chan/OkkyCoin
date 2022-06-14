package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	gen "github.com/okky-chan/OkkyCoin/gen"
)

func main() {
	b, err := ioutil.ReadFile("wallet/UTC--2022-06-12T14-54-48.670521200Z--7f0cc8f62f89ad5a899192d3922eb55815649c36")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://kovan.infura.io/v3/b475e2c5cc2445569728e45c43a0b5af")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, err := .DeployOkky(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("------------------------------")
}
