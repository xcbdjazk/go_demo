package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"go_demo/src/pkg/micro/protos/grpc"
	"log"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, res *grpc.LoginRequest, rep *grpc.LoginResponse) error {
	rep.Username = "123" + res.Username
	return nil
}

type Greeter struct {
	c client.Client
}

// 服务间调用
func (g *Greeter) Hello(ctx context.Context, req *grpc.Request, rsp *grpc.Response) error {
	fmt.Println(req.Name)
	s := grpc.NewUserService("com.foobar.api.example", g.c)
	rsp_, err := s.Login(ctx, &grpc.LoginRequest{Username: "333"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("rsp to Login", rsp_)

	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.foobar.api.example"),
		micro.Address("127.0.0.1:54783"),
	)
	service.Init()
	s := service.Server()
	//err := grpc.RegisterUserServiceHandler(s, &UserService{})
	//if err != nil {
	//	log.Panic("error")
	//}

	err := grpc.RegisterGreeterHandler(s, &Greeter{service.Client()})
	if err != nil {
		log.Panic("error")
	}
	err = grpc.RegisterUserServiceHandler(s, new(UserService))
	if err != nil {
		log.Panic("error")
	}
	if err := service.Run(); err != nil {
		log.Panic(" run error")
	}
}
