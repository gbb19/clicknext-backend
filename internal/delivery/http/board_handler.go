package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type BoardHandler struct {
	boardUsecase *usecase.BoardUsecase
}

func NewBoardHandler(boardUsecase *usecase.BoardUsecase) *BoardHandler {
	return &BoardHandler{boardUsecase: boardUsecase}
}

// CreateBoard godoc
// @Summary Create a new board
// @Description Create a new board
// @Tags boards
// @Accept json
// @Produce json
// @Param board body dto.BoardCreateRequest true "Create board"
// @Success 201 {object} dto.BoardResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /boards [post]
func (h *BoardHandler) CreateBoard(c *fiber.Ctx) error {
	var request dto.BoardCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	userID := c.Locals("userID").(uint)

	request.CreatedBy = userID

	board := request.ToBoardDomain()

	if err := h.boardUsecase.CreateBoard(board); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create board"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.FromBoardDomain(board))
}

// GetBoardByID godoc
// @Summary Get board by ID
// @Description Retrieve board details by board ID
// @Tags boards
// @Accept json
// @Produce json
// @Param id path int true "Board ID"
// @Success 200 {object} dto.BoardResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /boards/{id} [get]
func (h *BoardHandler) GetBoardByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid board id"})
	}

	board, err := h.boardUsecase.GetBoardByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "board not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromBoardDomain(board))
}

// เพิ่ม handler ใหม่สำหรับดึง boards ที่ user สร้าง
// GetMyCreatedBoards godoc
// @Summary Get boards created by current user
// @Description Retrieve all boards created by the current user
// @Tags boards
// @Accept json
// @Produce json
// @Success 200 {array} dto.BoardResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /boards/created [get]
func (h *BoardHandler) GetMyCreatedBoards(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	boards, err := h.boardUsecase.GetBoardsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve boards"})
	}

	var boardResponses []dto.BoardResponse
	for _, board := range boards {
		boardResponses = append(boardResponses, *dto.FromBoardDomain(board))
	}

	return c.Status(fiber.StatusOK).JSON(boardResponses)
}

// เพิ่ม handler ใหม่สำหรับดึง boards ที่ user เข้าร่วม
// GetMyJoinedBoards godoc
// @Summary Get boards joined by current user
// @Description Retrieve all boards that the current user has joined but did not create
// @Tags boards
// @Accept json
// @Produce json
// @Success 200 {array} dto.BoardResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /boards/joined [get]
func (h *BoardHandler) GetMyJoinedBoards(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	boards, err := h.boardUsecase.GetBoardsJoinedByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve boards"})
	}

	var boardResponses []dto.BoardResponse
	for _, board := range boards {
		boardResponses = append(boardResponses, *dto.FromBoardDomain(board))
	}

	return c.Status(fiber.StatusOK).JSON(boardResponses)
}

// เพิ่ม handler ใหม่สำหรับดึง boards ทั้งหมดที่ user มีส่วนร่วม (ทั้งสร้างเองและเข้าร่วม)
// GetAllMyBoards godoc
// @Summary Get all boards associated with current user
// @Description Retrieve all boards that the current user has created or joined
// @Tags boards
// @Accept json
// @Produce json
// @Success 200 {object} AllBoardsResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /boards/my [get]
func (h *BoardHandler) GetAllMyBoards(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// ดึง boards ที่สร้างโดย user
	createdBoards, err := h.boardUsecase.GetBoardsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve created boards"})
	}

	// ดึง boards ที่ user เข้าร่วม
	joinedBoards, err := h.boardUsecase.GetBoardsJoinedByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve joined boards"})
	}

	// แปลงเป็น response
	var createdBoardResponses []dto.BoardResponse
	for _, board := range createdBoards {
		createdBoardResponses = append(createdBoardResponses, *dto.FromBoardDomain(board))
	}

	var joinedBoardResponses []dto.BoardResponse
	for _, board := range joinedBoards {
		joinedBoardResponses = append(joinedBoardResponses, *dto.FromBoardDomain(board))
	}

	// สร้าง response รวม
	response := fiber.Map{
		"created_boards": createdBoardResponses,
		"joined_boards":  joinedBoardResponses,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
