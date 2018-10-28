package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num int64) []byte {

	//w:我们要写的数据 ,order:写的时候对齐方式，如大端对齐，数据
	//func Write(w io.Writer, order ByteOrder, data interface{}) error {

	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr("IntToByte", err)
	return buffer.Bytes()
}

func CheckErr(pos string, err error) {
	if err != nil {
		fmt.Printf("error, pos :", pos, err)
		os.Exit(1)
	}
}
