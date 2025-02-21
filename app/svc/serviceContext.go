package svc

import (
	"github.com/go-playground/validator/v10"
	"mangoTV/app/config"
	"mangoTV/app/svc/iface"
	"mangoTV/app/svc/repository"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	Conf *config.Config
	*repository.Provider
	Service   iface.IService
	Validator *validator.Validate
}

// ServiceComm 全局服务上下文
var ServiceComm *ServiceContext

// NewServiceContext 创建服务上下文
func NewServiceContext() *ServiceContext {
	ServiceComm = &ServiceContext{
		Conf:      config.Cfg,
		Provider:  repository.NewProvider(),
		Validator: validator.New(),
	}
	return ServiceComm
}
