package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ColumnHandler struct {
	columnUsecase *usecase.ColumnUsecase
}

func NewColumnHandler(columnUsecase *usecase.ColumnUsecase) *ColumnHandler {
	return &ColumnHandler{columnUsecase: columnUsecase}
}

// CreateColumn godoc
// @Summary Create a new column
// @Description Create a column in a board
// @Tags columns
// @Accept json
// @Produce json
// @Param column body dto.ColumnCreateRequest true "Column"
// @Success 201 {object} dto.ColumnResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /columns [post]
func (h *ColumnHandler) CreateColumn(c *fiber.Ctx) error {
	// Get userID from Locals
	userID := c.Locals("userID").(uint)

	// Parse request body
	var columnRequest dto.ColumnCreateRequest
	if err := c.BodyParser(&columnRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Set the CreatedBy field to the current user ID
	columnRequest.CreatedBy = userID

	// Convert to domain model
	column := columnRequest.ToColumnDomain()

	// Create the column
	if err := h.columnUsecase.CreateColumn(column); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create column"})
	}

	// Return the created column
	return c.Status(fiber.StatusCreated).JSON(dto.FromColumnDomain(column))
}

// GetColumnByID godoc
// @Summary Get column by ID
// @Description Retrieve column details by column ID
// @Tags columns
// @Accept json
// @Produce json
// @Param id path int true "Column ID"
// @Success 200 {object} dto.ColumnResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /columns/{id} [get]
func (h *ColumnHandler) GetColumnByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid column id"})
	}

	column, err := h.columnUsecase.GetColumnByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "column not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromColumnDomain(column))
}

// GetColumnsByBoardID godoc
// @Summary Get all columns for a board
// @Description Retrieve all columns associated with a board
// @Tags columns
// @Accept json
// @Produce json
// @Param board_id query int true "Board ID"
// @Success 200 {array} dto.ColumnResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /columns [get]
func (h *ColumnHandler) GetColumnsByBoardID(c *fiber.Ctx) error {
	boardID, err := strconv.Atoi(c.Query("board_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid board id"})
	}

	columns, err := h.columnUsecase.GetColumnsByBoardID(uint(boardID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no columns found"})
	}

	var columnResponses []dto.ColumnResponse
	for _, column := range columns {
		columnResponses = append(columnResponses, *dto.FromColumnDomain(column))
	}

	return c.Status(fiber.StatusOK).JSON(columnResponses)
}

// UpdateColumn godoc
// @Summary Update a column
// @Description Update a column's details
// @Tags columns
// @Accept json
// @Produce json
// @Param id path int true "Column ID"
// @Param column body dto.ColumnUpdateRequest true "Column"
// @Success 200 {object} dto.ColumnResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /columns/{id} [put]
func (h *ColumnHandler) UpdateColumn(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid column id"})
	}

	// Get the existing column
	column, err := h.columnUsecase.GetColumnByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "column not found"})
	}

	// Parse the request
	var updateRequest struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Update fields
	if updateRequest.Name != "" {
		column.Name = updateRequest.Name
	}
	if updateRequest.Color != "" {
		column.Color = updateRequest.Color
	}

	// Save the updated column
	if err := h.columnUsecase.UpdateColumn(column); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update column"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromColumnDomain(column))
}

// UpdateColumnPosition godoc
// @Summary Update a column's position
// @Description Move a column to a new position
// @Tags columns
// @Accept json
// @Produce json
// @Param id path int true "Column ID"
// @Param position body PositionUpdateRequest true "New Position"
// @Success 200 {object} dto.ColumnResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /columns/{id}/position [patch]
func (h *ColumnHandler) UpdateColumnPosition(c *fiber.Ctx) error {
	columnID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid column id"})
	}

	// Parse the request
	var updateRequest struct {
		Position int  `json:"position"`
		BoardID  uint `json:"board_id"`
	}
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Update the position
	if err := h.columnUsecase.UpdateColumnPosition(uint(columnID), updateRequest.Position, updateRequest.BoardID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update column position"})
	}

	// Get the updated column
	column, err := h.columnUsecase.GetColumnByID(uint(columnID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "column not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromColumnDomain(column))
}

// DeleteColumn godoc
// @Summary Delete a column
// @Description Delete a column by ID
// @Tags columns
// @Accept json
// @Produce json
// @Param id path int true "Column ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /columns/{id} [delete]
func (h *ColumnHandler) DeleteColumn(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid column id"})
	}

	if err := h.columnUsecase.DeleteColumn(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "column not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
