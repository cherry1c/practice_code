package main

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	proto "microDemo/proto"
)

func main() {
	registry := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"192.168.229.133:8500",
			}
		})
	// create a new service
	service := micro.NewService(
		micro.Registry(registry),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	//cl := hello.NewSayService("go.micro.srv.greeter", service.Client())
	cl := proto.NewGreeterService("go.micro.srv.greeter", service.Client())

	// Make request
	rsp, err := cl.Hello(context.Background(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
