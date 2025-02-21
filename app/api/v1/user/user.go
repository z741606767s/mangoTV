package user

import (
	"github.com/gofiber/fiber/v2"
	"mangoTV/app/config/constants"
	"mangoTV/app/utils"
)

// GetUserList
// @Summary 获取用户列表
// @Description 客户端获取用户列表
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param file formData file true "Upload file"
// @Param id path int true "User ID" // 参数名 参数位置 参数类型 是否必须 参数描述
// @Param name query string false "userName"
// @Param user body User true "User information"
// @Success 200 {object} User "用户信息"
// @Failure 400 {object} ErrorResponse "请求错误"
// @Failure 404 {object} ErrorResponse "未找到用户"
// @Router /users/{id} [get]
func GetUserList(c *fiber.Ctx) error {
	//svc.ServiceComm.Service.GetUsersService().GetUserList()
	return utils.ResponseOk(c, constants.ErrCodes.ErrNo, "获取用户列表成功")
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`      // 错误代码
	Message string `json:"message" example:"error"` // 错误信息
}

type User struct {
	Id int `json:"id" example:"1"` // user id
}
