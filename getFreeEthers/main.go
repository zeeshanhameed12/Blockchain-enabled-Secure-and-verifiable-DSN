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
	//key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := key.NewAccount("account1")
	// if err != nil {
	// 	log.Fatalf("Error while creating new encrypted wallet: %v", err)
	// }
	// _, err = key.NewAccount("account2")
	// if err != nil {
	// 	log.Fatalf("Error while creating new encrypted wallet: %v", err)
	// }
	ac1, err := os.ReadFile("./wallet/UTC--2024-03-06T18-26-46.040477236Z--4a02341be8e888617faedd979d25adc80061cece")
	if err != nil {
		log.Fatalf("Error while reading wallet file: %v", err)
	}
	// ac2, err := os.ReadFile("./wallet/UTC--2024-03-06T18-26-47.214027485Z--fc0d6b1a3bed7f18af443ad3768c710e888d22d0")
	// if err != nil {
	// 	log.Fatalf("Error while reading wallet file: %v", err)
	// }
	
	///Account 1
	key, err := keystore.DecryptKey(ac1, "account1")
	if err != nil {
		log.Fatalf(" Error while decrypting the encrypted wallet: %v", err)
	}
	// kData := crypto.FromECDSA(key.PrivateKey)
	// private_key := hexutil.Encode(kData)

	// fmt.Println("Private key of Account 1: ", private_key)

	// pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	// pub_key := hexutil.Encode(pubData)

	// fmt.Println("Public key Account 1: ", pub_key)
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println("Public address Account 1: ", add)

	balance, err := client.BalanceAt(context.Background(), add, nil)
	if err != nil {
		log.Fatalf("Error while getting the balance: %v", err)
	}
	fmt.Println("Balance in Account 1: ", float32(balance.Int64())/float32(1000000000000000000))












	///Account 2
	// key2, err := keystore.DecryptKey(ac2, "account2")
	// if err != nil {
	// 	log.Fatalf(" Error while decrypting the encrypted wallet: %v", err)
	// }
	// kDataAc2 := crypto.FromECDSA(key2.PrivateKey)
	// private_keyAc2 := hexutil.Encode(kDataAc2)

	// fmt.Println("Private key Account 2: ", private_keyAc2)

	// pubDataAc2 := crypto.FromECDSAPub(&key2.PrivateKey.PublicKey)
	// pub_keyAc2 := hexutil.Encode(pubDataAc2)

	// fmt.Println("Public key Account 2: ", pub_keyAc2)
	// addAc2 := crypto.PubkeyToAddress(key2.PrivateKey.PublicKey)
	// fmt.Println("Public address Account 2: ", addAc2)
	// balanceAc2, err := client.BalanceAt(context.Background(), addAc2, nil)
	// if err != nil {
	// 	log.Fatalf("Error while getting the balance: %v", err)
	// }
	// fmt.Println("Balance in Account 2: ", balanceAc2)

}
