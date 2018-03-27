package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/tecbot/gorocksdb"
)

const dbFile = "ideablox.db"

// NewBlockChain - create new blockchain with genesis block from NewGenesisBlock
func NewBlockChain() *Blockchain {
	var tip []byte
	// Set up connection to db
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	shouldCreateGenesisBlock := false
	/* If no existing blockchain create genesis block */
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		shouldCreateGenesisBlock = true
	}
	db, err := gorocksdb.OpenDb(opts, dbFile)
	if err != nil {
		log.Panic("Could not connect to database:", err)
	}
	if shouldCreateGenesisBlock == true {
		fmt.Println("No existing blockchain found. Creating a new one...")
		// Generate new block
		genesis := NewGenesisBlock()
		// serialise the genesis block
		serialisedGenesis, err := proto.Marshal(genesis)
		if err != nil {
			log.Fatalln("Could not serialise genesis block:", err)
		}
		// write to db
		wo := gorocksdb.NewDefaultWriteOptions()
		err = db.Put(wo, genesis.Hash, serialisedGenesis) // save block
		if err != nil {
			log.Fatalln("Could not store genesis block:", err)
		}
		err = db.Put(wo, []byte("l"), genesis.Hash) // last block of the hash chain
		if err != nil {
			log.Fatalln("Could not store last hash for genesis block:", err)
		}
		tip = genesis.Hash
	} else {
		// Get existing tip
		ro := gorocksdb.NewDefaultReadOptions()
		tipRaw, err := db.Get(ro, []byte("l"))
		defer tipRaw.Free()
		tip = tipRaw.Data()
		fmt.Println("Using existing blockchain")
		if err != nil {
			log.Fatalln("Could not get tip hash:", err)
		}
	}
	return &Blockchain{tip, db}
}
