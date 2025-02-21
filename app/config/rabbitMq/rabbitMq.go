package rabbitMq

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"mangoTV/app/config"
	"time"
)

type RabbitMQ struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	exchangeName string
	exchangeType string
	queueName    string
	routeKey     string
	duration     bool
}

var mqConnection *amqp.Connection

func NewRabbitMQ(exchangeName, exchangeType, queueName string, duration bool) (*RabbitMQ, error) {
	var err error
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Cfg.RabbitMQ.User, config.Cfg.RabbitMQ.Password, config.Cfg.RabbitMQ.Host, config.Cfg.RabbitMQ.Ports)
	logrus.Infof("RabbitMQ URL: " + dsn)

	// 加入连接重试机制
	for i := 1; i <= 20; i++ {
		mqConnection, err = amqp.Dial(dsn)
		if err == nil {
			logrus.Infof("RabbitMQ Connected!") // 连接成功
			break
		}
		time.Sleep(3 * time.Second)
		logrus.WithField("error", err).Errorf("Failed to connect to RabbitMQ, retrying in 3 seconds!")
	}

	channel, _ := mqConnection.Channel()
	mq := &RabbitMQ{
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queueName:    queueName,
		duration:     duration,
		conn:         mqConnection,
		routeKey:     fmt.Sprintf("%s-%s", exchangeName, queueName),
		channel:      channel,
	}

	err = mq.declareAndBindQueue()
	if err != nil {
		logrus.Debugf("Bind queue - %s to exchange - %s faild \n", mq.queueName, mq.exchangeName)
		return nil, err
	}
	return mq, err
}

// 声明交换机
func (mq *RabbitMQ) declareExchange() error {
	err := mq.channel.ExchangeDeclare(
		mq.exchangeName, //name
		mq.exchangeType, //exchangeType "direct"
		mq.duration,     //duration
		false,           //auto-delete
		false,           //internal
		false,           //nowait
		nil,
	)
	return err
}

// 声明队列
func (mq *RabbitMQ) declareQueue() error {
	_, err := mq.channel.QueueDeclare(
		mq.queueName, // name
		mq.duration,  // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	return err
}

// 队列与交换机绑定
func (mq *RabbitMQ) bindQueue() error {
	return mq.channel.QueueBind(
		mq.queueName,    //queue name
		mq.routeKey,     //router key
		mq.exchangeName, //exchange
		false,           //nowait
		nil,
	)
}

// 关闭管道(不想要： 未使用)
func (mq *RabbitMQ) forkChannel() *amqp.Channel {
	ch, err := mq.conn.Channel()
	defer func() {
		_ = ch.Close()
	}()
	if err != nil {
		return nil
	}
	return ch
}

// 声明交换机和队列并且绑定队列
func (mq *RabbitMQ) declareAndBindQueue() (err error) {
	err = mq.declareExchange()
	if err != nil {
		return
	}
	err = mq.declareQueue()
	if err != nil {
		return
	}
	err = mq.bindQueue()
	if err != nil {
		return
	}
	return nil
}

// 关闭管道和关闭连接(不想要： 未使用)
func (mq *RabbitMQ) clear() {
	_ = mq.channel.Close()
	_ = mq.conn.Close()
}

// Publish 发布
func (mq *RabbitMQ) Publish(body []byte) error {
	err := mq.channel.Publish(
		mq.exchangeName, // exchange
		mq.routeKey,     // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		logrus.Errorf("MQ Channel Publish ERROR: %s", err)
	}
	return err
}

// Consume 消费
func (mq *RabbitMQ) Consume() (<-chan amqp.Delivery, error) {
	_ = mq.channel.Qos(250, 0, false)
	deliveryChan, err := mq.channel.Consume(
		mq.queueName, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		logrus.Error("MQ Channel Consume ERROR: ", err)
	}
	return deliveryChan, err
}

// Info 队列info信息
func (mq *RabbitMQ) Info() map[string]string {
	duration := "None-Duration"
	if mq.duration {
		duration = "Duration"
	}
	return map[string]string{
		"exchangeName": mq.exchangeName,
		"exchangeType": mq.exchangeType,
		"queueName":    mq.queueName,
		"duration":     duration,
	}
}
