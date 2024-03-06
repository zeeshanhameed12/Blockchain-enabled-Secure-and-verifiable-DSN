package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// acc, err := key.NewAccount("Pakistan")

	// if err != nil {
	// 	log.Fatalf("Error while generating the wallet: %v", err)
	// }

	// fmt.Println("Account:", acc.Address)
	ew, err := os.ReadFile("./wallet/UTC--2024-03-06T15-34-17.401375233Z--e016fe9635ed22f72156c5a07442dd7662de53c0")
	if err != nil {
		log.Fatalf("Error while reading wallet file: %v", err)
	}

	key, err := keystore.DecryptKey(ew, "Pakistan")
	if err != nil {
		log.Fatalf(" Error while decrypting the encrypted wallet: %v", err)
	}

	kData := crypto.FromECDSA(key.PrivateKey)
	private_key := hexutil.Encode(kData)

	fmt.Println("Private key: ", private_key)

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	pub_key := hexutil.Encode(pubData)

	fmt.Println("Public key: ", pub_key)
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println("Public address: ",add)
}
