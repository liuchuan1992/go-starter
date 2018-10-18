package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

//简易区块链的数据结构
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
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

func main() {
	h := sha256.New()
	fmt.Println(h)
}
