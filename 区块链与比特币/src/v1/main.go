package main

import "fmt"

func main() {

	bc := NewBlockChain()
	bc.AddBlock("A send B 1BTC")
	bc.AddBlock("B send C 2BTC")

	for _, block := range bc.blocks {//遍历所有区块
		fmt.Printf("version:%d \n", block.Version)
		fmt.Printf("PrevBlockHash:%x \n", block.PrevBlockHash)
		fmt.Printf("Hash:%x \n", block.Hash)
		fmt.Printf("TimeStamp:%d \n", block.TimeStamp)
		fmt.Printf("Bits:%d \n", block.Bits)
		fmt.Printf("Nonce:%d \n", block.Nonce)
		fmt.Printf("Data:%s \n", block.Data)
	}
}
