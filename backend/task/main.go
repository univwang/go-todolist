package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/conf"
	"task/core"
	"task/service"
)

func main() {
	conf.Init()
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("39.101.1.158:2379"),
	)

	microService := micro.NewService(
		micro.Name("rpc.task"), //微服务名字
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg), //注册中心
	)

	//结构命令行参数
	microService.Init()

	//注册服务
	_ = service.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	_ = microService.Run()

}
