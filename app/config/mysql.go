package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func InitMysql() {
	var err error
	var logLevel logger.LogLevel

	mysqlCfg := Cfg.Mysql

	// 是否开启SQL日志
	switch Cfg.Mysql.LogLevel {
	case "Silent":
		logLevel = logger.Silent
	case "Info":
		logLevel = logger.Info
	default:
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel, //logger.Info 打印 //logger.Silent 关闭 SQL 日志
			Colorful:      true,
		},
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Ports, mysqlCfg.DBName)
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "c_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用迁移时的外键约束
		Logger:                                   newLogger,
	})
	if err != nil {
		fmt.Printf("mysql connect err:%v\n", err)
		panic("initMysql error")
	}

	return
}
