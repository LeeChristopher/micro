package main

import (
	"context"
	"fmt"
	"hello/proto"

	micro "github.com/micro/go-micro/v2"
)

//这里客户端的代码是以  RPC 方式 请求服务接口
func main() {
	//客户端同时也可能是一个微服务
	//所以这里需要创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.client.greeter"),
	)

	//初始化服务
	service.Init()

	// 创建 Greeter 客户端
	greeter := proto.NewGreeterService("go.micro.srv.greeter", service.Client())

	// 远程调用 Greeter 服务的 Hello 方法
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "学院君"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
