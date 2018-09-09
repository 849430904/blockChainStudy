package main
import (
	"fmt"
	"time"
)


/*

 如果我们起了一个协程，但是这个协程出现了panic，但是我们没有捕获这个panic，就会导致
 整个程序崩溃，这时我们可以在goroutine中使用recover来捕获panic进行处理，这样即使这个
 协程发生了问题，但是主线程仍然不受影响，可以继续执行；
*/
//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
//函数
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {

	go sayHello()
	go test()


	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}