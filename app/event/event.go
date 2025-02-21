package event

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"mangoTV/app/config/constants"
	"mangoTV/app/config/kafkaQueue"
	"mangoTV/app/svc"
	"mangoTV/app/utils"
)

type EventMsg struct {
	EventName string `json:"eventName"`
	Content   string `json:"content"`
}

type EventMQ struct {
	ctx *svc.ServiceContext
}

func NewEventMQ(ctx *svc.ServiceContext) *EventMQ {
	return &EventMQ{
		ctx: ctx,
	}
}

// Emit 使用生产消息投递队列调用方法
func (e *EventMQ) Emit(mqType string, body []byte) error {
	switch mqType {
	case constants.MqQueueTypeNotification: //消息通知队列
		utils.GoSafe(func() {
			if err := e.ctx.Provider.NotificationMq.SendMessage(body); err != nil {
				logrus.Errorf("MqQueueTypeNotification Err:[%+v]", err)
			}
		})
	case constants.MqQueueTypeBusiness: //业务队列
		utils.GoSafe(func() {
			if err := e.ctx.Provider.BusinessMq.SendMessage(body); err != nil {
				logrus.Errorf("MqQueueTypeBusiness Err:[%+v]", err)
			}
		})
	default:
		return errors.New("MqTypeError")
	}
	return nil
}

// RegisterConsumers 注册消费队列
func (e *EventMQ) RegisterConsumers() {
	// 注册 NotificationConsumerMq 的处理器
	e.registerConsumer(e.ctx.Provider.BusinessConsumerMq, map[string]kafkaQueue.MessageHandler{
		// 在这里添加 BusinessConsumerMq 需要处理的事件
		constants.EventTypeLogs: e.ctx.Service.GetLogService().CreateLogsMgEventHandler,
	})

	// 注册 BusinessConsumerMq 的处理器
	e.registerConsumer(e.ctx.Provider.NotificationConsumerMq, map[string]kafkaQueue.MessageHandler{
		// 在这里添加 BusinessConsumerMq 需要处理的事件
		constants.EventTypeNotification: e.ctx.Service.GetNotificationService().HandleNotification,
	})
}

// registerConsumer 注册消费者并启动消费
func (e *EventMQ) registerConsumer(consumer *kafkaQueue.KafkaConsumer, handlers map[string]kafkaQueue.MessageHandler) {
	utils.GoSafe(func() {
		// 注册处理器
		for eventType, handler := range handlers {
			consumer.AddHandler(eventType, handler)
		}

		// 启动消费者
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := consumer.ConsumeMessages(ctx)
		if err != nil {
			logrus.Errorf("%s Err:[%+v]", consumer.GroupID, err)
		}
	})
}
