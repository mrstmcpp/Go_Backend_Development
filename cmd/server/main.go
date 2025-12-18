package main

import (
	"go-backend-dev/internal/repository"
	"go-backend-dev/internal/routes"
	"go-backend-dev/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//creating fiber server
	app := fiber.New()
	//db connection
	db, err := repository.DbConnection()
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()

	utils.InitValidator()
	queries := repository.New(db)

	routes.RegisterRoutes(app, queries)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
