package main

import (
	"flag"
	"gin-micro/api/clients"
	"gin-micro/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"net/http"
	"time"
)

var (
	httpAddr = flag.String("http","0.0.0.0:8080","http/https access address")
	readTimeout = flag.Duration("read-timeout", 10,"read timeout")
	writeTimeout =flag.Duration("write-timeout",20,"write timeout")
	maxHeaderSize = flag.Int("max-header-size",1048576,"max header size")
)

func main() {
	flag.Parse()
	router := gin.Default()
	routers.Register(router)
	server := &http.Server{
		Addr:           *httpAddr,
		ReadTimeout:    *readTimeout * time.Second,
		WriteTimeout:   *writeTimeout * time.Second,
		MaxHeaderBytes: *maxHeaderSize,
		Handler:        router,
	}
	// 初始化服务
	service := web.NewService(
		web.Address(*httpAddr),
		web.Name("gin-micro.api"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
		web.Registry(clients.Registry),
		web.Server(server),
	)
	service.Handle("/",router)
	if err := service.Run();err != nil {
		panic(err)
	}

}
