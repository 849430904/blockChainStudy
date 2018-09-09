package main
import (
	"fmt"
	"os" 
)

/**

  文件以流的形式操作
  流：数据在数据源和程序之间经历的路径
  输入流：数据从数据源（文件）到程序（内存）的路径
  输出流：数据从程序（内存）到数据源（文件）的路径

  我们操作文件，会经常使用到 os.File 结构体

  打开文件：
  
*/

func main() {
	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file 文件句柄
	file , err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//输出下文件，看看文件是什么, 看出file 就是一个指针 *File
	fmt.Printf("file=%v", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}