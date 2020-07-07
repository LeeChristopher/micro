package main

import (
	"context"
	greeter "hello/greeter/proto"
	"log"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-micro/v2/service"

	"github.com/micro/go-micro/v2/service/grpc"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Println("获取 Greeter.Request 请求")
	rsp.Msg = "你好，" + req.Name
	return nil
}

func main() {
	servi := grpc.NewService(
		service.Name("go.micro.srv.greeter.grpc"),
		service.Address("localhost:9090"),
		service.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
	)

	servi.Init()

	_ = greeter.RegisterGreeterHandler(servi.Server(), new(Greeter))

	if err := servi.Run(); err != nil {
		log.Fatalln(err)
	}
}
