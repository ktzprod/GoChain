package main

import (
	"fmt"
	"gochain/model"
	"strconv"
)

func main() {
	bc := model.CreateBlockchain()

	bc.AddBlock("Send 1 Gochain to X")
	bc.AddBlock("Send 2 Gochain to Y")

	for _, block := range bc.Blocks {
		fmt.Printf("previous block hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := model.CreateProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
