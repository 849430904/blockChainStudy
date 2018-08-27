package main

var a = "G"

func main() {
   n()//G
   m()//O
   n()
}

func n() { print(a) }

func m() {
   a := "O"//注意这里的赋值方式，局部赋值
   print(a)
}


// GOG