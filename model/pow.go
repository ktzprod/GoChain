package model

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"gochain/utils"
	"math"
	"math/big"
)

const targetsBits = 24
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// create a Proof of work instance
func CreateProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetsBits))
	return &ProofOfWork{b, target}
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			p.block.PrevHash,
			p.block.Data,
			utils.IntToHex(p.block.Timestamp),
			utils.IntToHex(int64(targetsBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
}

// Mine a block given the block itself and the target hash
func (p *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", p.block.Data)
	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// validate that this proof of work is actually related to the given block
func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(p.target) == -1

	return isValid
}
