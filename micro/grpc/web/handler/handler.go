package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	// example "github.com/micro/examples/template/srv/proto/example"
	example "go-microservice/micro/grpc/srv/proto/example"

	"github.com/micro/go-grpc"
)

// (响应输出参数 w http.ResponseWriter,请求输入参数 r *http.Request)
func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	// 将传入的json参数解码到创建map中
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	// exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	// RPC调用客户端 连接 "go.micro.srv.srv" 服务
	// 返回RPC调用客户端句柄
	// exampleClient := example.NewExampleService("go.micro.srv.srv", client.DefaultClient)

	// 创建micro grpc服务
	server := grpc.NewService()
	// 服务初始化
	server.Init()
	exampleClient := example.NewExampleService("go.micro.srv.srv", server.Client())

	// RPC客户端调用远程RPC服务端Call服务
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	// 将RPC服务端返回数据最终返回给前端用户
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
