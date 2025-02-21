package authMiddleware

import (
	"github.com/gofiber/fiber/v2"
)

// CrossDomainMiddleware 跨越
func CrossDomainMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// TODO: 跨域逻辑...
		return nil
	}
}
