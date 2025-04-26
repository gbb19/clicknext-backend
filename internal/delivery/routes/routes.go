package routes

import (
	"clicknext-backend/internal/delivery/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *handlers.Handler) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/", h.UserHandler.CreateUser)
	user.Get("/:id", h.UserHandler.GetUserByID)
}
