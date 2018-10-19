package main

import (
"fmt"
)
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
