package main

import (
	"github.com/micro/go-micro"
	"log"

	"github.com/li-zeyuan/my-micro-service/handler"
	s "github.com/li-zeyuan/my-micro-service/proto/user"
)

func main() {
	// New Service   新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.service.user"),
		micro.Version("latest"),
	)

	// Initialise service  初始化服务
	service.Init()

	// Register Handler   注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service    启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//func registryOptions(ops *registry.Options) {
//	etcdCfg := config.GetEtcdConfig()
//	ops.Timeout = time.Second * 5
//	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
//}
