package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://sepolia.infura.io/v3/a2b0e9a36e9a481b8a5356585850b404"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to creat ether client: %v", err)
	}
	defer client.Close()

	ac1, err := os.ReadFile("./wallet/UTC--2024-03-06T18-26-46.040477236Z--4a02341be8e888617faedd979d25adc80061cece")
	if err != nil {
		log.Fatalf("Error while reading wallet file: %v", err)
	}
	key, err := keystore.DecryptKey(ac1, "account1")
	if err != nil {
		log.Fatalf("Error while getting the balance: %v", err)
	}
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce , err := client.PendingCodeAt(context.Background(),add)
	if err != nil {
		log.Fatalf("Error while getting pending nonce|: %v", err)
	}

	gPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error while getting pending nonce|: %v", err)
	}
	cID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error while getting pending nonce|: %v", err)
	}

	



}
