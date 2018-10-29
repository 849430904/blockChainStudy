package main

import "fmt"


func (cli *CLI)AddBlock(data string)  {
	//bc := GetBlockChainHandler()
	//bc.AddBlock(data)
}

//打印数据
func (cli *CLI)PrintChain()  {

	bc := GetBlockChainHandler()
	defer bc.db.Close()

	//打印数据
	it := bc.NewIterator()

	for  {
		block := it.Next()
		fmt.Printf("version:%d \n", block.Version)
		fmt.Printf("PrevBlockHash:%x \n", block.PrevBlockHash)
		fmt.Printf("Hash:%x \n", block.Hash)
		fmt.Printf("TimeStamp:%d \n", block.TimeStamp)
		fmt.Printf("Bits:%d \n", block.Bits)
		fmt.Printf("Nonce:%d \n", block.Nonce)
		//fmt.Printf("Data:%s \n", block.Data)
		fmt.Printf("isValild:%v \n", NewProofOfWork(block).isValid())

		if len(block.PrevBlockHash) == 0{
			fmt.Printf("end...")
			break
		}
	}

}


func (cli *CLI)CreateChain(address string)  {
	bc := NewBlockChain(address)
	bc.db.Close()
	fmt.Println("create CreateChain successfully...")
}

func (cli *CLI)GetBalance(address string) float64 {
	bc := GetBlockChainHandler()
	utxos := bc.FindUTXO(address)

	//余额
	var total float64 = 0

	//遍历所有utxo，获取总金额数
	for _,utxo := range utxos{
		total += utxo.Value
	}
	fmt.Println("The balance of %s is %f \n",address,total)
	return total
}