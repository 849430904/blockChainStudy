package main

import "fmt"

func main()  {

	//desTest()
	//threeDesTest()
	//rsaTest()
	hashTest()
}

//测试DES加解改密
func desTest()  {
	fmt.Println("=====des 加解密")

	src := []byte("少壮不努力，老大。。。")
	key := []byte("12345678")//密钥的长度为8字节
	str := enctyptDES(src ,key )

	println("加密后的数据 ：",str)

	str = descryptDES(src,key)

	println("解密后的数据 ：",str)
}

//测试3DES加解改密
func threeDesTest()  {
	fmt.Println("=====aes 加解密")

	src := []byte("少壮不努力，老大。。。")
	key := []byte("12345678abcdefgh12345678")//密钥的长度为24字节
	str := encrypt3DES(src ,key )

	println("加密后的数据 ：",str)

	str = decrypt3DES(src,key)

	println("解密后的数据 ：",str)
}



//测试AES加解改密
func aesTest()  {
	fmt.Println("=====3des 加解密")

	src := []byte("少壮不努力，老大。。。")
	key := []byte("12345678abcdefgh")//密钥的长度为16字节
	str := encryptAES(src ,key )

	println("加密后的数据 ：",str)

	str = decryptAES(src,key)

	println("解密后的数据 ：",str)
}

func rsaTest()  {
	fmt.Println("=====rsa 加解密")
	err := RsaGenKey(4096)//4096长度的私钥与公钥
	fmt.Println("错误信息:",err)

	//加密
	src := []byte("少壮不努力，老大。。。")
	data,err := RsaPublicEncrypt(src,"public.pem")

	//解密
	data,err = RsaPrivateDecrypt(data,"private.pem")
	fmt.Println("非对称解密的结果:",string(data))
	fmt.Println("err:",err)
}

func hashTest()  {

	data := []byte("少壮不努力，老大。。。")
	hash := GetMd5str_1(data)
	fmt.Println("hash:",hash)
	hash = GetMd5str_2(data)
	fmt.Println("hash:",hash)
}
