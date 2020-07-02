package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"

	"github.com/li-zeyuan/my-micro-service/basic"
	"github.com/li-zeyuan/my-micro-service/basic/config"
	"github.com/li-zeyuan/my-micro-service/handler"
	"github.com/li-zeyuan/my-micro-service/model"
	"github.com/micro/go-micro/config/source/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	service := micro.NewService(
		micro.Name("mu.micro.book.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
