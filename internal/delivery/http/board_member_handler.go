package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type BoardMemberHandler struct {
	boardMemberUsecase *usecase.BoardMemberUsecase
}

func NewBoardMemberHandler(boardMemberUsecase *usecase.BoardMemberUsecase) *BoardMemberHandler {
	return &BoardMemberHandler{boardMemberUsecase: boardMemberUsecase}
}

// GetBoardMembers godoc
// @Summary Get board members by Board ID
// @Description Retrieve members of a board by Board ID
// @Tags board_members
// @Accept json
// @Produce json
// @Param id path int true "Board ID"
// @Success 200 {array} dto.BoardMemberResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /boards/{id}/members [get]
func (h *BoardMemberHandler) GetBoardMembers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid board id"})
	}

	members, err := h.boardMemberUsecase.GetBoardMemberByBoardID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no members found for this board"})
	}

	var response []*dto.BoardMemberResponse
	for _, member := range members {
		response = append(response, dto.FromBoardMemberDomain(member))
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
