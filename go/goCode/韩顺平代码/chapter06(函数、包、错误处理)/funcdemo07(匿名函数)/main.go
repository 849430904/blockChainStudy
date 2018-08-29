package main
import (
	"fmt"
	"strings"
)


//累加器
func AddUpper() func (int) int {
	var n int = 10 
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(36) // => 36 = '$'   
		fmt.Println("str=", str) // 1. str="hello$" 2. str="hello$$" 3. str="hello$$$"
		return n
	}
}

/*
对上面代码的说明和总结
1) AddUpper 是一个函数，返回的数据类型是 fun (int) int 2) 闭包的说明
2),返回的是一个匿名函数, 但是这个匿名函数引用到函数外的 n ,因此这个匿名函数就和 n 形成一 个整体，构成闭包。
3) 大家可以这样理解: 闭包是类, 函数是操作，n 是字段。函数和它使用到 n 构成闭包。
4) 当我们反复的调用 f 函数时，因为 n 是初始化一次，因此每调用一次就进行累计。
5) 我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引
用到的变量共同构成闭包。
6) 对上面代码的一个修改，加深对闭包的理解
*/


//
// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix)  {
			return name + suffix
		}

		return name
	}
}


/**

上面代码的总结和说明:
1) 返回的匿名函数和 makeSuffix (suffix string) 的 suffix 变量 组合成一个闭包,因为 返回的函数引用 到 suffix 这个变量
2) 我们体会一下闭包的好处，如果使用传统的方法，也可以轻松实现这个功能，
  但是传统方法需要每 次都传入 后缀名，比如 .jpg ,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复 使用。大家可以仔细的体会一把!


  **/

func makeSuffix2(suffix string, name string)  string {


	//如果 name 没有指定后缀，则加上，否则就返回原来的名字
	if !strings.HasSuffix(name, suffix)  {
		return name + suffix
	}

	return name
	
}

func main() {
	
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1))// 11 
	fmt.Println(f(2))// 13
	fmt.Println(f(3))// 16


	//测试makeSuffix 的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg") //如果使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter")) // winter.jgp
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg

	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter")) // winter.jgp
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "bird.jpg")) // bird.jpg



}