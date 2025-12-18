package handler

import (
	"go-backend-dev/internal/logger"
	"go-backend-dev/internal/models"
	"go-backend-dev/internal/repository"
	"go-backend-dev/internal/service"
	"go-backend-dev/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	userId, err := h.queries.CreateUser(c.Context(), repository.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Fialed to create user.",
		})
	}

	//logging action
	logger.Log.Info("New user created",
		zap.Int32("UserID", userId),
		zap.String("Name", req.Name),
	)

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

	//loggin action
	logger.Log.Info("User fetched",
		zap.Int32("UserID", user.ID),
		zap.String("Name", user.Name),
	)

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

	//logging action
	logger.Log.Info("User updated",
		zap.Int32("UserID", userId),
		zap.String("Name", req.Name),
	)

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

	//logging action
	logger.Log.Info("User deleted",
		zap.Int32("UserID", userId),
	)

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListAllUsers(c *fiber.Ctx) error {
	page, limit, offset := utils.GetPagination(c)

	users, err := h.queries.ListAllUsersWithPagination(
		c.Context(),
		repository.ListAllUsersWithPaginationParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users.",
		})
	}

	totalRecordsResult, err := h.queries.CountUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users count.",
		})
	}

	responseUsers := make([]models.GetUserByIdResponseDTO, 0, len(users))

	for _, user := range users {
		responseUsers = append(responseUsers, models.GetUserByIdResponseDTO{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(user.Dob),
		})
	}

	response := models.ListUsersResponseDTO{
		Page:         page,
		Limit:        limit,
		TotalRecords: int(totalRecordsResult),
		Users:        responseUsers,
	}

	//logging action
	logger.Log.Info("Listed all users",
		zap.Int("Page", page),
		zap.Int("Limit", limit),
		zap.Int("TotalRecords", int(totalRecordsResult)),
	)

	return c.Status(fiber.StatusOK).JSON(response)
}
