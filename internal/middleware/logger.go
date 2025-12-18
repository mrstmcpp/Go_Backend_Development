package middleware

import (
	"go-backend-dev/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Log.Info("Incoming request",
			zap.String("Method", c.Method()),
			zap.String("Path", c.Path()),
			zap.Int("Status", c.Response().StatusCode()),
			zap.Duration("Duration", duration),
		)
		return err
	}
}
