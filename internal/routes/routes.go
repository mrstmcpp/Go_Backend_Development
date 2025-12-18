package routes

import (
	"go-backend-dev/internal/handler"
	"go-backend-dev/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, queries *repository.Queries) {
	//routes here
	userHandler := handler.NewUserHandler(queries)

	app.Post("/users", userHandler.CreateUser)
}
