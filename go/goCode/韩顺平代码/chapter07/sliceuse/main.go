package main
import (
	"fmt"
)

func main() {


	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片  方式一
	//slice := intArr[1:3]
	//1. slice 就是切片名
	//2. intArr[1:3] 表示 slice 引用到intArr这个数组 
	//3. 引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)    
	slice1 := intArr[1:3] 
	fmt.Println("intArr=", slice1)
	fmt.Println("slice 的元素是 =", slice1) //  22, 33
	fmt.Println("slice 的元素个数 =", len(slice1)) // 2
	fmt.Println("slice 的容量 =", cap(slice1)) // 切片的容量是可以动态变化  

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v\n", &slice1[0], slice1[0])
	slice1[1] = 34
	fmt.Println()
	fmt.Println()
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice1) //  22, 33


	//演示切片的使用 make  方式二
	//使用make创建切片时，底层也会创建一个数组，只不过这个数组由底层维护
	//，这个数组对程序员不可见
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	//对于切片，必须make使用.
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))


	//方式3
	fmt.Println()
	//第3种方式：定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice)) //3
	fmt.Println("strSlice cap=", cap(strSlice)) // ?

}