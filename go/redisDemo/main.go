package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 连接
func ConnRedis2() {
	rd := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.4:6379", // url
		Password: "",
		DB:       0, // 数据库
	})
	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Println("ping err :", err)
		return
	}
	fmt.Println(result)
}

// 全局变量：连接数据库
var rd *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "192.168.1.4:6379", // url
	Password: "",
	DB:       0, // 数据库
})

// string操作
func SetAndGet() {
	// set操作：第三个参数是过期时间，如果是0表示不会过期。
	err := rd.Set("k1", "v1", 0).Err()
	if err != nil {
		fmt.Println("set err :", err)
		return
	}
	// get操作
	val, err := rd.Get("k1").Result()
	if err != nil {
		fmt.Println("get err :", err)
		return
	}
	fmt.Println("k1 ==", val) // k1 == v1
	// get获取一个不存在的key，err会返回redis.Nil，因此要注意判断err
	val2, err := rd.Get("k2").Result()
	if err == redis.Nil {
		fmt.Println("k2 does not exist") // k2 does not exist
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("k2", val2)
	}
	rd.Close()
}

// set操作
func Set() {
	defer rd.Close()
	// 添加元素
	rd.SAdd("key", "v1", "v2", "v3")
	// 取出全部元素
	members := rd.SMembers("key")
	for _, value := range members.Val() {
		fmt.Printf("%s\t", value) // v2	v1	v3
	}
	fmt.Println()
	// 删除某个元素
	rd.SRem("key", "v2")
	// 判断某个元素是否存在
	fmt.Println(rd.SIsMember("key", "v2").Val()) // false
	// 获取当前set集合的长度
	fmt.Println(rd.SCard("key")) // scard key: 2
}

func main() {
	SetAndGet()
}
