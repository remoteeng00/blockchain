package core

import (
	"crypto/rsa"
	"time"

	"../config"
)

//InitializeBlockchainWithDiff creates a blockchain from scratch
func InitializeBlockchainWithDiff(gensisAddress *rsa.PublicKey, diff Difficulty) Blockchain {
	var chain Blockchain
	chain.txMap = make(map[[config.HashSize]byte]*Transaction)
	chain.utxoMap = make(map[UTXO]bool)
	chain.blockMap = make(map[[config.HashSize]byte]*Block)
	chain.difficulty = diff
	chain.AddressMap = make(map[rsa.PublicKey]map[UTXO]bool)
	chain.TransactionPool = make(map[string]*Transaction)

	gensisBlock := CreateFirstBlock(uint64(time.Now().UnixNano()/1000000), gensisAddress)
	chain.performMinerTransactionAndAddBlock(gensisBlock)

	return chain
}
