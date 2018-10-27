package main

//区块相关
import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块头信息,
type Block struct {
	Version       int64  //版本
	PrevBlockHash []byte //上一个区块的hash
	Hash          []byte //实际中的hash根据区块计算，这里中简化代码
	MerKelRoot    []byte //梅克尔根
	TimeStamp     int64  //时间戳
	Bits          int64  //难度值
	Nonce         int64  //随机数

	//交易信息
	Data []byte
}

//data:交易信息， 前一个区块的哈希
func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash todo
		MerKelRoot: []byte{}, //先为空
		TimeStamp:  time.Now().Unix(),
		Bits:       1,
		Nonce:      1,
		Data:       []byte(data)}

	block.SetHash()
	return &block
}

func (block *Block) SetHash() {

	temp := [][]byte{ //定义一个二维的切片

		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Nonce),
		block.Data}

	//func Join(s [][]byte, sep []byte) []byte {
	data := bytes.Join(temp, []byte{}) //切片拼接
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

//创世块
func NewGenersisBlock() *Block {
	return NewBlock("GenersisBlock", []byte{})
}
