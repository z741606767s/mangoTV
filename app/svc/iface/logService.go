package iface

import "github.com/IBM/sarama"

type ILogService interface {
	CreateLogEventHandler(message *sarama.ConsumerMessage) error
}
