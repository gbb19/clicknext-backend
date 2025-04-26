package handlers

import (
	handler "clicknext-backend/internal/delivery/http"
	"clicknext-backend/internal/repository/postgres"
	"clicknext-backend/internal/usecase"

	"gorm.io/gorm"
)

type Handler struct {
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}

func NewHandler(db *gorm.DB) *Handler {
	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	return &Handler{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
