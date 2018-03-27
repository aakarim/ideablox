package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tecbot/gorocksdb"
)

// Iterator creates a new BlockchainIterator and is associated with blockchain due to db connection
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	block := &Block{}

	ro := gorocksdb.NewDefaultReadOptions()

	value, err := i.db.Get(ro, i.currentHash)
	defer value.Free()
	if err != nil {
		log.Fatalln("Failed to get block from db:", err)
	}
	if err := proto.Unmarshal(value.Data(), block); err != nil {
		log.Fatalln("Failed to parse Block:", err)
	}

	i.currentHash = block.PrevBlockHash
	return block
}
