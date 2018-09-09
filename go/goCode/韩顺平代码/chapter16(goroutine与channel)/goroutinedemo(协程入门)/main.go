package main
import (
	"fmt"
	"strconv"
	"time"
)


/*
Go 主线程(有程序员直接称为线程/也可以理解成进程): 一个 Go 线程上，可以起多个协程，你可以
这样理解，协程是轻量级的线程[编译器做优化]。

Go协程的特点
	1) 有独立的栈空间
	2) 共享程序堆空间
	3) 调度由用户控制
	4) 协程是轻量级的线程


*/

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

/*
请编写一个程序，完成如下功能:
	1) 在主线程(可以理解成进程)中，开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world" 2) 在主线程中也每隔一秒输出"hello,golang", 输出 10 次后，退出程序
	3) 要求主线程和 goroutine 同时执行.
	4) 画出主线程和协程执行流程图
*/
//编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	/*
	   协程在这里for循环执行
	   1，如果主线程退出了，协程即使还没有执行完毕，也会退出
	   2，当然协程也可以再主线程没有退出前，就自己结束了，比如完成了自己的任务
	*/
	go test() // 开启了一个协程,在协程里面第1秒输出一下打印

	for i := 1; i <= 10; i++ {//goroutine（主线程）中输出
		fmt.Println(" main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

/*
请编写一个程序，完成如下功能:
	1) 在主线程(可以理解成进程)中，开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world" 
	2) 在主线程中也每隔一秒输出"hello,golang", 输出 10 次后，退出程序
	3) 要求主线程和 goroutine 同时执行.
	4) 画出主线程和协程执行流程图
*/