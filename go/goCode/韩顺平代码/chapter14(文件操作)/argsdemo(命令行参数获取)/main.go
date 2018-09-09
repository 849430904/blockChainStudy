package main
import (
	"fmt"
	"os"
)


/**


  
*/
func main() {

	/*
	var Args []string:Args保管了命令行参数
   */

	fmt.Println("命令行的参数有", len(os.Args))
	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	
}