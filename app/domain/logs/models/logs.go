package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

const (
	LevelSuccess = "success"
	LevelError   = "error"
	LevelDebug   = "debug"
	LevelInfo    = "info"
	LevelSystem  = "system"
)

type Logs struct {
	ID                 int64     `gorm:"primary_key;auto_increment;column:id;comment:'ID'" json:"id"`
	LogType            int       `gorm:"column:log_type;not null;comment:'应用端: 1 cms管理系统 2 web应用'" json:"logType"`
	Level              string    `gorm:"column:level;type:varchar(20);not null;comment:'日志等级'" json:"level"`
	Message            string    `gorm:"column:level;type:text;not null;comment:'消息'" json:"message"`
	ApiName            string    `gorm:"column:api_name;type:varchar(100);not null;comment:'接口名称'" json:"apiName"`
	Method             string    `gorm:"column:method;type:varchar(100);not null;comment:'请求方法'" json:"method"`
	ReqParams          string    `gorm:"column:req_params;type:text;not null;comment:'入参数'" json:"reqParams"`
	ResParams          string    `gorm:"column:res_params;type:text;not null;comment:'出参数'" json:"resParams"`
	ActionName         string    `gorm:"column:action_name;type:varchar(160);comment:'操作名称'" json:"actionName"`
	ActionIp           string    `gorm:"column:action_ip;type:varchar(20);not null;comment:'操作IP'" json:"actionIp"`
	ActionMan          string    `gorm:"column:action_man;type:varchar(100);not null;comment:'操作人'" json:"actionMan"`
	LatencyTime        string    `gorm:"column:latency_time;type:varchar(20);not null;comment:'请求耗时'" json:"latencyTime"`
	CreatedAt          time.Time `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP;comment:'记录时间'" json:"-"`
	FormattedCreatedAt string    `gorm:"-" json:"createdAt"`
}

type LogsMg struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LogType     int                `bson:"log_type" json:"logType"`
	Level       string             `bson:"level" json:"level"`
	Message     string             `bson:"message" json:"message"`
	ApiName     string             `bson:"api_name" json:"apiName"`
	Method      string             `bson:"method" json:"method"`
	ReqParams   string             `bson:"req_params" json:"reqParams"`
	ResParams   string             `bson:"res_params" json:"resParams"`
	ActionName  string             `bson:"action_name" json:"actionName"`
	ActionIp    string             `bson:"action_ip" json:"actionIp"`
	ActionMan   string             `bson:"action_man" json:"actionMan"`
	LatencyTime string             `bson:"latency_time" json:"latencyTime"`
	Timestamp   primitive.DateTime `bson:"timestamp" json:"timestamp"`
}

func (l *Logs) Table() map[string]string {
	return map[string]string{"ENGINE": "InnoDB", "COMMENT": "log日志表"}
}

func (l *Logs) TableName() string {
	return "c_logs"
}

// AfterFind 是一个 GORM 钩子，在查询记录后调用
func (l *Logs) AfterFind(tx *gorm.DB) (err error) {
	l.FormattedCreatedAt = l.CreatedAt.Format(time.DateTime)
	return nil
}
