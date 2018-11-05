package main

import (
	"crypto/md5"
	"encoding/hex"
)

//使用md5对数据进行哈希运算
func GetMd5str_1(src []byte) string  {
	//1，给哈希算法添加数据
	res := md5.Sum(src)

	//2,对数据进行格式化
	//myRes := fmt.Sprintf("%x",res)

	myRes := hex.EncodeToString(res[:])//res[:]表示把固定长度的切片转换为不固定长度的切片
	return myRes
}



//使用md5对数据进行哈希运算
func GetMd5str_2(src []byte) string  {
	//1，给哈希算法添加数据
	myHash := md5.New()

	//2,添加数据
	//添加数据的第一种方式
	//io.WriteString(myHash,string(src))


	//添加数据的第二种方式
	myHash.Write(src)

	//3，计算结果
	res := myHash.Sum(nil)

	//4，散列值格式化
	myRes := hex.EncodeToString(res[:])//res[:]表示把固定长度的切片转换为不固定长度的切片
	return myRes
}
