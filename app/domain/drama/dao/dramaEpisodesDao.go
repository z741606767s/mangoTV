package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
)

type DramaEpisodesDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewDramaEpisodesDao(db *gorm.DB) *DramaEpisodesDao {
	return &DramaEpisodesDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}
