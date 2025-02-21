package iface

import "github.com/IBM/sarama"

type INotificationService interface {
	HandleNotification(message *sarama.ConsumerMessage) error
	CreateAndSendSystemMessage(message *sarama.ConsumerMessage) error
}
