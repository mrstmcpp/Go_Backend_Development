package handler

import (
	"go-backend-dev/internal/models"
	"go-backend-dev/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	queries *repository.Queries
}

var validate = validator.New()

func NewUserHandler(q *repository.Queries) *UserHandler {
	return &UserHandler{
		queries: q,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	//parsing incoming request bdoy
	var req models.CreateUserRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Response Body",
		})
	}

	//validation
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//parsing dob from string
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format. Please use YYYY-MM-DD",
		})
	}

	//db save
	result, err := h.queries.CreateUser(c.Context(), repository.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Fialed to create user.",
		})
	}

	//getting id from mysql
	userId, err := result.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch the user id of recently created user.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   userId,
		"name": req.Name,
		"dob":  dob.Format("2006-01-02"),
	})

}
