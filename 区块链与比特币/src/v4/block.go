package main



//区块相关

import (
	"bytes"
	"encoding/gob"
	"time"
)




//区块头信息,
type Block struct {
	Version       int64  //版本
	PrevBlockHash []byte //上一个区块的hash
	MerKelRoot    []byte //梅克尔根
	TimeStamp     int64  //时间戳
	Bits          int64  //难度值
	Nonce         int64  //随机数

	//交易信息
	//Data []byte
	Transactions []*Transaction
	Hash []byte //实际中的hash根据区块计算，这里中简化代码
}

//区块的序列化
func (block *Block)Seriallize() []byte {
	var buffer bytes.Buffer

	//func NewEncoder(w io.Writer) *Encoder {
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)//encode后，所有的字节流在buffer里面
	CheckErr("Seriallize",err)

	return  buffer.Bytes()

}

//自由函数,返序列化
func Deserialize(data []byte) *Block  {

	if len(data) == 0{
		return nil
	}

	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block) //解码后的数据在block里面
	CheckErr("Deserialize",err)
	return  &block
}


//data:交易信息， 前一个区块的哈希
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash todo
		MerKelRoot: []byte{}, //先为空
		TimeStamp:  time.Now().Unix(),
		Bits:       targetBits,//难度值
		Nonce:      0,
		Transactions: txs}

	//block.SetHash()
	pow := NewProofOfWork(&block)//用工作量证明
	nonce,hash := pow.Run()//工作量证明，计算hash值
	block.Nonce = nonce
	block.Hash = hash

	return &block
}

//
//func (block *Block) SetHash() {
//
//	temp := [][]byte{ //定义一个二维的切片
//
//		IntToByte(block.Version),
//		block.PrevBlockHash,
//		block.MerKelRoot,
//		IntToByte(block.TimeStamp),
//		IntToByte(block.Nonce), block.Data}
//
//	//func Join(s [][]byte, sep []byte) []byte {
//	data := bytes.Join(temp, []byte{}) //切片拼接
//	hash := sha256.Sum256(data)
//	block.Hash = hash[:]
//}

//创世块
func NewGenersisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})

}



