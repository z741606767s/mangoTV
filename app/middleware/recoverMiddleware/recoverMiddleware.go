package recoverMiddleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
)

func NewRecoverConfig() recover.Config {
	return recover.Config{
		Next:             nil,
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			buf := debug.Stack()
			_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s", e, buf))

			logrus.Errorf("Fiber-Panic-Trace:[%+v]", map[string]string{
				"uri":         string(c.Request().URI().FullURI()),
				"x-token":     c.Get("x-token", ""),
				"Body":        string(c.Body()),
				"QueryParams": c.Request().URI().QueryArgs().String(),
				"PanicInfo":   fmt.Sprintf("panic: %v\n%s\n", e, buf),
			})
		},
	}
}
