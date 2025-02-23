package repository

import (
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"mangoTV/app/config"
	"mangoTV/app/config/kafkaQueue"
	dramaDao "mangoTV/app/domain/drama/dao"
	logsDao "mangoTV/app/domain/logs/dao"
	notificationDao "mangoTV/app/domain/notification/dao"
	usersDao "mangoTV/app/domain/users/dao"
)

// Provider DB服务提供器 用于创建DB服务实例
type Provider struct {
	DB                     *gorm.DB
	RedisDB                *redis.Client
	MongoDB                *mongo.Client
	BusinessMq             *kafkaQueue.KafkaProducer
	BusinessConsumerMq     *kafkaQueue.KafkaConsumer
	NotificationMq         *kafkaQueue.KafkaProducer
	NotificationConsumerMq *kafkaQueue.KafkaConsumer
	CloseFunc              func()

	NotificationDao  *notificationDao.NotificationDao
	LogDao           *logsDao.LogDao
	LogsMgDao        *logsDao.LogsMgDao
	UsersDao         *usersDao.UsersDao
	UsersDeviceDao   *usersDao.UsersDeviceDao
	UsersFlowsDao    *usersDao.UserFlowsDao
	DramaDao         *dramaDao.DramaDao
	DramaEpisodesDao *dramaDao.DramaEpisodesDao
}

// NewProvider 创建DB服务实例
func NewProvider() *Provider {
	return &Provider{
		DB:      config.Db,
		RedisDB: config.Client,
		MongoDB: config.MongoClient,

		NotificationDao:  notificationDao.NewNotificationDao(config.Db),
		LogDao:           logsDao.NewLogDao(config.Db),
		LogsMgDao:        logsDao.NewLogsMgDao(config.MongoClient),
		UsersDao:         usersDao.NewUsersDao(config.Db),
		UsersDeviceDao:   usersDao.NewUsersDeviceDao(config.Db),
		UsersFlowsDao:    usersDao.NewUsersFlowsDao(config.Db),
		DramaDao:         dramaDao.NewDramaDao(config.Db),
		DramaEpisodesDao: dramaDao.NewDramaEpisodesDao(config.Db),
	}
}
