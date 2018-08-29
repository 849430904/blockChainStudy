package main
import (
	"fmt"
)

func main() {

// 	为什么需要切片
//  先看一个需求:我们需要一个数组用于保存学生的成绩，但是学生的个数是不确定的，请问怎么
// 办?解决方案:-》使用切片。


/**
1) 切片的英文是 slice
2) 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
3) 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice)都一样。 
4) 切片的长度是可以变化的，因此切片是一个可以动态变化数组。
5) 切片定义的基本语法:
	var 切片名 []类型 
	比如:var a [] int
**/
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//slice := intArr[1:3]
	//1. slice 就是切片名
	//2. intArr[1:3] 表示 slice 引用到intArr这个数组 
	//3. 引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)    
	slice := intArr[1:3] 
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //  22, 33
	fmt.Println("slice 的元素个数 =", len(slice)) // 2
	fmt.Println("slice 的容量 =", cap(slice)) // 切片的容量是可以动态变化  

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Println()
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //  22, 33

}