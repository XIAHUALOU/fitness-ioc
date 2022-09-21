package main

import (
	"fmt"
	Injector "github.com/XIAHUALOU/fitness-ioc"
	"github.com/XIAHUALOU/fitness-ioc/examples/Config"
	"github.com/XIAHUALOU/fitness-ioc/examples/services"
)

func main() {
	serviceConfig := Config.NewServiceConfig()

	Injector.BeanFactory.Config(serviceConfig) //展开方法
	//  BeanFactory.Set()
	{
		//这里 测试 userServices
		userService := services.NewUserService()
		Injector.BeanFactory.Apply(userService) //处理依赖
		fmt.Println(userService.Order.Name())
		userService.GetUserInfo(3)

	}
}
