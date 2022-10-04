package main

import (
	"context"
	"fmt"
	proto "microDemo/proto"
	"net/http"
	"time"

	"github.com/go-micro/plugins/v4/registry/consul"
	promwrapper "github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

const (
	consulAddress = "192.168.229.133:8500"
	localIp       = "172.18.70.226"
	localPort     = 8002
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
	registry := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"192.168.229.133:8500",
			}
		})

	//go func() {
	//	for {
	//		grpc.DialContext(context.TODO(), "127.0.0.1:9091")
	//		time.Sleep(time.Second)
	//	}
	//}()

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Address("0.0.0.0:8001"),
		micro.Registry(registry),
		micro.WrapHandler(promwrapper.NewHandlerWrapper(
			promwrapper.ServiceName("go.micro.srv.greeter"),
			// promwrapper.ServiceVersion(version),
			// promwrapper.ServiceID(id),
		)),
	)
	go func() {
		consulRegister()
		//定义一个http接口
		http.HandleFunc("/actuator/health", Handler)

		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", localIp, localPort), nil)
		if err != nil {
			fmt.Println("error: ", err.Error())
		}
	}()

	service.Init()
	proto.RegisterGreeterHandler(service.Server(), &Greeter{})

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
