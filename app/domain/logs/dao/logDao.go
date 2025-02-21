package dao

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mangoTV/app/config/base"
	"mangoTV/app/domain/logs/models"
)

type LogDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewLogDao(db *gorm.DB) *LogDao {
	return &LogDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}

func (dao *LogDao) AddLogs(logs models.Logs, args ...interface{}) error {
	err := dao.AutoDb(args...).Model(&models.Logs{}).Create(&logs).Error
	if err != nil {
		logrus.Errorf("add logs err Err:[%+v] Data:[%+v]", err, logs)
		return err
	}
	return nil
}
