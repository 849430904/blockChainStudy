package main
import (
	"fmt"
)

//切片的遍历方式
func main() {

	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4] // 20, 30, 40
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	fmt.Println()
	//使用for--range 方式遍历切片
	for i, v := range slice {//如果遍历时不想要i,可以用_替换旧
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	//切片可以继续切片,slice2就是从slice中切出来的
	slice2 := slice[1:2] //  slice [ 20, 30, 40]    [30]
	slice2[0] = 100  // 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	fmt.Println()

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) //100, 200, 300,400, 500, 600

	//通过append将切片slice3追加给slice3； 把切片再追加到切片
	slice3 = append(slice3, slice3...) // 100, 200, 300,400, 500, 600 100, 200, 300,400, 500, 600
	fmt.Println("slice3", slice3)
	/**
	 append函数底层原理：
	    1，切片append操作本质就是对数组扩容
		2.go底层会创建一个新的数组newArr（容量一般为原来的2倍）
		3，将slice原来包含的元素拷贝到新的数组newArr
		4，slice重新引用到newArr
		5,注意newArr由底层维护，程序员看不到
	**/


	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝，举例说明
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)// 1, 2, 3, 4, 5
	fmt.Println("slice5=", slice5) // 1, 2, 3, 4, 5, 0 , 0 ,0,0,0
	//按照上面的代码来看，slice4 slice5的数据空间是独立的，相互不影响，也不是如果slice4=99，而
	//slice5[0] 的值仍然为1
}

/**

切片使用注意事项与细节：
   1，切片初始化时要指明开始、结束下标
   2，切片初始化时，仍然不能越界，范围在[0~len(xx)]
	  如果引用一个存在的数组，可以有以下三种写法:
	  var slice = arr[0:end] 可以简写 var slice = arr[:end]
	  var slice = arr[start:len(arr)] 可以简写 var slice = arr[start:]
	  var slice = arr[0:len(arr)] 可以简写 var slice = arr[:] 
   3,cap是一个内置函数，用于统计切片容量，即最大可以存放多个元素
   4，切片定义后，还不能使用，因为本身是一个空的，需要让其引用到一个数组 ，或者make一个空间供切片使用
   5，切片可以继续切片

   切片可以动态增涨，使用内置函数append
**/