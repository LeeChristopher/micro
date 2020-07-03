package main

import (
	"context"
	"fmt"
	proto "hello/proto"

	micro "github.com/micro/go-micro/v2"
)

type GreeterServiceHandler struct{}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 你好, " + req.Name
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// 初始化，会解析命令行参数
	service.Init()

	// 注册处理器，调用 Greeter 服务接口处理请求
	_ = proto.RegisterGreeterHandler(service.Server(), new(GreeterServiceHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
