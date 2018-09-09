package main
import (
	"fmt"
	_ "time"
	"sync"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成 

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中  //协程往map里面放数据的时候，需要考虑同步
// 3. map 应该做出一个全局的.

var (
	myMap = make(map[int]int, 10)  
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁， 
	//sync 是包: synchornized 同步
	//Mutex : 是互斥
	lock sync.Mutex
)


//使用协程+管道  ===>???

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func test(n int) {
	
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里我们将 res 放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res //concurrent map writes?
	//解锁
	lock.Unlock()
}

func main() {

	// 我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}


	//休眠10秒钟【第二个问题 】
	//time.Sleep(time.Second * 5)

	//这里我们输出结果,变量这个结果  //读的时候为什么要加锁？
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}

/*
为什么需要 channel:（channel问题的引出）
	1) 前面使用全局变量加锁同步来解决 goroutine 的通讯，但不完美
	2) 主线程在等待所有 goroutine 全部完成的时间很难确定，我们这里设置 10 秒，仅仅是估算。 
	3) 如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有 goroutine 处于工作
	状态，这时也会随主线程的退出而销毁
	4) 通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操作。 
	5) 上面种种分析都在呼唤一个新的通讯机制-channel

*/