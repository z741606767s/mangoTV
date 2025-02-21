package models

import (
	"gorm.io/gorm"
	"time"
)

type UsersDevice struct {
	ID                 int64     `gorm:"primary_key;auto_increment;column:id;comment:'ID'" json:"id"`
	UserID             int64     `gorm:"index;not null;column:user_id;comment:'用户ID'" json:"userId"`
	IP                 string    `gorm:"size:60;column:ip;comment:'IP'" json:"ip"`
	UserDevice         string    `gorm:"type:text;column:user_device;comment:'用户设备信息'" json:"userDevice"`
	CreatedAt          time.Time `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP" json:"-"`
	FormattedCreatedAt string    `gorm:"-" json:"createdAt"` // 用于存储格式化后的时间
	UpdatedAt          time.Time `gorm:"type:datetime;not null;column:updated_at;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"-"`
	FormattedUpdatedAt string    `gorm:"-" json:"updatedAt"` // 用于存储格式化后的时间
}

func (u *UsersDevice) Table() map[string]string {
	return map[string]string{
		"ENGINE": "InnoDB", "COMMENT": "用户设备表",
	}
}

func (u *UsersDevice) TableName() string {
	return "c_user_devices"
}

// AfterFind 是一个 GORM 钩子，在查询记录后调用
func (u *UsersDevice) AfterFind(tx *gorm.DB) (err error) {
	u.FormattedCreatedAt = u.CreatedAt.Format(time.DateTime)
	u.FormattedUpdatedAt = u.UpdatedAt.Format(time.DateTime)
	return nil
}
