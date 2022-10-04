package main

import (
	"fmt"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	consulAddress = "192.168.229.133:8500"
	localIp       = "172.18.70.226"
	localPort     = 8010
)

func consulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "337"
	registration.Name = "service337"
	registration.Port = localPort
	registration.Tags = []string{"testService"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/actuator/health", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("you are visiting health check api"))
}

func main() {
	consulRegister()
	//定义一个http接口
	http.HandleFunc("/actuator/health", Handler)
	// 1、初始化metric
	httpRequestCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "countDemo",
			Name:      "count_request_total",
			Help:      "Total number of http_request",
		},
	)
	// 2、注册metric
	if err := prometheus.Register(httpRequestCounter); err != nil {
		fmt.Println("prometheus register failed errMessage: ", err.Error())
		return
	}

	go func() {
		// 3、数据上报
		for {
			httpRequestCounter.Inc()
			time.Sleep(time.Millisecond)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("0.0.0.0:8010", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
