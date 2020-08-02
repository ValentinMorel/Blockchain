package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	// Converts the hash, prev hash and data in byte format into a sha256 hash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (b *BlockChain) ShowInfo() {
	for _, blck := range b.Blocks {
		fmt.Printf("PrevHash Block : %x \n", blck.PrevHash)
		fmt.Printf("Data Block : %s \n", blck.Data)
		fmt.Printf("Hash Block : %x \n\n", blck.Hash)

		pow := NewProof(blck)
		fmt.Printf("PoW : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
