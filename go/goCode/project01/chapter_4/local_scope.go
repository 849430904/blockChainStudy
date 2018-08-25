package main

var a = "G"

func main() {
   n()//G
   m()//O
   n()
}

func n() { print(a) }

func m() {
   a := "O"
   print(a)
}


// GOG