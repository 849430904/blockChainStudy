package main
import (
	"fmt"
	"sort"
)

/**

1)golang没有一个专门的方法针对map的key进行排序
2)golang中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历，得到的输出 可能不一样. 【案例演示 1】
3) golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
**/
func main() {

	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key 放入到 切片中
	//2. 对切片排序 
	//3. 遍历切片，然后按照key来输出map的值

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys{
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
	
}