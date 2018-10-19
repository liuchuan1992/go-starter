package main

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"
)

//挖矿的难度值 这里表示哈希的前 24 位必须是 0
const targetBits  = 24

//设置nonce的最大值
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork( b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits)) //将big.Int的1 向左移256-24位，所以比这个值小的
	return &ProofOfWork{b,target}
}

//工作量证明的数据准备，需要用到PreBlockHash,Data,timestamp,nonce，targetbits
func (pow * ProofOfWork)prepareData(nonce int) []byte  {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//工作量证明的核心方法
func (pow *ProofOfWork) Run() (int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data) //sha256.Sum用来计算byte数组的hash值，返回一个32位长度的字节数组 所以这边的hash定义为【32】byte
		hashInt.SetBytes(hash[:]) //setByte方法接受slice 这边将数组转slice 使用hash[:]

		if hashInt.Cmp(pow.target)  == -1 {
			fmt.Printf("\r%x", hash)
			break
		}
			nonce++
	}
	return nonce , hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}