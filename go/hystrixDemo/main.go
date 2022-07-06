package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
)

func init() {
	//Timeout: 执行 command 的超时时间。
	//MaxConcurrentRequests：command 的最大并发量 。
	//SleepWindow：当熔断器被打开后，SleepWindow 的时间就是控制过多久后去尝试服务是否可用了。
	//RequestVolumeThreshold： 一个统计窗口 10 秒内请求数量。达到这个请求数量后才去判断是否要开启熔断
	//ErrorPercentThreshold：错误百分比，请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断

	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})
}

func test01() {
	output := make(chan bool, 1)
	errors := hystrix.Go("my_command", func() error {
		// talk to other services
		fmt.Println("normal flow")
		output <- true
		return nil
	}, func(err error) error {
		// do this when services are down
		fmt.Println("failure flow")

		return nil
	})

	select {
	case out := <-output:
		// success
		fmt.Printf("successful %v", out)
	case err := <-errors:
		// failure
		fmt.Printf("failure %v", err)
	}
}
func main() {
	//hystrixStreamHandler := hystrix.NewStreamHandler()
	//hystrixStreamHandler.Start()
	//go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	for i := 0; i < 10000; i++ {
		go test01()
	}
	select {}

}
