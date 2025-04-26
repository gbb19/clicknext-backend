package handlers

import (
	userhandler "clicknext-backend/internal/delivery/http"
	"clicknext-backend/internal/repository/postgres"
	"clicknext-backend/internal/usecase"

	"gorm.io/gorm"
)

type Handler struct {
	UserHandler *userhandler.UserHandler
}

func NewHandler(db *gorm.DB) *Handler {
	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHandler(userUsecase)

	return &Handler{
		UserHandler: userHandler,
	}
}
