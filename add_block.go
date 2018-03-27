package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tecbot/gorocksdb"
)

// AddBlock - add block to existing blockchain
func (bc *Blockchain) AddBlock(data []byte) {
	// Get the last saved hash
	ro := gorocksdb.NewDefaultReadOptions()
	lastHash, err := bc.db.Get(ro, []byte("l"))
	defer lastHash.Free()
	if err != nil {
		log.Fatalln("Error getting last hash from blockchain:", err)
	}
	newBlock := NewBlock(data, lastHash.Data())

	wo := gorocksdb.NewDefaultWriteOptions()

	serialisedNewBlock, err := proto.Marshal(newBlock)
	if err != nil {
		log.Fatalln("Could not serialise new block:", err)
	}
	bc.db.Put(wo, newBlock.Hash, serialisedNewBlock)
	bc.db.Put(wo, []byte("l"), newBlock.Hash) // Save tip
	bc.tip = newBlock.Hash
}
