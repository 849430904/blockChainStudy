package main
import "fmt"

//演示golang中+的使用
func main() {
	
	var i = 1
	var j = 2
	var r = i + j //做加法运算； 如果是数字自动做加法运算
	fmt.Println("r=", r)

	var str1 = "hello "
	var str2 = "world"
	var res = str1 + str2 //做拼接操作； 如果是字符串自动做拼接
	fmt.Println("res=", res)

}