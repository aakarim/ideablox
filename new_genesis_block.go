package main

import (
	"log"

	"github.com/golang/protobuf/proto"
)

// NewGenesisBlock - generate the genesis block (which is always owned by aakarim)
func NewGenesisBlock() *Block {
	idea := &Idea{Text: "Create IdeaBlox", Author: "aakarim"}
	out, err := proto.Marshal(idea)
	if err != nil {
		log.Panic(err)
	}
	return NewBlock(out, []byte{})
}
