package main

var a string

func main() {
   a = "G"
   print(a)//G
   f1()//O
}

func f1() {
   a := "O"//局部修改
   print(a)
   f2()//G
}

func f2() {
   print(a)
}