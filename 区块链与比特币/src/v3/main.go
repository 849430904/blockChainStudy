package main

func main() {

	//""
	//bc := NewBlockChain()
	//bc.AddBlock("A send B 1BTC")
	//bc.AddBlock("B send C 2BTC")
	//
	//for _, block := range bc.blocks {
	//	fmt.Printf("version:%d \n", block.Version)
	//	fmt.Printf("PrevBlockHash:%x \n", block.PrevBlockHash)
	//	fmt.Printf("Hash:%x \n", block.Hash)
	//	fmt.Printf("TimeStamp:%d \n", block.TimeStamp)
	//	fmt.Printf("Bits:%d \n", block.Bits)
	//	fmt.Printf("Nonce:%d \n", block.Nonce)
	//	fmt.Printf("Data:%s \n", block.Data)
	//	fmt.Printf("isValild:%v \n", NewProofOfWork(block).isValid())
	//}



	//通过命令行添加
	bc := NewBlockChain()
	cli := CLI{bc:bc}
	cli.Run()

	//执行：
	// ./block addBlock --data "F to E 1TBC"
	// ./block printChain
}


