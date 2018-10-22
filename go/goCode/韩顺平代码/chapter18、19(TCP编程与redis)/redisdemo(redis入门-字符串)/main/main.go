package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)


/*
Redis基本介绍：
	1,redis是NoSql数据库，而不是传统的关系型数据库
	  官网：https:redis.io  和 http://www.redis.cn
	2,Redis：REmote Dlctionart Server（远程字典服务器），Redis性能非常高，
	  单机能够达到15w gps，通常适合做缓存 ，也可以持久化
	3，完全开源免费的，高性能的(key/value)分布式内存数据库，基于内存运行支持
	  持久化的NoSQL数据库，是最热门的NoSQL数据库之一，也称为数据结构服务器

Redis的安装：
Redis命令参考：http://redisdoc.com/

Redis的CRUD操作：
    Redis 的五大数据类型: String(字符串) 、Hash (哈希)、List(列表)、Set(集合) 和 zset(sorted set:有序集合)

String(字符串) -介绍：  
	string 是 redis 最基本的类型，一个 key 对应一个 value。
	string 类型是二进制安全的。除普通的字符串外，也可以存放图片等数据。
	redis 中字符串 value 最大是 512M
	字符串的CRUD:
	set//如果存在就相当于修改，不存在就是添加 ，
		如：存放一个地址信息： address beijing
		   set address beijing  //key=address,value = beijing
	del,如： del address
	get,如: get address
String(字符串)-使用细节和注意事项:
	setex(set with expire)键秒值
	mset[同时设置一个或多个 key-value 对]
	mget[同时获取多个 key-val]


Hash (哈希，类似 golang 里的 Map)-介绍:
	Redis hash 是一个键值对集合。var user1 map[string]string
	Redis hash 是一个 string 类型的 field 和 value 的映射表，hash 特别适合用于存储对 象。
	例如：存放一个User信息:(user1)
		 user1 name "smith" age 30 job "golang coder"  
		 说明:
			 key : user1
			name "smith" age 30 job "golang coder"  就是三对 field-value

	Hash(哈希，类似 golang 里的 Map)-CRUD：hset/hget/hgetall/hdel
Hash-使用细节和注意事项：
	1，在给user设置name和age时，使用hmset和hmget可以一次设置多个 filed 的值和返回多个 field 的值 。	 
	2，hlen 统计一个 hash 有几个元素.
	3，hexists key field
       查看哈希表 key 中，给定域 field 是否存在

List(列表)-介绍:
    列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部(左边)或者尾部(右边)。
	List本质是链表，List的元素是有序的，元素的值可以重复

	例如：存放多个地址信息：
	     city 北京 天津 上海
	说明：
		key:city
		北京 天津 上海 就是三个元素

List(列表)-CRUD：lpush/rpush/lrange/lpop/rpop/del/
List使用细节与注意事项：
	1），index按照索引下标获得元素（从左到右，从编号0开始）
	2），LLEN key
		 返回列表key的长度，如果key不存在，则key被解释为空列表，返回0
	3），List的其它说明：
		List数据，可以从左或者右，插入添加
		如果值全移除，对应的键也就消失了

Set(集合) - 介绍
    Redis的Set是string类型的无序集合。	
    底层是 HashTable 数据结构, Set 也是存放很多字符串元素，字符串元素是无序 的，而且元素的值不能重复
	举例,存放多个邮件列表信息:
		email sgg@sohu.com tom@sohu.com
	说明：
		key:email
		tn@sohu.com tom@sohu.com 就是二个元素

 Set(集合)- CRUD
   举例说明 Redis 的 Set 的 CRUD 操作:
       sadd
	   smembers[取出所有值] 
	   sismember[判断值是否是成员] 
	   srem [删除指定值]

	   
Golang操作redis: ....



常用用法：
    批量 Set/Get 数据
		说明: 通过 Golang 对 Redis 操作，一次操作可以 Set / Get 多个 key-val 数据
		核心代码:
			_, err = c.Do("MSet", "name", "尚硅谷", "address", "北京昌平~")
			r, err := redis.Strings(c.Do("MGet", "name", "address"))
			for _, v := range r { 
				fmt.Println(v)
			}

	给数据设置有效时间：
		说明: 通过 Golang 对 Redis 操作，给 key-value 设置有效时间
		核心代码:
		//给 name 数据设置有效时间为 10s
		_, err = c.Do("expire", "name", 10)

	操作 List：
		说明: 通过 Golang 对 Redis 操作 List 数据类型
		核心代码:
		_, err = c.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
		r, err := redis.String(c.Do("rpop", "heroList"))
	
	*/
func main() {
	//通过go 向redis 写入数据和读取数据
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return 
	}
	defer conn.Close() //关闭..

	//2. 通过go 向redis写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "tomjerry猫猫")
	if err != nil {
		fmt.Println("set  err=", err)
		return 
	}

	//3. 通过go 向redis读取数据 string [key-val]

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set  err=", err)
		return 
	}

	//因为返回 r是 interface{}
	//因为 name 对应的值是string ,因此我们需要转换
	//nameString := r.(string)

	fmt.Println("操作ok ", r)
}