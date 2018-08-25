package main
import "fmt"
func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	 max := 5
	if val := 10; val > max {
		// do something
		fmt.Printf("val > max \n")
	}
}