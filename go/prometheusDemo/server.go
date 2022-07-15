package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	proto "microDemo/proto"
	"time"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// 统计请求数量
	httpRequestCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "countDemo",
			Name:      "count_request_total",
			Help:      "Total number of http_request",
		},
	)
	go func() {
		httpRequestCounter.Inc()
		time.Sleep(time.Millisecond)
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
	)
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), &Greeter{})

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
