package main

import (
	"fmt"
)

/**

1) 结构体可以使用嵌套匿名结构体所有的字段和方法，即:首字母大写或者小写的字段、方法， 都可以使用。【举例说明】
2) 匿名结构体字段访问可以简化，如图
**/

type A struct {
	Name string
	age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string 
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {

	// var b B
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk()
	// b.A.hello()

	// //上面的写法可以简化

	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	var b B
	b.Name = "jack" // ok
	b.A.Name = "scott"
	b.age = 100  //ok
	b.SayOk()  // B SayOk  jack
	b.A.SayOk() //  A SayOk scott
	b.hello() //  A hello ? "jack" 还是 "scott"

	/**
	
	对上面的代码小结
	(1) 当我们直接通过 b 访问字段或方法时，其执行流程如下比如 b.Name
	(2) 编译器会先看 b 对应的类型有没有 Name, 如果有，则直接调用 B 类型的 Name 字段
	(3) 如果没有就去看 B 中嵌入的匿名结构体 A 有没有声明 Name 字段，如果有就调用,如果没有
	继续查找..如果都找不到就报错.

	**/


}