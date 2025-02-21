package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
)

type DramaDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewDramaDao(db *gorm.DB) *DramaDao {
	return &DramaDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}
