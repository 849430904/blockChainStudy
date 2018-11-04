package main

import "fmt"

func main()  {

	//desTest()
	threeDesTest()
}

//测试DES加解改密
func desTest()  {
	fmt.Println("=====des 加解密")

	src := []byte("少壮不努力，老大。。。")
	key := []byte("12345678")
	str := enctyptDES(src ,key )

	println("加密后的数据 ：",str)

	str = descryptDES(src,key)

	println("解密后的数据 ：",str)
}

//测试3DES加解改密
func threeDesTest()  {
	fmt.Println("=====3des 加解密")

	src := []byte("少壮不努力，老大。。。")
	key := []byte("12345678abcdefgh12345678")
	str := encrypt3DES(src ,key )

	println("加密后的数据 ：",str)

	str = decrypt3DES(src,key)

	println("解密后的数据 ：",str)
}