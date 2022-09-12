package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rpc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/web"
	"go_demo/src/pkg/micro/protos/grpc"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Handle("GET", "test1", func(ctx *gin.Context) {

		// 调用 grpc 服务
		c := rpc.NewClient()
		s := grpc.NewUserService("com.foobar.api.example", c)
		rsp_, err := s.Login(ctx, &grpc.LoginRequest{Username: "333"})
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(200, gin.H{"res": rsp_})
	})
	service := web.NewService(web.Name("com.foobar.api.example"), web.Address("127.0.0.1:9000"), web.Handler(engine))
	service.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test 123"))
	})
	service.Init()
	service.Run()

}
