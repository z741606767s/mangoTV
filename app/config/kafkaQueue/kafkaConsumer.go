package kafkaQueue

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
	"time"
)

type MessageHandler func(message *sarama.ConsumerMessage) error

type KafkaConsumer struct {
	cfg        *config.Config // 配置信息
	Consumer   sarama.ConsumerGroup
	Topic      string        // 主题
	GroupID    string        // 消费者组ID
	maxRetries int           // 最大重试次数
	retryDelay time.Duration // 重试延迟时间
	handlers   map[string]MessageHandler
}

// CMessage 要接收的消息格式
type CMessage struct {
	ID        string `json:"id"`
	EventName string `json:"eventName"`
	Content   string `json:"content"`
}

func NewKafkaConsumer(cfg *config.Config, topic string, groupID string, maxRetries int, retryDelay time.Duration) (*KafkaConsumer, error) {
	newConfig := sarama.NewConfig()
	newConfig.Consumer.Return.Errors = true                  // 如果消费出现错误，返回错误信息
	newConfig.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的消息开始消费
	//newConfig.Version = sarama.V1_0_0_0     // Kafka 版本

	consumer, err := sarama.NewConsumerGroup(cfg.KafkaMQ.Brokers, groupID, newConfig)
	if err != nil {
		logrus.Errorf("Error creating consumer group client: %v", err)
		return nil, err
	}

	return &KafkaConsumer{Consumer: consumer, Topic: topic, GroupID: groupID, maxRetries: maxRetries, retryDelay: retryDelay, handlers: make(map[string]MessageHandler)}, nil
}

func (c *KafkaConsumer) AddHandler(eventName string, handler MessageHandler) {
	c.handlers[eventName] = handler
}

// ConsumeMessages 消费消息
func (c *KafkaConsumer) ConsumeMessages(ctx context.Context) error {
	// 订阅主题
	topics := []string{c.Topic}

	// 开始消费消息
	for {
		err := c.Consumer.Consume(ctx, topics, c)
		if err != nil {
			logrus.Errorf("Error from consumer: %v", err)
			return err
		}

		if ctx.Err() != nil {
			logrus.Errorf("Error from consumer: %v", ctx.Err())
			return ctx.Err()
		}
	}
}

// Setup Cleanup ConsumeClaim 实现sarama.ConsumerGroupHandler接口
func (c *KafkaConsumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *KafkaConsumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 启动一个goroutine处理消息
func (c *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var msg CMessage
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			logrus.Errorf("Failed to unmarshal message: %v", err)
			continue
		}
		message.Value = []byte(msg.Content)
		handler, ok := c.handlers[msg.EventName]
		if !ok {
			logrus.Errorf("No handler for event: %s", msg.EventName)
			continue
		}

		err = handler(message)
		if err != nil {
			logrus.Errorf("Error processing message: %v", err)
		} else {
			// 确认消息
			session.MarkMessage(message, "")
		}
	}
	return nil
}

func (c *KafkaConsumer) Close() error {
	return c.Consumer.Close()
}
