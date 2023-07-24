package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block headers
type Block struct {
	Time     	int64
	Value   	[]byte
	PrevHash 	[]byte
	Hash	[]byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	// Parameters:
	// Data: A string that goes into the block
	// prevBlockHash: The Hash of the previous block, represented as an array of bytes.

	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	timestamp := []byte(strconv.FormatInt(b.Time, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Value, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]	
	return block
}

func CreateGenesisBlock() *Block {
	// This function creates the first block in the block chain.
	return NewBlock("Genesis Block", []byte{})
}



type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlockchain() *Blockchain {
	// This function creates a new block chain where the first block is the Genesis Block.
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}
