package kafkaQueue

import (
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
	"mangoTV/app/config/constants"
	"time"
)

// InitKafkaQueue 初始化MQ，创建生产和消费队列(需要新类型的队列在这里加)
func InitKafkaQueue() (mqBusiness, mqNotification *KafkaProducer, mqBusinessConsumer, mqNotificationConsumer *KafkaConsumer, closeFunc func()) {
	var err error
	mqBusiness, err = NewKafkaProducer(
		config.Cfg,
		constants.MqQueueTypeBusiness,
		10,
		1*time.Second)
	if err != nil {
		logrus.Errorf("Failed to init kafka queue: %v", err)
		panic(err.Error())
	}

	mqBusinessConsumer, err = NewKafkaConsumer(
		config.Cfg,
		constants.MqQueueTypeBusiness,
		constants.MqQueueTypeBusiness,
		10,
		1*time.Second)
	if err != nil {
		logrus.Errorf("Failed to init kafka consumer: %v", err)
		panic(err.Error())
	}

	mqNotification, err = NewKafkaProducer(
		config.Cfg,
		constants.MqQueueTypeNotification,
		10,
		1*time.Second)
	if err != nil {
		logrus.Errorf("Failed to init kafka queue: %v", err)
		panic(err.Error())
	}

	mqNotificationConsumer, err = NewKafkaConsumer(
		config.Cfg,
		constants.MqQueueTypeNotification,
		constants.MqQueueTypeNotification,
		10,
		1*time.Second)
	if err != nil {
		logrus.Errorf("Failed to init kafka consumer: %v", err)
		panic(err.Error())
	}

	// 返回关闭函数，在需要的时候调用
	closeFunc = func() {
		mqBusiness.Close()
		mqNotification.Close()
		mqBusinessConsumer.Close()
		mqNotificationConsumer.Close()
	}

	return
}
