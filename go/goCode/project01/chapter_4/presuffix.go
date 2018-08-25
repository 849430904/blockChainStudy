package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))


	fmt.Printf("%t\n", strings.Contains(str,"This"))
	fmt.Printf("%t\n", strings.Index(str,"is"))
	fmt.Printf("%t\n", strings.LastIndex(str,"is"))
	fmt.Printf("%t\n", strings.IndexRune(str,rune('a')))
	fmt.Printf("%t\n", strings.Count(str,"of"))
	fmt.Printf("%s\n", strings.ToUpper(str))
	fmt.Printf("%s\n", strings.ToLower(str))
}