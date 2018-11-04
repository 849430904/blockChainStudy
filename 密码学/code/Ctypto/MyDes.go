package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

//DES加密

//填充最后一个分组的函数
/*
src:原始数据
blockSize:每一个分组的数据长度
*/
func paddingText(src []byte,blockSize int) []byte {
	//1、求出一个分组要填充多少个字节
	padding := blockSize - len(src)%blockSize

	//2,创建新的切片，切片的字节数为padding,并初始化，每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)},padding)

	//3,将创建出的新切片和原始数据进行拼接
	newText := append(src,padText...)

	//返回新的字符串
	return newText


}

//删除末尾填充的字节数
func unPaddingText(src []byte) []byte  {

	//1，求出要处理的切片的长度
	len := len(src)

	//2,取出最后一个字符，并把它的整型值给取出来
	number := int(src[len - 1])// byte -> int

	//3，将切片末尾的Number个字节删除掉
	newText := src[:len - number]

	return newText
}


//使用des进行对称加密
//src 源数据
//key 密钥
func enctyptDES(src , key []byte)  [] byte{
	//创建并返回一个使用des算法的clpher.Block接口
	block , err := des.NewCipher(key)
	if err != nil{
		panic(err)
	}

	//2,对最后一个明文分组进行数据填充
	paddingText(src,block.BlockSize())

	//3,创建一个密码分组为链接模式,底层使用DES加密的BlockMode接口
	iv := []byte("aaabbbbb")
	blockMode := cipher.NewCBCDecrypter(block,iv)

	//4,加密连续的数据块
	dst := make([]byte,len(src))
	blockMode.CryptBlocks(dst,src)

	return dst

}

//使用des解密
func descryptDES(src, key []byte) [] byte  {

	//1，创建并返回一个使用des算法的clpher.Block接口
	block , err := des.NewCipher(key)
	if err != nil{
		panic(err)
	}

	//2,创建一个密码分组为链接模式的，底层使用des解密的blockMode接口
	iv := []byte("aaabbbbb")
	blockMode := cipher.NewCBCDecrypter(block,iv)

	//3,数据块解密
	blockMode.CryptBlocks(src,src)

	//4,去掉最后一组的填充数据
	newText := unPaddingText(src)//得到明文

	return newText

}


//使用3des进行加密操作
func encrypt3DES(src,key []byte)[]byte  {
	//1，创建并返回一个使用3des算法的clpher.Block接口
	block , err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}

	//2，对最后一个明文分组进行数据填充
	src = paddingText(src,block.BlockSize())

	//3，创建一个密码分组为链接模式的，底层使用3DES加密的BlockMode接口
	//key[:block.BlockSize()],表示截取开始位置到结束位置
	blockMode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])

	//4，加密连接的数据块
	blockMode.CryptBlocks(src,src)

	return src

}

//使用3DES对数据解密

func decrypt3DES(src,key []byte) []byte  {

	//1，创建并返回一个使用3des算法的clpher.Block接口
	block , err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}

	//2，创建一个密码分组为链接模式的，底层使用3DES加密的BlockMode接口
	//key[:block.BlockSize()],表示截取开始位置到结束位置
	blockMode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])

	//3,数据块解密
	blockMode.CryptBlocks(src,src)

	//4,去掉最后一组填充数据
	src  = unPaddingText(src)//解密后的明文

	return src
}