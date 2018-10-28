package main

import (
	"flag"
	"fmt"
	"os"
)

//从命令行接收参数

//func test()  {
//	len(os.Args)//参数的个数
//	os.Args[0]//获取参数的内容
//}


//多行字符串
const useage  = `
	addBlock --data DATA  "add a block to blockChain"
	printChain            "printChain"
`

const AddBlockCmdString  = "AddBlockCmdString"
const PrintChainCmdString  = "PrintChainCmdString"

type CLI struct {
	bc *BlockChain

}



func (cli *CLI)PrintUsage()  {
	fmt.Println(useage)
	os.Exit(1)
}


	//参数检查
func (cli *CLI)paramCheck()  {

	if len(os.Args) < 2{
		fmt.Println("invalid param")
		fmt.Println(useage)
		os.Exit(1)
	}
}


func (cli *CLI)Run()  {

	cli.paramCheck()

	//解析命令
	addBlockCmd := flag.NewFlagSet(AddBlockCmdString,flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString,flag.ExitOnError)

	//func (f *FlagSet) String(name string, value string, usage string) *string {
	//把参数以str的形式返回,data表示参数名称，第二个参数表示默认值
	addBlockCmdPara := addBlockCmd.String("data","","block transatcion info")


	//监听命令
	switch os.Args[1] {
	case AddBlockCmdString:
		//添加动作

		//解析前先检测
		err := addBlockCmd.Parse(os.Args[2:])//解析参数，形如：./block addBlock --data "A to B" , os.Args[2] = (--data "A to B")
		CheckErr("Run() AddBlockCmdString",err)
		if addBlockCmd.Parsed(){
			if *addBlockCmdPara == "" {
				cli.PrintUsage()
			}
			cli.AddBlock(*addBlockCmdPara)
		}

	case PrintChainCmdString:


		//打印输出
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run() PrintChainCmdString",err)

		if printChainCmd.Parsed(){
			cli.PrintUsage()
		}

	default:
		fmt.Println("无效 param")
		cli.PrintUsage()

	}
}