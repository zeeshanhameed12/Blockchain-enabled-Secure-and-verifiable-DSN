package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://sepolia.infura.io/v3/a2b0e9a36e9a481b8a5356585850b404"
var ganacheURL = "http://localhost:8545"

func main() {

	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		log.Fatalf("Error to creat ether client: %v", err)
	}

	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to creat ether client: %v", err)
	}
	fmt.Println(block.Number())

	add := "0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"

	address := common.HexToAddress(add)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error while getting the balance: %v", err)
	}

	fmt.Println("Blance of the account", float32(balance.Int64())/float32(1000000000000000000))
}
