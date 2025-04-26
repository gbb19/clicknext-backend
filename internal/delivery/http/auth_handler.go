package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/domain"
	"clicknext-backend/internal/usecase"
	"clicknext-backend/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user by providing user details
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserCreateRequest true "User Info"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/register [post]
// Register สร้างผู้ใช้ใหม่
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.UserCreateRequest
	// BodyParser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid input",
			"details": err.Error(),
		})
	}

	// Validate
	if errorsMap, err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	} else if errorsMap != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": errorsMap,
		})
	}

	// Create
	user := req.ToUserDomain()
	if err := h.authUseCase.Register(user); err != nil {
		if validationErr, ok := err.(*domain.ValidationError); ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   validationErr.Message,
				"details": validationErr.Errors,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create user",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.FromUserDomain(user))
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "User Info"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/login [post]
// Login เข้าสู่ระบบ
// Login เข้าสู่ระบบและรับ token
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate
	if errorsMap, err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	} else if errorsMap != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": errorsMap,
		})
	}

	user, token, err := h.authUseCase.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.AuthResponse{
		Token: token,
		User:  *dto.FromUserDomain(user),
	})
}
