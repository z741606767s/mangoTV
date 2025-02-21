package factory

import (
	dramaServices "mangoTV/app/domain/drama/services"
	logsServices "mangoTV/app/domain/logs/services"
	notificationServices "mangoTV/app/domain/notification/services"
	usersServices "mangoTV/app/domain/users/services"
	"mangoTV/app/event"
	"mangoTV/app/svc"
	iface2 "mangoTV/app/svc/iface"
)

// ServiceFactory 服务工厂 用于创建服务实例
type ServiceFactory struct {
	eventService        *event.EventMQ
	NotificationService *notificationServices.NotificationService
	LogService          *logsServices.LogService
	UsersService        *usersServices.UsersService
	DramaService        *dramaServices.DramaService
}

// NewServiceFactory 创建服务实例
func NewServiceFactory(p *svc.ServiceContext) *ServiceFactory {
	return &ServiceFactory{
		event.NewEventMQ(p),
		notificationServices.NewNotificationService(p),
		logsServices.NewLogService(p),
		usersServices.NewUsersService(p),
		dramaServices.NewDramaService(p),
	}
}

// GetEventService 获取MQ队列事件服务实例
func (s *ServiceFactory) GetEventService() iface2.IEventService {
	return s.eventService
}

// GetNotificationService 获取通知服务实例
func (s *ServiceFactory) GetNotificationService() iface2.INotificationService {
	return s.NotificationService
}

// GetLogService 获取日志服务实例
func (s *ServiceFactory) GetLogService() iface2.ILogService {
	return s.LogService
}

// GetUsersService 获取用户服务实例
func (s *ServiceFactory) GetUsersService() iface2.IUsersService {
	return s.UsersService
}

// GetDramaService 获取剧服务实例
func (s *ServiceFactory) GetDramaService() iface2.IDramaService {
	return s.DramaService
}
