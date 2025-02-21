package kafkaQueue

import (
	"errors"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
	"time"
)

type KafkaProducer struct {
	cfg        *config.Config      // 配置
	producer   sarama.SyncProducer // 同步生产者
	Topic      string              // 主题
	maxRetries int                 // 最大重试次数
	retryDelay time.Duration       // 重试延迟时间
}

func NewKafkaProducer(cfg *config.Config, topic string, maxRetries int, retryDelay time.Duration) (*KafkaProducer, error) {
	newConfig := sarama.NewConfig()
	newConfig.Producer.RequiredAcks = sarama.WaitForAll // 等待所有副本确认
	newConfig.Producer.Retry.Max = maxRetries           // 最大重试次数
	newConfig.Producer.Return.Successes = true          // 返回成功消息
	//config.Version = sarama.V2_8_0_0                 // Kafka 版本

	producer, err := sarama.NewSyncProducer(cfg.KafkaMQ.Brokers, newConfig)
	if err != nil {
		logrus.Errorf("NewKafkaProducer Failed to start Sarama producer: %v", err)
		return nil, err
	}

	return &KafkaProducer{producer: producer, Topic: topic, maxRetries: maxRetries, retryDelay: retryDelay}, nil
}

// SendMessage 发送消息
func (p *KafkaProducer) SendMessage(message []byte) error {
	// 序列化消息为 JSON
	/*msgBytes, err := json.Marshal(message)
	if err != nil {
		logrus.Errorf("topic:%s SendMessage Failed to marshal message: %v", p.Topic, err)
		return err
	}*/

	msgBytes := message

	// 构建 Kafka 消息
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.StringEncoder(msgBytes),
	}

	// 发送消息
	partition, offset, sendErr := p.producer.SendMessage(msg)
	if partition != -1 && msg.Partition != partition {
		logrus.Errorf("topic:%s Unexpected partition", p.Topic)
		return errors.New("unexpected partition")
	}
	if offset != -1 && msg.Offset != offset {
		logrus.Errorf("topic:%s Unexpected offset", p.Topic)
		return errors.New("unexpected offset")
	}
	/*if str, ok := msg.Metadata.(string); !ok || str != "test" {
		logrus.Errorf("topic:%s Unexpected metadata", p.Topic)
		return errors.New("unexpected metadata")
	}*/

	// 发送失败
	if sendErr != nil {
		logrus.Errorf("topic:%s Failed to send message: %v", p.Topic, sendErr)
		return nil
	}

	return nil
}

// SendMessageWithRetry 发送消息并在失败时重试
func (p *KafkaProducer) SendMessageWithRetry(message []byte, maxRetries int, retryDelay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = p.SendMessage(message)
		if err == nil {
			return nil
		}
		logrus.Warnf("Retry %d: Failed to send message: %v", i+1, err)
		time.Sleep(retryDelay)
	}
	return err
}

// Close 关闭生产者
func (p *KafkaProducer) Close() {
	if err := p.producer.Close(); err != nil {
		logrus.Errorf("topic:%s Failed to close Sarama producer: %v", p.Topic, err)
	}
}
