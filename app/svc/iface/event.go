package iface

// IEventService MQ队列事件服务接口
type IEventService interface {
	Emit(mqType string, body []byte) error
	RegisterConsumers()
}
