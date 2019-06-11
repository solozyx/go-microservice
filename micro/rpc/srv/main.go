package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"go-microservice/micro/rpc/srv/handler"
	"go-microservice/micro/rpc/srv/subscriber"

	example "go-microservice/micro/rpc/srv/proto/example"
)

func main() {
	// New Service
	// 创建RPC服务
	service := micro.NewService(
		// 服务名称
		micro.Name("go.micro.srv.srv"),
		// 服务版本
		micro.Version("latest"),
		// RPC服务不设定端口 由consul完成端口分配
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	// 可以注释掉 不用 不影响
	micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// 可以注释掉 不用 不影响
	micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
