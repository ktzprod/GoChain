package model

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}
