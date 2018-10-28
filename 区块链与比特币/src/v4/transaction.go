package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

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
	ScriptSig string//解锁收款方的地址，它是一个锁定的脚本。只有收款方才能打开
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





