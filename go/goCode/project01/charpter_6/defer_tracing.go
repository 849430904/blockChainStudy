package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")//进入某个函数
	defer untrace("b")//离开函数
	fmt.Println("in b")
	a()
}
func main() {
	b()
}
