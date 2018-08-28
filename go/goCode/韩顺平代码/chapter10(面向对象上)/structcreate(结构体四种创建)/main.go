package main
import (
	"fmt"
)

/**

创建结构体变量和访问结构体字段

方式 1-直接声明 
   案例演示: var person Person
方式 2-{}
   案例演示: var person Person = Person{}
方式 3-&
   案例: var person *Person = new (Person)
方式 4-{}
   案例: var person *Person = &Person{}

说明:
	1) 第 3 种和第 4 种方式返回的是 结构体指针。
	2) 结构体指针访问字段的标准方式应该是:(*结构体指针).字段名 ，比如 (*person).Name = "tom" 
	3) 但 go 做了一个简化，也支持 结构体指针.字段名, 比如 person.Name = "tom"。更加符合程序员
	使用的习惯，go 编译器底层 对 person.Name 做了转化 (*person).Name。
**/

type Person struct{
	Name string
	Age int
}
func main() {
	//方式1

	//方式2
	p2 := Person{"mary", 20}
	// p2.Name = "tom"
	// p2.Age = 18
	fmt.Println(p2)

	//方式3-&
	//案例: var person *Person = new (Person)

	var p3 *Person= new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith" 也可以这样写 p3.Name = "smith"

	//原因: go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	//会给 p3 加上 取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith" 
	p3.Name = "john" //

	(*p3).Age = 30
	p3.Age = 100
	fmt.Println(*p3)

	//方式4-{}
	//案例: var person *Person = &Person{}

	//下面的语句，也可以直接给字符赋值
	//var person *Person = &Person{"mary", 60} 
	var person *Person = &Person{}

	//因为person 是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以 person.Name = "scott"
	// 原因和上面一样，底层会对 person.Name = "scott" 进行处理， 会加上 (*person)
	(*person).Name = "scott"
	person.Name = "scott~~"

	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)

}