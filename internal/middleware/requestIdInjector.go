package middleware

import (
	"go-backend-dev/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestIdInjector() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//gen new request id
		requestId := uuid.New().String()

		//storing in contexts
		c.Locals("requestId", requestId)

		//adding to response header
		c.Set("requestId", requestId)

		//startiing time
		start := time.Now()
		//going to next middleware
		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("Request Completed",
			zap.String("request_id", requestId),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)
		return err
	}
}
