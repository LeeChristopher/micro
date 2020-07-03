package main

import (
	"context"
	"encoding/json"
	"hello/proto"
	"log"
	"strings"

	micro "github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/errors"

	api "github.com/micro/go-micro/v2/api/proto"
)

type Say struct {
	Client proto.GreeterService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Say.Hello API 请求")

	// 从请求参数中获取 name 值
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "名字不能为空")
	}

	//将参数交由底层服务处理
	response, err := s.Client.Hello(ctx, &proto.HelloRequest{ //这里调用了Client 也就是实现了服务的接口对象 并且调用了这个对象的方法
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)
	rsp.StatusCode = 200

	return nil
}

func main() {
	//创建一个新服务
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	//服务初始化
	service.Init()

	//将请求转发给底层的 go.micro.srv.greeter 服务处理
	_ = service.Server().Handle(
		service.Server().NewHandler(
			&Say{
				Client: proto.NewGreeterService("go.micro.srv.greeter", service.Client()),
			}),
	)

	//运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
