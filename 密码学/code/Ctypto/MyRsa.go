package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//非对称加密


//生成公钥与私钥的函数
func RsaGenKey(bits int)  error{

	/**************生成私钥文件***********/

	//1,使用rsa中的GenerateKey方法生成私钥
	//bits指的是私钥长度
	privateKey,err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		return err
	}

	//2,通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)

	//3,将私钥字符串设置到pem格式中
	block := pem.Block{
		Type:"RSA Private Key",
		Bytes:privateStream,

	}
	//4，通过pem包将设置好的数据进行编码，并写入到磁盘中去
	privateFile ,err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	err = pem.Encode(privateFile,&block)//将key写磁盘（会写到文件中去）
	if err != nil {
		return err
	}

	/**************生成公钥文件***********/

	//1，将得到的私钥对象中将公钥信息取出
	pubKey := privateKey.PublicKey //得到公钥

	//2,通过x509标准将得到的rsa公钥序列化为字符串
	pubStream ,err  := x509.MarshalPKIXPublicKey(&pubKey)//注意：这里一定要传地址
	if err != nil {
		return err
	}
	//3,将公钥字符串设置到pem格式块中
	publicBlock := pem.Block{
		Type:"RSA Public Key",
		Bytes:pubStream,
	}

	//4,通过pem将设置好的数据进行编码，并写入到磁盘中
	publicFile ,err := os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(publicFile,&publicBlock)//写磁盘
	if err != nil {
		return err
	}
	defer publicFile.Close()
	return nil
}


//利用公钥对数据进行加密
//src 待加密的数据 ，pathName：公钥文件的路径
func RsaPublicEncrypt(src []byte,pathName string)([]byte, error) {

	msg := []byte("")
	//1,将公钥文件中的公钥读出，得到使用pem编码的字符串
	file , err :=  os.Open(pathName)
	defer file.Close()
	if err != nil {
		return msg,err
	}
	//1.1 先得到文件属性信息，通过属性信息对象得到文件大小
	fileInfo , err := file.Stat()
	if err != nil {
		return msg,err
	}
	recvBuffer := make([]byte,fileInfo.Size())
	file.Read(recvBuffer)//把数据读取到了recvBuffer中
	//2，将得到的字符串解码（因为读出来的是pem格式的字符串）
	block ,_  := pem.Decode(recvBuffer)

	//3,使用x509将编码之后的公钥解析出来
	pubInter,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return msg,err
	}

	pubKey := pubInter.(*rsa.PublicKey)//类型转换  这里得到实际的公钥
	//4,使用得到的公钥通过rsa进行数据加密
	msg,err = rsa.EncryptPKCS1v15(rand.Reader,pubKey,src)
	if err != nil {
		return msg,err
	}

	return msg,nil
}


//使用私钥解密

func RsaPrivateDecrypt(src []byte,pathName string)([]byte,error){

	msg := []byte("")
	//1,打开私钥文件
	file , err :=  os.Open(pathName)
	defer file.Close()
	if err != nil {
		return msg,err
	}

	//2,读取文件内容
	fileInfo , err := file.Stat()
	if err != nil {
		return msg,err
	}
	recvBuffer := make([]byte,fileInfo.Size())
	file.Read(recvBuffer)//把数据读取到了recvBuffer中

	//3，将得到字符串进行解码（因为读出来的是pem格式的字符串）
	block ,_  := pem.Decode(recvBuffer)

	//4,使用x509将编码之后的私钥解析出来
	privateKey,_ := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg,err
	}

	//5，对数据进行解密
	msg,err = rsa.DecryptPKCS1v15(rand.Reader,privateKey,src)
	if err != nil {
		return msg,err
	}
	return msg,nil
}

