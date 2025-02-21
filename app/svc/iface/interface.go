package iface

// IService 服务接口 定义了服务的基本行为
type IService interface {
	GetEventService() IEventService
	GetNotificationService() INotificationService
	GetLogService() ILogService
	GetUsersService() IUsersService
	GetDramaService() IDramaService
}
