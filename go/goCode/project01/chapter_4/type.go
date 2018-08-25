package main

import "fmt"

type TZ int//别名

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Println("c has the value: %d", c)
}