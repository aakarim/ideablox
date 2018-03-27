package main

import (
	"crypto/sha256"
	"math/big"
)

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1 // Compare against target hash

	return isValid
}
