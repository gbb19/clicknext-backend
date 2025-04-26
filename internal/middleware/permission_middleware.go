package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CheckOwnership() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loggedInUserID, ok := c.Locals("userID").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		paramID := c.Params("id")
		userID, err := strconv.ParseUint(paramID, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID format",
			})
		}

		if uint(userID) != loggedInUserID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "You can only access your own profile",
			})
		}

		return c.Next()
	}
}
