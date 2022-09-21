package Config

import (
	"github.com/XIAHUALOU/fitness-ioc/examples/services"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}
