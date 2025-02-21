package logsMiddleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
	"mangoTV/app/domain/logs/models"
	"net/http"
	"strings"
	"time"
)

func LogrusMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now() // **记录开始时间**
		apiName := c.Path()
		method := c.Method()

		// **解析请求体**
		var requestBody string
		if method == http.MethodPost || method == http.MethodPut {
			bodyBytes := c.Body()
			requestBody = string(bodyBytes)
			c.Request().SetBody(bodyBytes) // **确保请求体可复用**
		}

		ip := c.IP()

		// **执行请求**
		err := c.Next()

		// **计算延迟**
		latencyTime := time.Since(startTime)

		// **排除不需要记录的 API**
		if shouldExcludeAPI(apiName) {
			return err
		}

		// **解析响应体**
		responseBody := string(c.Response().Body())
		/*level, message := getLogLevelAndMessage(responseBody)

		var adminId int64
		var uuid string
		logType := 1

		// **异步写入 MQ 日志**
		utils.GoSafe(func() {
			_ = svc.ServiceComm.Service.GetEventService().Emit(constants.MqQueueTypeBusiness, utils.JsonMarshal(event.EventMsg{
				EventName: constants.EventTypeLogs,
				Content: string(utils.JsonMarshal(v1.MqLogsMsg{
					AdminId:      adminId,
					UUID:         uuid,
					LogType:      logType,
					Level:        level,
					Message:      message,
					ApiName:      apiName,
					Method:       method,
					RequestBody:  requestBody,
					ResponseBody: responseBody,
					Ip:           ip,
					LatencyTime:  latencyTime,
					CreatedAt:    time.Now(),
				})),
			}))
		})*/

		// **日志记录**
		if config.Cfg.App.EnableLogging {
			logrus.WithFields(logrus.Fields{
				"headers":      c.GetReqHeaders(),
				"method":       method,
				"query":        c.Request().URI().QueryArgs().String(),
				"path":         apiName,
				"ip":           ip,
				"latency":      latencyTime,
				"requestBody":  requestBody,
				"responseBody": responseBody,
			}).Debugf("%s %s", method, apiName)
		}

		return err
	}
}

// **判断是否需要排除该API路径**
func shouldExcludeAPI(apiName string) bool {
	excludedPaths := []string{
		"/oss/logs",
	}

	for _, path := range excludedPaths {
		if apiName == path || strings.Contains(apiName, path) {
			return true
		}
	}
	return false
}

// **获取日志级别和消息**
func getLogLevelAndMessage(responseBody string) (string, string) {
	var res struct {
		Success bool   `json:"success"`
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
	}

	if err := json.Unmarshal([]byte(responseBody), &res); err != nil {
		logrus.Warnf("Error unmarshalling response body: %v", err)
		return models.LevelError, "Failed to parse response"
	}

	if res.Code == 200 {
		return models.LevelSuccess, ""
	}
	return models.LevelError, res.Msg
}
