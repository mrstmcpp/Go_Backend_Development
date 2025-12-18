package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetIdParam(c *fiber.Ctx) (int32, error) {
	idParam := c.Params("id")

	if idParam == "" {
		return 0, fiber.NewError(fiber.StatusBadRequest, "id parameter is required.")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "invalid id parameter.")
	}
	return int32(id), nil
}
