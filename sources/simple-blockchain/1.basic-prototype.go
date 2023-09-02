package prototype

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"fmt"
)


func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to bosung")
	bc.AddBlock("Send 2 more BTC to bosung")

	for i, block := range bc.blocks {
		fmt.Printf("%d: Prev. hash: %x\n", i, block.PrevBlockHash)
		fmt.Printf("%d: Data: %s\n", i, block.Data)
		fmt.Printf("%d: Hash: %x\n", i, block.Hash)
		fmt.Println()
	}
}

// block that store valuable information.
type Block struct {
	Timestamp     int64  // when the block is created
	Data          []byte  // actual valuable information containing in the block
	PrevBlockHash []byte  // the hash of the previous block
	Hash          []byte  // the hash of the block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// blockchain
type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// In its essence, blockchain is just a distributed database of records. 
// But what makes it unique is that it’s not a private database, but a public one.
// blockchain is a distributed database that has no single decision maker. 