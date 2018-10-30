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
	createChain --address ADDRESS  "create a  blockchain"
	addBlock --data DATA  "add a block to blockChain"
    send    --from From --to TO --amount AMOUNT  "send coin from FROM to TO"
    getbalance --address ADDRESS  "get balance of address"
	printChain            "printChain"
`

const CreateChainCmdString  = "createChain"
const AddBlockCmdString  = "addBlock"
const PrintChainCmdString  = "printChain"
const SendCmdString  = "send"//转账
const GetbalanceCmdString  = "getbalance"//获取某个账户的余额

type CLI struct {

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
	createChainCmd := flag.NewFlagSet(CreateChainCmdString,flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet(AddBlockCmdString,flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString,flag.ExitOnError)
	getbalanceCmd := flag.NewFlagSet(GetbalanceCmdString,flag.ExitOnError)
	sendCmd := flag.NewFlagSet(SendCmdString,flag.ExitOnError)

	//func (f *FlagSet) String(name string, value string, usage string) *string {
	//把参数以str的形式返回,data表示参数名称，第二个参数表示默认值
	addBlockCmdPara := addBlockCmd.String("data","","block transatcion info")

	//创建区块链相关参数
	createChainCmdPara := addBlockCmd.String("address","","address info")

	//余额相关
	getbalanceCmdPara := addBlockCmd.String("address","","address info")

    //send相关参数
	fromPara := addBlockCmd.String("from","","sender addrss info ")
	toPara := addBlockCmd.String("to","","to address info")
	amountPara := addBlockCmd.Float64("amount",0,"amount info")


	//监听命令
	switch os.Args[1] {

	case CreateChainCmdString:

		//创建区块链
		err := createChainCmd.Parse(os.Args[2:])//解析参数，形如：./block addBlock --data "A to B" , os.Args[2] = (--data "A to B")
		CheckErr("Run()1 AddBlockCmdString",err)
		if createChainCmd.Parsed(){
			if *createChainCmdPara == "" {
				fmt.Println("address should not be empty")
				cli.PrintUsage()
			}
			cli.CreateChain(*createChainCmdPara)
		}
	case AddBlockCmdString:
		//添加动作

		//解析前先检测
		err := addBlockCmd.Parse(os.Args[2:])//解析参数，形如：./block addBlock --data "A to B" , os.Args[2] = (--data "A to B")
		CheckErr("Run() AddBlockCmdString",err)
		if addBlockCmd.Parsed(){
			if *getbalanceCmdPara == "" {
				fmt.Println("address should not be empty")
				cli.PrintUsage()
			}
			cli.GetBalance(*addBlockCmdPara)
		}

	case GetbalanceCmdString://获取余额

		err := getbalanceCmd.Parse(os.Args[2:])//解析参数，形如：./block addBlock --data "A to B" , os.Args[2] = (--data "A to B")
		CheckErr("Run() getbalanceCmd",err)
		if getbalanceCmd.Parsed(){
			if *addBlockCmdPara == "" {
				cli.PrintUsage()
			}
			//	cli.AddBlock(*addBlockCmdPara)
		}

	case PrintChainCmdString:

		//打印输出
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run() PrintChainCmdString",err)

		if printChainCmd.Parsed(){
			cli.PrintChain()
		}

	case SendCmdString:

		err := sendCmd.Parse(os.Args[2:])
		CheckErr("Run() sendCmd",err)
		if sendCmd.Parsed(){
			if *fromPara == "" && *toPara == "" && *amountPara == 0{//参数校验
			    fmt.Println("send cmd paramter invalid")
				cli.PrintUsage()
			}
			cli.send(*fromPara,*toPara,*amountPara)
		}
	default:
		fmt.Println("无效 param")
		cli.PrintUsage()

	}

	//fromPara := addBlockCmd.String("from","","sender addrss info ")
	//toPara := addBlockCmd.String("to","","to address info")
	//amountPara := addBlockCmd.String("amount","","amount info")
}

//转账
func (cli *CLI)send(from,to string,amount float64)  {
	bc := GetBlockChainHandler()
	defer  bc.db.Close()
    tx := NewTransaction(from,to,amount,bc)
    bc.AddBlock([]*Transaction{tx})
    fmt.Println("send successfully")

}

