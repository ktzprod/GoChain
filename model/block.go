package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func (b *Block) Serialize() []byte {
	var serialized bytes.Buffer
	encoder := gob.NewEncoder(&serialized)
	if err := encoder.Encode(b); err != nil {
		fmt.Printf("Failed to serialize block: %s\n", err)
	}
	return serialized.Bytes()
}

func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := CreateProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		fmt.Printf("Failed to deserialize block: %s\n", err)
		return nil
	}

	return &block
}
