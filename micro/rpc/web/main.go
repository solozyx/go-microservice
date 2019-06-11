package main

import (
	"github.com/micro/go-log"
	"net/http"

	"github.com/micro/go-web"
	"go-microservice/micro/rpc/web/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		// 服务名称
		web.Name("go.micro.web.web"),
		// 服务版本
		web.Version("latest"),
		// 设置web服务端口号
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	// 注册web服务静态资源
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
