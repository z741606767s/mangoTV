package constants

import "time"

// Queue Topic队列类型
const (
	MqQueueTypeBusiness     = "topic-business"     // 业务类事件
	MqQueueTypeNotification = "topic-notification" // 通知类事件
)

// 队列事件名称
const (
	EventTypeLogs         = "SystemLogs"   // 系统日志
	EventTypeNotification = "Notification" // 通知
)

const (
	RedisApiExpireTime     = 30 * time.Second
	RedisRand          int = 10
)
