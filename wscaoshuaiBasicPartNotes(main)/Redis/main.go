package main

func main() {
	//Redis数据库是NoSQL数据库，不是传统的关系型数据库
	//远程字典服务器，性能非常好，单机能够达到15w qps，非常适合做缓存，也可以持久化
	//开源的，高性能，分布式的内存数据库，也称为数据结构服务器

	//redis基本使用(指令：redisdoc.com)
	//1.添加key-val [set]
	//2.查看当前redis的所有key [key*]
	//3.获取key对应的值 [get key]
	//4.切换redis的数据库 [select index]
	//5.查看当前数据库的key-val数量 [dbsize]
	//6.清空当前数据库的key-val和清空所有数据库的key-val [flushdb flushall]
	//7.键秒值 [setex(set with expire)]

	//redis五大数据类型(指令：redisdoc.com)
	//String,Hash,List,Set,zset
	//1.1 同时设置一个或者多个key-value对 [hmset key value [key value]]
	//1.2 同时返回多个key值 [hmget key [key]]
	//2.1 redis hash 是一个键值对集合 [var user1 map[string]string]
	//2.2 redis hash 是一个string类型的field和value的映射表，hash很适合存储对象
	//3.1 list本质是一个链表，元素是有序的
	//3.2 list的crud基本操作：lpush,rpush,lrange,lpop,rpop,del
	//4.1 set的底层是hashTable数据结构，可以存放很多字符串元素但是无序的且不能重复
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//通过golang向redis写入和读取数据
	// conn, err := redis.Dial("tcp", "127.0.0.1:6379") //1.链接到redis
	// if err != nil {
	// 	fmt.Println("redis.Dial err=", err)
	// 	return
	// }
	// fmt.Println("conn success", conn)
	// defer conn.Close() //关闭链接

	// _, err = conn.Do("Set", "name", "Tom") //2.写入数据string[key-val]
	// if err != nil {
	// 	fmt.Println("set err=", err)
	// 	return
	// }
	// fmt.Println("执行正确", conn)

	// r, err := redis.String(conn.Do("Get", "name")) //3.读取数据string[key-val]
	// if err != nil {
	// 	fmt.Println("get err=", err)
	// 	return
	// }
	// fmt.Println("执行正确", r) //因为返回的r 是interface{},name对应的类型是string，所以需要转换
	// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// //golang操作redis(hash类型)
	// _, err = conn.Do("HMSet", "user02", "name", "Tom", "age", 22) //2.写入数据string[key-val]
	// if err != nil {
	// 	fmt.Println("HMset err=", err)
	// 	return
	// }
	// fmt.Println("执行正确", conn)

	// r, err = redis.String(conn.Do("HMGet", "user02", "name", "age")) //3.读取数据string[key-val]
	// if err != nil {
	// 	fmt.Println("hget err=", err)
	// 	return
	// }
	// for i, v := range r {
	// 	fmt.Println("执行正确", i, v)
	// }
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//golang操作redis
	//1.实现初始化一定数量的链接，放入连接池内
	//2.当Go需要操作redis时，直接到连接池中取出即可
	//3.好处：节省获取redis链接的时间，从而提高效率
	//核心代码：
	// var poor *redis.Pool
	// pool = &redis.Pool{
	// 	Maxidle:     8,//最大空闲链接数
	// 	MaxActive:   0,//表示与数据库的最大连接数，0：没有限制
	// 	idleTimeout: 100,//最大空闲时间
	// 	Dial: func() (redis.Conn, error) {//初始化链接代码
	// 		return redis.Dial("tcp", "localhost:0000")
	// 	},
	// }
	// c := pool.Get()//从连接池中取出一个链接
	// pool.Close()//关闭连接池，就不能从连接池中取出连接

}
