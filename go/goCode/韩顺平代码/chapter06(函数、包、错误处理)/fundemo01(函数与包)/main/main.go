package main
import (
	"fmt"
	"go_code/chapter06/fundemo01/utils"//引入包

	/**

	在 import 包时，路径从 $GOPATH 的 src 下开始，不用带 src , 编译器会自动从 src 下开始引入
	为了让其它包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其它语言 的 public ,这样才能跨包访问。比如 utils.go 的
	在访问其它包函数，变量时，其语法是 包名.函数名， 比如这里的 main.go 文件中
	如果包名较长，Go 支持给包取别名， 注意细节:取别名后，原来的包名就不能使用了
	如果你要编译成一个可执行程序文件，就需要将这个包声明为 main , 即 package main .这个就
    是一个语法规范，如果你是写一个库 ，包名可以自定义
	**/
)


func main() {

	fmt.Println("utils.go Num~=", utils.Num1)
	//请大家完成这样一个需求:
	//输入两个数,再输入一个运算符(+,-,*,/)，得到结果.。
	//分析思路....
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := utils.Cal(n1, n2 , operator) 
	fmt.Println("result~=", result)


	//代码...
	//代码...
	//代码...

	//有需求，输入两个数num1, num2，计算 + / *  - 的值

	n1 = 4.5
	n2 = 6.7
	operator = '*'
	result = utils.Cal(n1, n2 , operator)
	fmt.Printf("result~=%.2f", result)
	// //..需求


}