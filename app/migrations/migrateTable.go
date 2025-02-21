package migrations

import (
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
)

// MigrateTable 用户表迁移
func MigrateTable() {
	err := config.Db.AutoMigrate(
	// 添加要迁移的表模型
	//&models.Logs{},
	)
	if err != nil {
		logrus.Errorf("migrate rbac error Err:[%+v]", err)
		return
	}

	// 设置自增起始值为 100000
	//config.Db.Exec("ALTER TABLE c_users AUTO_INCREMENT = 100000")

	// 写入初始数据
	//SeedData()
	return
}

func SeedData() {

	return
}
