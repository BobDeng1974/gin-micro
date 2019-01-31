package main

import (
	"context"
	"fmt"
	"gin-micro/protos/user"
	"gin-micro/user/services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
	"log"
	"os"
	"time"
)

func main() {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://182.61.45.175:2379"}
	})

	service := micro.NewService(
		micro.Name("micro.user.svr"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
		micro.WrapHandler(logWrapper),
	)
	service.Init()
	// 注册所有的Handler
	err := user.RegisterUserServiceHandler(service.Server(), new(services.UserService))
	if err != nil {
		log.Println("handler注册失败："+err.Error())
		os.Exit(0);
	}
	// 启动服务
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v", req.Method())
		err := fn(ctx, req, rsp)
		return err
	}
}

