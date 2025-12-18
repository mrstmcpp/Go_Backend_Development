package handler

import (
	"go-backend-dev/internal/models"
	"go-backend-dev/internal/repository"
	"go-backend-dev/internal/service"
	"go-backend-dev/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	queries *repository.Queries
}

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
	if err := utils.Validate.Struct(req); err != nil {
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

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	//parsing id
	// userId := c.Params("id")
	userId, err := utils.GetIdParam(c)
	if err != nil {
		return err
	}

	//fetching user from db
	user, err := h.queries.GetUserById(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No user found with provided id.",
		})
	}

	//age calculation
	age := service.CalculateAge(user.Dob)

	//res
	response := models.GetUserByIdResponseDTO{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}

	return c.Status(fiber.StatusOK).JSON(response)
	// return c.JSON(fiber.Map{
	// 	"id":   user.ID,
	// 	"name": user.Name,
	// 	"dob":  user.Dob.Format("2006-01-02"),
	// 	"age":  age,
	// })
}

func (h *UserHandler) UpdateUserById(c *fiber.Ctx) error {
	userId, err := utils.GetIdParam(c)
	if err != nil {
		return err
	}

	var req models.UpdateUserByIdRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Response Body",
		})
	}

	//validation
	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//parsing
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format. Please use YYYY-MM-DD",
		})
	}
	result, err := h.queries.UpdateUser(c.Context(), repository.UpdateUserParams{
		ID:   userId,
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user.",
		})
	}

	//checking if updated
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Nothing to update or user not found.",
		})
	}

	return c.JSON(fiber.Map{
		"id":   userId,
		"name": req.Name,
		"dob":  dob.Format("2006-01-02"),
	})

}

func (h *UserHandler) DeleteUserById(c *fiber.Ctx) error {
	userId, err := utils.GetIdParam(c)
	if err != nil {
		return err
	}

	result, err := h.queries.DeleteUser(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}

	// check affected rows
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
