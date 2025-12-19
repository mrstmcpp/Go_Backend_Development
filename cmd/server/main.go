package main

import (
	"go-backend-dev/config"
	"go-backend-dev/internal/logger"
	"go-backend-dev/internal/middleware"
	"go-backend-dev/internal/repository"
	"go-backend-dev/internal/routes"
	"go-backend-dev/internal/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	//laoding env
	// if err := godotenv.Load("../../.env"); err != nil {
	// 	logger.Log.Fatal("Error loading .env file", zap.Error(err))
	// }

	if os.Getenv("ENV") != "docker" {
		_ = godotenv.Load()
	}

	//logger
	logger.InitLogger()
	defer logger.Log.Sync()
	logger.Log.Info("Starting server...")

	//validator playgraound
	utils.InitValidator()

	//creating fiber app
	app := fiber.New()

	//middlewares
	app.Use(middleware.RequestIdInjector())
	app.Use(middleware.RequestLogger())

	//db connection
	db, err := config.DbConnection()
	if err != nil {
		logger.Log.Fatal("failed to connnect to db", zap.Error(err))
	}
	defer db.Close()

	//routes
	queries := repository.New(db)
	routes.RegisterRoutes(app, queries)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log.Info("Server is running on port " + port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World...")
	})
}
