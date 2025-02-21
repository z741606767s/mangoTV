package services

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"mangoTV/app/domain/logs/models"
	v1 "mangoTV/app/query/api/v1"
	"mangoTV/app/svc"
	"strconv"
)

type LogService struct {
	ctx *svc.ServiceContext
}

func NewLogService(ctx *svc.ServiceContext) *LogService {
	return &LogService{
		ctx: ctx,
	}
}

// CreateLogEventHandler 操作日志队列消费方法
func (l *LogService) CreateLogEventHandler(message *sarama.ConsumerMessage) error {
	var mqMsg v1.MqLogsMsg
	var username string
	err := json.Unmarshal(message.Value, &mqMsg)
	if err != nil {
		logrus.Errorf("MqCreateLog Unmarshal Err:[%+v] msg:[%s]", err, message.Value)
		return err
	}

	err = l.ctx.Provider.LogDao.AddLogs(models.Logs{
		LogType:     mqMsg.LogType,
		Level:       mqMsg.Level,
		Message:     mqMsg.Message,
		ApiName:     mqMsg.ApiName,
		Method:      mqMsg.Method,
		ReqParams:   mqMsg.RequestBody,
		ResParams:   mqMsg.ResponseBody,
		ActionName:  mqMsg.ApiName,
		ActionIp:    mqMsg.Ip,
		ActionMan:   username,
		LatencyTime: strconv.FormatInt(int64(mqMsg.LatencyTime), 10),
		CreatedAt:   mqMsg.CreatedAt,
	})
	if err != nil {
		logrus.Errorf("MqCreateLog err:[%+v]", err)
		return err
	}
	return nil
}
