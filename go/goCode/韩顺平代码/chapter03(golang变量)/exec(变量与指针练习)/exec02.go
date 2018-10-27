package main
import "fmt"

func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)//%v输出地址

	var ptr *int 
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num =" , num)
	var a_b int = 20
	fmt.Println(a_b)

	var int int = 30//变量可以与关键字同名
	fmt.Println(int)
}