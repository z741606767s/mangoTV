package dao

import (
	"gorm.io/gorm"
	"mangoTV/app/config/base"
	"mangoTV/app/domain/users/models"
	"time"
)

type UsersDao struct {
	db *gorm.DB
	*base.BaseDao
}

func NewUsersDao(db *gorm.DB) *UsersDao {
	return &UsersDao{
		db:      db,
		BaseDao: base.NewBaseDao(db),
	}
}

func (dao *UsersDao) GetUserList(limit, page int, id int64, uuid string, isVip, status int8, sTime, eTime int64, args ...interface{}) ([]*models.Users, int64, error) {
	var users []*models.Users
	var total int64
	query := dao.AutoDb(args...).Model(&models.Users{})
	if id > 0 {
		query = query.Where("id = ?", id)
	}
	if uuid != "" {
		query = query.Where("uuid = ?", uuid)
	}

	switch isVip {
	case 1:
		query = query.Where("is_vip = ?", models.VipTrue)
	case 2:
		query = query.Where("is_vip = ? OR vip_ex IS NOT NULL", models.VipTrue)
	case 3:
		query = query.Where("is_vip = ? AND vip_ex IS NULL", models.VipFalse)
	default:
	}

	if status > 0 {
		status = status - 1
		query = query.Where("status = ?", status)
	}
	if sTime > 0 && eTime > 0 {
		startTime := time.UnixMilli(sTime).UTC()
		endTime := time.UnixMilli(eTime).UTC()
		query = query.Where("created_at >= ? AND created_at <= ?", startTime, endTime)
	}

	query.Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
