package main

import (
	"math/big"

	"github.com/tecbot/gorocksdb"
)

// Blockchain - current working blockchain
type Blockchain struct {
	tip []byte
	db  *gorocksdb.DB
}

// ProofOfWork - holds pointer to block and pointer to target
/*
* target < 256 bits in memory. Bigger the difference the more difficult to find proper hash
 */
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// BlockchainIterator is used to avoid all the blocks being loaded into memory on read
type BlockchainIterator struct {
	currentHash []byte
	db          *gorocksdb.DB
}

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}
