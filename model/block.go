package model

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp int64  // when the block has been created
	Data      []byte // data of the block
	PrevHash  []byte // hash of the previous block in the blockchain
	Hash      []byte // has of the current block
	Nonce     int    // nonce
}

// Serialize a block
func (b *Block) Serialize() []byte {
	var serialized bytes.Buffer
	encoder := gob.NewEncoder(&serialized)
	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}
	return serialized.Bytes()
}

// Create a new block
func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := CreateProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Deserialize a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}

	return &block
}
