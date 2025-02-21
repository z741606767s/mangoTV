package utils

import (
	"github.com/gofiber/fiber/v2"
	"mangoTV/app/config/constants"
)

type C struct {
	c fiber.Ctx
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ResponseOk 请求成功响应
func ResponseOk(c *fiber.Ctx, errCode constants.ErrCode, data interface{}) error {
	return c.Status(errCode.HTTPCode).JSON(Response{
		Success: true,
		Code:    errCode.Code,
		Msg:     errCode.Desc,
		Data:    data,
	})
}

func ResponseErr(c *fiber.Ctx, errCode constants.ErrCode) error {
	return c.Status(errCode.HTTPCode).JSON(Response{
		Success: false,
		Code:    errCode.Code,
		Msg:     errCode.Desc,
		Data:    nil,
	})
}
