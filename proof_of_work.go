package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce int64 = math.MaxInt64 // highest value of nonce to avoid overflow
)

const targetBits = 24

// NewProofOfWork - generate proof of work for block
// target is upper boundary of range of acceptable values
// if produced hash is lower than the target then it is valid
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

// make the data ready to hash
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int // Integer representation of hash
	var hash [32]byte
	var nonce int64
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < int64(maxNonce) { // maxNonce is to avoid possible overflows of nonce
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 { // Compare against the target
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}
