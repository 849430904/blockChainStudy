package main

import (
	"crypto/aes"
	"crypto/cipher"
)

//aes加密
func encryptAES(src,key []byte) []byte   {

	//1，创建并返回一个使用aes算法的clpher.Block接口
	block , err := aes.NewCipher(key)
	if err != nil{
		panic(err)
	}

	//2,填充数据
	src = paddingText(src,block.BlockSize())

	//3,创建一个密码分组为链接模式的，底层使用AES加密的blockMode接口
	blockMode := cipher.NewCBCEncrypter(block,key)

	//4,数据加密
	blockMode.CryptBlocks(src,src)//不创建数组了，直接使用src

	return src
}


//aes解密
//src 要解密的数据
//key :密钥
func decryptAES(src,key []byte) []byte {

	//1，创建并返回一个使用aes算法的clpher.Block接口
	block , err := aes.NewCipher(key)
	if err != nil{
		panic(err)
	}

	//2,创建一个密码分组为链接模式的，底层使用AES加密的blockMode接口
	blockMode := cipher.NewCBCEncrypter(block,key)

	//3,数据块解密
	blockMode.CryptBlocks(src,src)

	//4 去掉最后一组的填充数据
	src = unPaddingText(src)

	return src
}