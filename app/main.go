package main

import (
	"github.com/sirupsen/logrus"
	"mangoTV/app/config"
	"mangoTV/app/config/kafkaQueue"
	"mangoTV/app/cron"
	"mangoTV/app/migrations"
	"mangoTV/app/routers"
	"mangoTV/app/svc"
	"mangoTV/app/svc/factory"
	"mangoTV/app/utils"
	"os"
	"os/signal"
)

func main() {
	// 加载配置
	config.InitConfig()

	// 设置工作目录
	/*err := os.Chdir(config.Cfg.App.AppUploadDir)
	if err != nil {
		logrus.Errorf("Failed to set working directory: %+v", err)
	}*/

	// 是否开启数据迁移
	if config.Cfg.App.IsMigrate {
		migrations.MigrateTable()
	}

	// 加载路由
	app := routers.InitRouter()

	// 加载全局服务
	svc.NewServiceContext()
	svc.ServiceComm.Service = factory.NewServiceFactory(svc.ServiceComm)

	// 初始化MQ
	mqBusiness, mqNotification, mqBusinessConsumer, mqNotificationConsumer, closeFunc := kafkaQueue.InitKafkaQueue()
	svc.ServiceComm.Provider.BusinessMq = mqBusiness
	svc.ServiceComm.Provider.NotificationMq = mqNotification
	svc.ServiceComm.Provider.BusinessConsumerMq = mqBusinessConsumer
	svc.ServiceComm.Provider.NotificationConsumerMq = mqNotificationConsumer
	svc.ServiceComm.Provider.CloseFunc = closeFunc
	defer svc.ServiceComm.Provider.CloseFunc()

	// 加载MQ
	svc.ServiceComm.Service.GetEventService().RegisterConsumers()

	// 加载定时任务
	cron.NewJobRouter(svc.ServiceComm).StartJob()

	// 启动服务
	utils.GoSafe(func() {
		if err := app.Listen(":" + config.Cfg.App.HttpPort); err != nil {
			logrus.Fatalf("start http server error Err:[%s]", err)
		}
	})

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	if err := app.Shutdown(); err != nil {
		logrus.Fatalf("shutdown http server error Err:[%s]", err)
	}
	logrus.Info("Server exiting")
}
