package utils 
import (
	"fmt"
)

var Num1 int = 300
//将计算的功能，放到一个函数中，然后在需要使用，调用即可
//为了让其它包的文件使用Cal函数，需要将C大小类似其它语言的public
func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符号错误...")
	}
	return res
}

// func Cal(){

// }


/**
包的三大作用:
	 区分相同名字的函数、变量等标识符
	 当程序文件很多时,可以很好的管理项目
	 控制函数、变量等访问范围，即作用域
**/