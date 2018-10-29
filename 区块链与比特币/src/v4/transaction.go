package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const reward  =  12.5
type Transaction struct {
	TXID []byte//交易ID
	TXInputs  []TXInput//输入
	TXOutputs []TXOutput//输出

}

type TXInput struct {
	TXID  []byte//所引用 输出的交易ID
	Vout  int64//所引用output的索引值
	ScriptSig string//解锁脚本，指明可以使用某个output的条件
}

type TXOutput struct {
	Value  float64//交易金额；支付给收款方的金额数
	ScriptPubKey string//解锁收款方的地址，它是一个锁定的脚本。只有收款方才能打开
}


//设置交易ID,是一个哈希值
func (tx *Transaction)SetTXID()  {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	CheckErr("SetTXID",err)
	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

//创建coinbase,只有收款人，没有付款人，是矿工的奖励交易
func NewCoinbaseTx(address string,data string) *Transaction  {

	if data == ""{
		data = fmt.Sprintf("reward to %s %d btc",address,reward)
	}
	inputs := TXInput{[]byte{},-1,""}
	outputs := TXOutput{reward,address}//挖矿奖励

	tx := Transaction{[]byte{},[]TXInput{inputs},[]TXOutput{outputs}}
	tx.SetTXID()
	
	return &tx

}

//utxo能否解锁 ; 检查当前的用户能否解开引用的utxo
func (input *TXInput)CanUnlockUTXOWith(unlockedData string)bool  {

	return input.ScriptSig == unlockedData
}


//检查当前用户是否是这个utxo的所有者
func (output *TXOutput)CanBeUnlockUTXOWith(unlockedData string)bool  {

	return output.ScriptPubKey == unlockedData
}





//Coinbase交易iD为空，vout = -1
func (tx *Transaction)IsCoinbase() bool {
	if len(tx.TXInputs) == 1 {
		if len(tx.TXInputs[0].TXID) == 0 && tx.TXInputs[0].Vout == -1{
			return true
		}
	}
	return false
}


//创建普通交易，send的辅助函数
func NewTransaction(from , to string ,amount float64,bc *BlockChain) *Transaction {

	//vaildUTXOs:所需要的，合适的utxo  map[string][]int64
	//total:返回的utxo的金额的总和
	vaildUTXOs,total := bc.FindSuitableUTXOs(from,amount)//总到一些合适的utxo，以便支付


	if total < amount {
		fmt.Println("Not enouth money!")
	}

	//进行output转换到input

	tx := Transaction{[]byte{},[]TXInput{},[]TXOutput{}}
	tx.SetTXID()
	return &tx
}

