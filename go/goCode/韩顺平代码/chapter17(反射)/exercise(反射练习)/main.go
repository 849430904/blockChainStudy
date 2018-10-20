package main
import (
	"fmt"
	"reflect"
)
func main() {
	var str string = "tom"   //ok
	fs := reflect.ValueOf(&str) //ok fs -> string  要特别注意传入的str的地址,否则下面一行会报错
	fs.Elem().SetString("jack") //ok
	fmt.Printf("%v\n", str) // jack
}
