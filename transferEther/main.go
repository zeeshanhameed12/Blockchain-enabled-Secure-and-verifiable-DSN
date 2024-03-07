package main

import (
	"context"
	"fmt"

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
	ac2, err := os.ReadFile("./wallet/UTC--2024-03-06T18-26-47.214027485Z--fc0d6b1a3bed7f18af443ad3768c710e888d22d0")
	if err != nil {
		log.Fatalf("Error while reading wallet file: %v", err)
	}

	///Account 1
	key, err := keystore.DecryptKey(ac1, "account1")
	if err != nil {
		log.Fatalf("Error while getting the balance: %v", err)
	}
	key2, err := keystore.DecryptKey(ac2, "account2")

	if err != nil {
		log.Fatalf(" Error while decrypting the encrypted wallet: %v", err)
	}

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	add2 := crypto.PubkeyToAddress(key2.PrivateKey.PublicKey)

	balance, err := client.BalanceAt(context.Background(), add, nil)
	if err != nil {
		log.Fatalf("Error while getting the balance: %v", err)
	}
	balance2, err := client.BalanceAt(context.Background(), add2, nil)
	if err != nil {
		log.Fatalf("Error while getting the balance: %v", err)
	}

	// //nonce , err := client.PendingCodeAt(context.Background(),add)
	// if err != nil {
	// 	log.Fatalf("Error while getting pending nonce|: %v", err)
	// }
	// nonce, err := client.PendingNonceAt(context.Background(), add)

	// amount := big.NewInt(100000000000000000)
	// gPrice, err := client.SuggestGasPrice(context.Background())
	// tX := types.NewTransaction(nonce, add2, amount, 21000, gPrice, nil)
	// cID, err := client.NetworkID(context.Background())

	// t, err := types.SignTx(tX, types.NewEIP155Signer(cID), key.PrivateKey)

	// client.SendTransaction(context.Background(),t)

	// fmt.Println("tx sent", t.Hash().Hex())

	fmt.Println("Balance in Account 1: ", float32(balance.Int64())/float32(1000000000000000000))
	fmt.Println("Balance in Account 2: ", float32(balance2.Int64())/float32(1000000000000000000))

}
