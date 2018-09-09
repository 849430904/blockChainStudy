package main
import (
	"fmt"
	"encoding/json"//https://studygolang.com/static/pkgdoc/pkg/encoding_json.htm
)

//json 序列化是指，将有 key-value 结构的数据类型(比如结构体、map、切片)序列化成 json 字符串的操作。

/**

	//https://studygolang.com/static/pkgdoc/pkg/encoding_json.htm#Marshal
	//json.Marshal 接收任何形式的参数


将结构体序列化：
 	json.Marshal:将结构体转成json字符串；它返回的是一个byte切片
将map进行序列化：
    json.Marshal()，map序列化出来的Key顺序不固定
将切片进行序列化：
	json.Marshal()
对基本数据类型序列化:对基本类型序列化意义不大
	json.Marshal()
也可以对数组进行序列化：	
*/

//定义一个结构体
type Monster struct {
	Name string `json:"monster_name"` //反射机制
	Age int `json:"monster_age"`//指定序列后的Key为monster_age,而不是Age
	Birthday string //....
	Sal float64
	Skill string
}



func testStruct() {
	//演示
	monster := Monster{
		Name :"牛魔王",
		Age : 500 ,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster) //要写上地址

	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
//monster序列化后={"monster_name":"牛魔王","monster_age":500,"Birthday":"2011-11-11","Sal":8000,"Skill":"牛魔拳"}
//注意age的key为monster_age，而不是Age
}

//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}//map的key为字符串，vlaue为任意类型
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	//将monster 序列化
	data, err := json.Marshal(a)//map本来就是引用传递，不需要写地址
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))
/*

a map 序列化后={"address":"洪崖洞","age":30,"name":"红孩儿"}
*/
}

//演示对切片进行序列化, 我们这个切片 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}//map的key为字符串，vlaue为任意类型
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)//切片里面包含多个map,将m1这个map放到切片里面

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥","夏威夷"}//数组
	slice = append(slice, m2)//切片里面包含多个map,将m2这个map放到切片里面

	//将切片进行序列化操作
	data, err := json.Marshal(slice)//这里不需要传地址
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
	//slice 序列化后=[{"address":"北京","age":"7","name":"jack"},{"address":["墨西哥","夏威夷"],"age":"20","name":"tom"}]
}

//对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
	//num1 序列化后=2345.67
}

func main() {
	//演示将结构体, map , 切片进行序列号
	testStruct()
	testMap()
	testSlice()//演示对切片的序列化
	testFloat64()//演示对基本数据类型的序列化
}