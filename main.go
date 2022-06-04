package main

import (
	"fmt"
	"gochain/model"
)

func main() {
	bc := model.CreateBlockchain()

	bc.AddBlock("Send 1 Gochain to X")
	bc.AddBlock("Send 2 Gochain to Y")

	for _, block := range bc.Blocks {
		fmt.Printf("previous block hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Println()
	}
}
