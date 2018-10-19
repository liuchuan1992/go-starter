package main

import (
	"time"
)

//简易区块链的数据结构
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}


/**
生成一个区块
 */
func NewBlock( data string, preBlockhash []byte ) *Block{
	block := &Block{time.Now().Unix(),[]byte(data),preBlockhash,[]byte{},0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}


func NewGenesisBlock() *Block{
	return NewBlock("lvcf's first block !",[]byte{})
}
