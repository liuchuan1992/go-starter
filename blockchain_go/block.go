package main

import (
	"crypto/sha256"
	"strconv"
	"time"
	"fmt"
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

/**
往区块链中添加一个区块
 */
func(bc * Blockchain) AddBlock(data string){
	preBlock := bc.blocks[len(bc.blocks) - 1] //获取当前的区块
	curBlock :=NewBlock(data,preBlock.Hash) //生成新的区块
	bc.blocks = append(bc.blocks,curBlock)
}

func NewGenesisBlock() *Block{
	return NewBlock("lvcf's first block !",[]byte{})
}

func InitBlockChain() * Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
func main() {
	blockchain := InitBlockChain()
	blockchain.AddBlock("lvcf -> Yoga , 1BTC")
	blockchain.AddBlock("Yoga -> M. ,0.5BTC")


	for _, block := range blockchain.blocks {
		fmt.Printf("Prev. hash: %x \n",block.PrevBlockHash)
		fmt.Printf("Data: %s \n",block.Data)
		fmt.Printf("Hash: %x \n",block.Hash)
		fmt.Println()
	}
}
