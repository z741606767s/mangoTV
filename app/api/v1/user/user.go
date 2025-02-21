package user

import (
	"github.com/gofiber/fiber/v2"
	"mangoTV/app/config/constants"
	"mangoTV/app/utils"
)

// @title mangoTV short video
// @version 1.0
// @description 海外短视频项目
// @BasePath /api/v1

// GetUserList
// @Summary 获取用户列表
// @Description 客户端获取用户列表
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" // 参数名 参数位置 参数类型 是否必须 参数描述
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserList(c *fiber.Ctx) error {
	//svc.ServiceComm.Service.GetUsersService().GetUserList()
	return utils.ResponseOk(c, constants.ErrCodes.ErrNo, "获取用户列表成功")
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Code    int    `json:"code" example:"200"`   // 错误代码
	Message string `json:"message" example:"ok"` // 错误信息
}

type User struct {
	Id int `json:"id" example:"1"` // user id
}
