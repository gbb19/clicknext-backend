package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TagHandler struct {
	tagUsecase *usecase.TagUsecase
}

func NewTagHandler(tagUsecase *usecase.TagUsecase) *TagHandler {
	return &TagHandler{tagUsecase: tagUsecase}
}

// CreateTag godoc
// @Summary Create a new tag
// @Description Create a new tag for tasks
// @Tags tags
// @Accept json
// @Produce json
// @Param tag body dto.TagCreateRequest true "Tag"
// @Success 201 {object} dto.TagResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tags [post]
func (h *TagHandler) CreateTag(c *fiber.Ctx) error {
	// Get userID from Locals
	userID := c.Locals("userID").(uint)

	// Parse request body
	var tagRequest dto.TagCreateRequest
	if err := c.BodyParser(&tagRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Set the CreatedBy field to the current user ID
	tagRequest.CreatedBy = userID

	// Check if tag with same name already exists
	existingTag, err := h.tagUsecase.FindTagByName(tagRequest.Name)
	if err == nil && existingTag != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "tag with this name already exists"})
	}

	// Convert to domain model
	tag := tagRequest.ToTagDomain()

	// Create the tag
	if err := h.tagUsecase.CreateTag(tag); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create tag"})
	}

	// Return the created tag
	return c.Status(fiber.StatusCreated).JSON(dto.FromTagDomain(tag))
}

// GetTagByID godoc
// @Summary Get tag by ID
// @Description Retrieve tag details by tag ID
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} dto.TagResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tags/{id} [get]
func (h *TagHandler) GetTagByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	tag, err := h.tagUsecase.GetTagByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tag not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTagDomain(tag))
}

// GetAllTags godoc
// @Summary Get all tags
// @Description Retrieve all available tags
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {array} dto.TagResponse
// @Router /tags [get]
func (h *TagHandler) GetAllTags(c *fiber.Ctx) error {
	tags, err := h.tagUsecase.GetAllTags()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve tags"})
	}

	var tagResponses []dto.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, *dto.FromTagDomain(tag))
	}

	return c.Status(fiber.StatusOK).JSON(tagResponses)
}

// UpdateTag godoc
// @Summary Update a tag
// @Description Update a tag's name
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Param tag body dto.TagCreateRequest true "Tag"
// @Success 200 {object} dto.TagResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tags/{id} [put]
func (h *TagHandler) UpdateTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	// Get the existing tag
	tag, err := h.tagUsecase.GetTagByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tag not found"})
	}

	// Get userID from Locals
	userID := c.Locals("userID").(uint)

	// Check if user is the creator of the tag
	if tag.CreatedBy != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "you don't have permission to update this tag"})
	}

	// Parse the request
	var updateRequest dto.TagCreateRequest
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Check if tag with new name already exists (if name is changing)
	if updateRequest.Name != tag.Name {
		existingTag, err := h.tagUsecase.FindTagByName(updateRequest.Name)
		if err == nil && existingTag != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "tag with this name already exists"})
		}
	}

	// Update fields
	tag.Name = updateRequest.Name

	// Save the updated tag
	if err := h.tagUsecase.UpdateTag(tag); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update tag"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTagDomain(tag))
}

// DeleteTag godoc
// @Summary Delete a tag
// @Description Delete a tag by ID
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /tags/{id} [delete]
func (h *TagHandler) DeleteTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	// Get the existing tag
	tag, err := h.tagUsecase.GetTagByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tag not found"})
	}

	// Get userID from Locals
	userID := c.Locals("userID").(uint)

	// Check if user is the creator of the tag
	if tag.CreatedBy != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "you don't have permission to delete this tag"})
	}

	if err := h.tagUsecase.DeleteTag(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete tag"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetTagsByUser godoc
// @Summary Get tags created by user
// @Description Retrieve all tags created by the specified user
// @Tags tags
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {array} dto.TagResponse
// @Router /tags/by-user [get]
func (h *TagHandler) GetTagsByUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	tags, err := h.tagUsecase.GetTagsByCreatedBy(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve tags"})
	}

	var tagResponses []dto.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, *dto.FromTagDomain(tag))
	}

	return c.Status(fiber.StatusOK).JSON(tagResponses)
}
