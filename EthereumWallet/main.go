package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvt, err := crypto.GenerateKey()

	if err != nil {
		log.Fatalf("Error while generating private key: %v", err)
	}
	pvtData := crypto.FromECDSA(pvt)
	Private_key := hexutil.Encode(pvtData)
	fmt.Println("Private Key is:", Private_key)

	pubK := pvt.PublicKey
	pubData := crypto.FromECDSAPub(&pubK)
	public_key := hexutil.Encode(pubData)
	fmt.Println("Public key is:", public_key)

	address := crypto.PubkeyToAddress(pubK)

	fmt.Println("Public Address is:", address)

}
