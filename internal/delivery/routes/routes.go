package routes

import (
	"clicknext-backend/internal/delivery/http/handlers"
	"clicknext-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *handlers.Handler) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", h.AuthHandler.Register)
	auth.Post("/login", h.AuthHandler.Login)

	user := api.Group("/users")
	user.Use(middleware.Protected())
	user.Get("/:id", middleware.CheckOwnership(), h.UserHandler.GetUserByID)

}
