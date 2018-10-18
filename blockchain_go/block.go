package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//简易区块链的数据结构
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

type Blockchain struct {
	blocks []*Block
}
func SetHash(b *Block) {
	h := sha256.New()
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	var data []byte
	data = append(data,timestamp...)
	data = append(data,b.PrevBlockHash...)
	data = append(data,b.Hash...)
	h.Write(data)
	b.Hash = h.Sum(nil)
}

/**
生成一个区块
 */
func NewBlock( data string, preBlockhash []byte ) *Block{
	block := &Block{time.Now().Unix(),[]byte(data),preBlockhash,[]byte{}}
	SetHash(block)
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


func main() {
	h := sha256.New()
	fmt.Println(h)

}
