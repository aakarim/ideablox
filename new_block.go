package main

import "time"

// NewBlock - create new block with hash
func NewBlock(data []byte, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), data, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run() // Run proof of work before creating block
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
