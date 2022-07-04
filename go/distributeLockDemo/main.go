package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// 全局变量
var counter int

func test01() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}

func test02() {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func test03() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}

func main() {
	test03()
}
