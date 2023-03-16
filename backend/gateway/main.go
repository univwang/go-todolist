package main

import (
	"api-gateway/service"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("39.101.1.158:2379"),
	)
	// 用户
	userMicroService := micro.NewService(
		micro.Name("user.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := service.NewUserService("rpc.user", userMicroService.Client())

	//task
	taskMicroService := micro.NewService(
		micro.Name("task.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := service.NewTaskService("rpc.task", taskMicroService.Client())

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		//将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Handler(weblib.NewRouter(userService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
