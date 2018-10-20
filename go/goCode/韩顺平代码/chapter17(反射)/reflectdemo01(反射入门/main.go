package main
import (
	"reflect"
	"fmt"
)


/*
 反射基本介绍：
	 1，反射可以在运行时动态获取变量的各种信息。比如变量的类型（type）,类别(kind)
	 2,如果是结构体变量，还可以获取到结构体身的信息（包括结构体的字段与方法）
	 3，通过反射，可以修改变量的值，可以调用关联的就读
	 4，使用反射需要: import ("reflect")
	
 
反射应用常景：
	1，不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法
	   反射。例如以下这种桥接模式，比如：
	   func bridge(funcPtr interface{},agrs....interface{})
	  第一个参数funcPtr以接口的形式传函数指针，函数参数args以可变参数的形式传入，bridge函数中可以
	  用反射来动态执行funcPtr函数
	2，对结构体序列化时，如果结构体有指令tag,也会使用到反射生成的对应字符串


反射重要的函数和概念：
	1，reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型
	2，reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型reflect.Value是一个结构体类型。看文档，通过
	   reflect.Value，可以获取到关于该变量的很多信息
    3，变量 、 infterface{}和reflect.Value是可以相互转换的，这点在实际开发中，会经常用到。(通过.Interface()方法转换)
*/

//专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)//通过反射获取变量的类型

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)//通过反射获取变量的值  
	
	n2 := 2 + rVal.Int() //rVal = 100 ,n2 = 102
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)
	
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)


}

//专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==> 
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)
	


	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age int
}

type Monster struct {
	Name string
	Age int
}

func main() {

	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)

	//2. 定义一个Student的实例
	stu := Student{
		Name : "tom",
		Age : 20,
	}
	reflectTest02(stu)


}