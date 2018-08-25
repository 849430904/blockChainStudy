package main

var a = "G"

func main() {
   n()//G
   m()//O
   n()//O
}

func n() {
   print(a)
}

func m() {
   a = "O"//全局赋值
   print(a)
}

// GOO