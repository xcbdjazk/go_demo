package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"go_demo/src/pkg/micro_/Users/grpc"
	"log"
)

type S struct {
}

func (s *S) Login(ctx context.Context, res *grpc.LoginRequest, rep *grpc.LoginResponse) error {
	rep.Username = "123" + res.Username
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("123"),
	)
	service.Init()
	err := grpc.RegisterUserServiceHandler(service.Server(), &S{})
	if err != nil {
		log.Panic("error")
	}
	if err := service.Run(); err != nil {
		log.Panic(" run error")
	}
}
