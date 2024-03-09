package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"main.go/todo"
	//"Blockchain-enabled-Secure-and-verifiable-DSN/todo"
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

	// nonce, err := client.PendingNonceAt(context.Background(), add)
	// if err != nil {
	// 	log.Fatalf("Error while getting pending nonce|: %v", err)
	// }

	// gPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error while getting gass price|: %v", err)
	// }
	// cID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error while getting pCID|: %v", err)
	// }

	// txOption, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, cID)
	// if err != nil {
	// 	log.Fatalf("Error while making new transaction|: %v", err)
	// }
	// //fmt.Println(nonce)

	// txOption.GasPrice = gPrice
	// txOption.GasLimit = uint64(3000000)
	// txOption.Nonce = big.NewInt(int64(nonce))
	// // contadd, txt,_, err := todo.DeployTodo(auth,client)
	// // if err != nil {
	// // 	log.Fatalf("Error while getting pending nonce|: %v", err)
	// // }

	cAdd := common.HexToAddress("0x981a8edF8753Ef7C1485a4874877c991a859493d")
	todoStructure, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatalf("error while getting todo structure: %v", err)
	}

	// tx,err := todoStructure.Add(txOption,"This is my first Task")
	// if err !=nil {
	// 	log.Fatalf("err while getting transaction ID %v", err)
	// }

	taskList, err := todoStructure.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		log.Fatalf("err while getting task list %v", err)
	}

	fmt.Println(taskList)
	//fmt.Println(tx.Hash().Hex())

}
