package authMiddleware

import (
	"github.com/gofiber/fiber/v2"
)

// JwtAuthMiddleware jwt鉴权
func JwtAuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		// TODO: 鉴权逻辑 ...
		return nil
	}
}
