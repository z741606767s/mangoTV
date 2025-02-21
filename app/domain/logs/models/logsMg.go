package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogsMg struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LogType            int                `bson:"logType" json:"logType"`
	Level              string             `bson:"level" json:"level"`
	Message            string             `bson:"message" json:"message"`
	ApiName            string             `bson:"apiName" json:"apiName"`
	Method             string             `bson:"method" json:"method"`
	ReqParams          string             `bson:"reqParams" json:"reqParams"`
	ResParams          string             `bson:"resParams" json:"resParams"`
	ActionName         string             `bson:"actionName" json:"actionName"`
	ActionIp           string             `bson:"actionIp" json:"actionIp"`
	ActionMan          string             `bson:"actionMan" json:"actionMan"`
	LatencyTime        string             `bson:"latencyTime" json:"latencyTime"`
	CreatedAt          time.Time          `bson:"createdAt" json:"-"`
	FormattedCreatedAt string             `bson:"-" json:"createdAt"`
}
