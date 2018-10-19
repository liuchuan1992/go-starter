package main

type Blockchain struct {
	blocks []*Block
}

/**
往区块链中添加一个区块
 */
func(bc * Blockchain) AddBlock(data string){
	preBlock := bc.blocks[len(bc.blocks) - 1] //获取当前的区块
	curBlock :=NewBlock(data,preBlock.Hash) //生成新的区块
	bc.blocks = append(bc.blocks,curBlock)
}

func InitBlockChain() * Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

