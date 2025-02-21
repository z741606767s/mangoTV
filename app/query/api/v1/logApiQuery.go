package v1

import "time"

type MqLogsMsg struct {
	AdminId      int64         `json:"adminId"`
	UUID         string        `json:"uuid"`
	LogType      int           `json:"logType"`
	Level        string        `json:"level"`
	Message      string        `json:"message"`
	ApiName      string        `json:"apiName"`
	Method       string        `json:"method"`
	RequestBody  string        `json:"requestBody"`
	ResponseBody string        `json:"responseBody"`
	Ip           string        `json:"ip"`
	LatencyTime  time.Duration `json:"latencyTime"`
	CreatedAt    time.Time     `json:"createdAt"`
}
