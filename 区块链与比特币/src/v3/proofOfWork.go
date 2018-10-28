package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWork struct {
	block    *Block
	target   *big.Int //目标值，比这个值小就算找到了hash值
}

const  targetBits  = 24
//实际中需要动态调整，这里找出前置6个0为我们的目标
func NewProofOfWork(block *Block) *ProofOfWork{

	//000000000000...01
	target := big.NewInt(1)
	//0x000001000000...01  16进制，256bit
	target.Lsh(target,256 - targetBits)//左移 256-24位,6位
	pow := ProofOfWork{block:block,target:target}
	return &pow
}

//构造hash对应的Data
func (pow *ProofOfWork)PrepareData(nonce int64)[]byte  {

	block := pow.block
	temp := [][]byte{ //定义一个二维的切片

		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(targetBits),
		IntToByte(nonce),
		block.Data}

	//func Join(s [][]byte, sep []byte) []byte {
	data := bytes.Join(temp, []byte{}) //切片拼接
	return data
}

func (pow *ProofOfWork)Run()(int64, []byte){

    //1，拼装数据

    //2, 类型转换 哈希值转成big.Int类型

    var hash [32]byte
	var nonce int64 = 0
	var hashInt big.Int

	fmt.Printf("开始挖矿啦....\n")
	fmt.Printf("target hash:%x \n",pow.target.Bytes())//切片化输出

    for nonce < math.MaxInt64 {//不断的尝试

    	data := pow.PrepareData(nonce)//nonce 随机值，返回目标data
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])//[:] 表示切片

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if hashInt.Cmp(pow.target) == -1{//跟目标值对比，找到了
			fmt.Printf("found nonce,nonce : %d ,hash:%x \n",nonce,hash)
			break //退出循环
		}else {
			//fmt.Printf("not found nonce,current nonce : %d ,current hash:%x \n",nonce,hash)
			nonce++
		}
	}
	return nonce,hash[:]
}



//校验工作量 hash
func (pow *ProofOfWork)isValid() bool {
	var hashInt big.Int

    data := pow.PrepareData(pow.block.Nonce)
    hash := sha256.Sum224(data)
	hashInt.SetBytes(hash[:])

    return hashInt.Cmp(pow.target) == -1
}
