package services

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"mangoTV/app/svc"
)

type NotificationService struct {
	ctx *svc.ServiceContext
}

func NewNotificationService(ctx *svc.ServiceContext) *NotificationService {
	return &NotificationService{
		ctx: ctx,
	}
}

// HandleNotification 处理通知发送 -> 处理 MQ
func (n *NotificationService) HandleNotification(message *sarama.ConsumerMessage) error {
	logrus.Debugf("====1111111=====NotificationHandler -> HandleNotification msgBody:[%+v]", message.Value)
	return nil
}

// CreateAndSendSystemMessage 消费者消息处理
func (n *NotificationService) CreateAndSendSystemMessage(message *sarama.ConsumerMessage) error {
	// 处理队列消息
	logrus.Infof("====22222=====CreateAndSendSystemMessage -> CreateAndSendSystemMessage:[%+v]", message.Value)
	return nil
}
