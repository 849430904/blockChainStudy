package main

//区块链相关
type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	//创世块
	block := NewGenersisBlock()
	return &BlockChain{blocks: []*Block{block}}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash //当前链条最后一个hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, block)
}
