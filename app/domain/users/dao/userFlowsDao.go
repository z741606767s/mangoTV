package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
)

type UserFlowsDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewUsersFlowsDao(db *gorm.DB) *UserFlowsDao {
	return &UserFlowsDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}
