package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
)

type UsersDeviceDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewUsersDeviceDao(db *gorm.DB) *UsersDeviceDao {
	return &UsersDeviceDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}
