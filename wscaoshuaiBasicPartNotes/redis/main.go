package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool //定义一个全局pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示与数据库的最大连接数，0：没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接代码
			return redis.Dial("tcp", "localhost:0000")
		},
	}
}

func main() {
	conn := pool.Get() //先从pool取出一个链接
	defer conn.Close() //及时关闭
	_, err := redis.String(conn.Do("set", "name", "Tom"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	r, err := conn.Do("Get", "name") //取出数据
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

}
