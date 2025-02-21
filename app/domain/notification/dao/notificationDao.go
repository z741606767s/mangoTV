package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
)

type NotificationDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewNotificationDao(db *gorm.DB) *NotificationDao {
	return &NotificationDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}
