package middleware

import (
	"go-backend-dev/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		logger.Log.Info("Incoming request",
			zap.String("Method", c.Method()),
			zap.String("Path", c.Path()),
			zap.Int("Status", c.Response().StatusCode()),
		)
		return err
	}
}
