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
	res := md5.Sum(src)

	//2,对数据进行格式化
	//myRes := fmt.Sprintf("%x",res)

	myRes := hex.EncodeToString(res[:])//res[:]表示把固定长度的切片转换为不固定长度的切片
	return myRes
}
