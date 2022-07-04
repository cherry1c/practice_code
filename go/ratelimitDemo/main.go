package main

import (
	"fmt"
	"github.com/juju/ratelimit"
	"time"
)

var bucketMap map[int64]*ratelimit.Bucket

func main() {
	bucketMap = make(map[int64]*ratelimit.Bucket)
	b := newBucket()
	bucketMap[1] = b
	for i := 0; i < 10000; i++ {
		before := b.Available()
		if b.TakeAvailable(1) != 0 {
			fmt.Printf("获取到了令牌index %d 前后数量-> 前：%d 后：%d\n", i+1, before, b.Available())
		} else {
			fmt.Printf("未获取到令牌，拒绝 index %d\n", i+1)
		}
		time.Sleep(time.Second / 10)
	}
}

func newBucket() *ratelimit.Bucket {
	bucket := ratelimit.NewBucketWithQuantumAndClock(2*time.Second, 10, 10, nil)
	return bucket
}
