package model

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	Id                 int64     `gorm:"column:id" json:"id" form:"id"`
	MsgType            int64     `gorm:"column:msg_type;is:'1:系统消息', '" json:"msgType" form:"msgType"`
	SendUid            int64     `gorm:"column:send_uid;is:'发送者uid'" json:"sendUid" form:"sendUid"`
	ReceiveUid         int64     `gorm:"column:receive_uid;is:'接收者uid'" json:"receiveUid" form:"receiveUid"`
	Subject            string    `gorm:"column:subject;is:'消息主题'" json:"subject" form:"subject"`
	Content            string    `gorm:"column:content;is:'消息正文'" json:"content" form:"content"`
	ReadAt             int64     `gorm:"column:read_at;is:'已读时间'" json:"readAt" form:"readAt"`
	Status             int64     `gorm:"column:status;is:'1：已发，2：自动发送，3：等待发送'" json:"status" form:"status"`
	CreatedAt          time.Time `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP" json:"-"`
	FormattedCreatedAt string    `gorm:"-" json:"createdAt"` // 用于存储格式化后的时间
	UpdatedAt          time.Time `gorm:"type:datetime;not null;column:updated_at;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"-"`
	FormattedUpdatedAt string    `gorm:"-" json:"updatedAt"` // 用于存储格式化后的时间
}

func (a *Notification) Table() map[string]string {
	return map[string]string{"ENGINE": "InnoDB"}
}

func (a *Notification) TableName() string {
	return "c_notifications"
}

// AfterFind 是一个 GORM 钩子，在查询记录后调用
func (a *Notification) AfterFind(tx *gorm.DB) (err error) {
	a.FormattedCreatedAt = a.CreatedAt.Format(time.DateTime)
	a.FormattedUpdatedAt = a.UpdatedAt.Format(time.DateTime)
	return nil
}
