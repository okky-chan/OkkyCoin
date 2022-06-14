package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	/*key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Address)*/
	password := "password"

	b, err := ioutil.ReadFile("./wallet/UTC--2022-06-12T14-22-03.014479000Z--c466513b2b5f9c7f8ff1d32647095d45c680647e")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv Key: ", hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public Key: ", hexutil.Encode(pubData))

	fmt.Println("Address: ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
