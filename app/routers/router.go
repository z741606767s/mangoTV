package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"mangoTV/app/api/v1/user"
	_ "mangoTV/app/docs"
	"mangoTV/app/middleware/logsMiddleware"
	"mangoTV/app/middleware/recoverMiddleware"
)

func InitRouter() *fiber.App {

	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // 20M
	})

	// 设置静态文件目录
	//app.Use("/public", static.New(fmt.Sprintf("%s/public", config.Cfg.App.AppUploadDir)))

	// 注册中间件
	app.Use(recover.New(recoverMiddleware.NewRecoverConfig()), logsMiddleware.LogrusMiddleware())

	// 注册路由
	userApi := app.Group("/api/user")
	userApi.Get("/getUserList", user.GetUserList)

	// 路由，展示 Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
