package main

import (
	"go-backend-dev/internal/logger"
	"go-backend-dev/internal/middleware"
	"go-backend-dev/internal/repository"
	"go-backend-dev/internal/routes"
	"go-backend-dev/internal/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	//validator playgraound
	utils.InitValidator()

	//logger
	logger.InitLogger()
	defer logger.Log.Sync()
	logger.Log.Info("Starting server...")

	//creating fiber app
	app := fiber.New()

	//middleware
	app.Use(middleware.RequestLogger())

	//db connection
	db, err := repository.DbConnection()
	if err != nil {
		logger.Log.Fatal("failed to connnect to db", zap.Error(err))
	}
	defer db.Close()

	queries := repository.New(db)

	routes.RegisterRoutes(app, queries)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	logger.Log.Info("Server is running on port 3000")

	app.Listen(":3000")
}
