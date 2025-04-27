package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskTagHandler struct {
	taskTagUsecase *usecase.TaskTagUsecase
}

func NewTaskTagHandler(taskTagUsecase *usecase.TaskTagUsecase) *TaskTagHandler {
	return &TaskTagHandler{taskTagUsecase: taskTagUsecase}
}

// CreateTaskTag godoc
// @Summary Assign a tag to a task
// @Description Create a new task-tag association
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param taskTag body dto.TaskTagCreateRequest true "Task Tag"
// @Success 201 {object} dto.TaskTagResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tasks/tags [post]
func (h *TaskTagHandler) CreateTaskTag(c *fiber.Ctx) error {
	// Parse request body
	var taskTagRequest dto.TaskTagCreateRequest
	if err := c.BodyParser(&taskTagRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Convert to domain model
	taskTag := taskTagRequest.ToTaskTagDomain()

	// Create the task tag
	if err := h.taskTagUsecase.CreateTaskTag(taskTag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the created task tag
	return c.Status(fiber.StatusCreated).JSON(dto.FromTaskTagDomain(taskTag))
}

// GetTaskTagByID godoc
// @Summary Get task-tag by ID
// @Description Retrieve task-tag details by ID
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param id path int true "Task Tag ID"
// @Success 200 {object} dto.TaskTagResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags/{id} [get]
func (h *TaskTagHandler) GetTaskTagByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task tag id"})
	}

	taskTag, err := h.taskTagUsecase.GetTaskTagByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task tag not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTaskTagDomain(taskTag))
}

// GetTaskTagsByTaskID godoc
// @Summary Get all tags for a task
// @Description Retrieve all tags associated with a task
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param task_id query int true "Task ID"
// @Success 200 {array} dto.TaskTagResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags/by-task [get]
func (h *TaskTagHandler) GetTaskTagsByTaskID(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Query("task_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	taskTags, err := h.taskTagUsecase.GetTaskTagsByTaskID(uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var taskTagResponses []dto.TaskTagResponse
	for _, taskTag := range taskTags {
		taskTagResponses = append(taskTagResponses, *dto.FromTaskTagDomain(taskTag))
	}

	return c.Status(fiber.StatusOK).JSON(taskTagResponses)
}

// GetTaskTagsByTagID godoc
// @Summary Get all tasks for a tag
// @Description Retrieve all tasks associated with a tag
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param tag_id query int true "Tag ID"
// @Success 200 {array} dto.TaskTagResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags/by-tag [get]
func (h *TaskTagHandler) GetTaskTagsByTagID(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Query("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	taskTags, err := h.taskTagUsecase.GetTaskTagsByTagID(uint(tagID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var taskTagResponses []dto.TaskTagResponse
	for _, taskTag := range taskTags {
		taskTagResponses = append(taskTagResponses, *dto.FromTaskTagDomain(taskTag))
	}

	return c.Status(fiber.StatusOK).JSON(taskTagResponses)
}

// DeleteTaskTag godoc
// @Summary Delete a task-tag association
// @Description Delete a task-tag association by ID
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param id path int true "Task Tag ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags/{id} [delete]
func (h *TaskTagHandler) DeleteTaskTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task tag id"})
	}

	if err := h.taskTagUsecase.DeleteTaskTag(uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteTaskTagByTaskAndTag godoc
// @Summary Delete a task-tag association by task and tag
// @Description Delete a task-tag association by task ID and tag ID
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param task_id query int true "Task ID"
// @Param tag_id query int true "Tag ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags [delete]
func (h *TaskTagHandler) DeleteTaskTagByTaskAndTag(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Query("task_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	tagID, err := strconv.Atoi(c.Query("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	if err := h.taskTagUsecase.DeleteTaskTagByTaskAndTag(uint(taskID), uint(tagID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteTaskTagsByTaskID godoc
// @Summary Delete all tags from a task
// @Description Delete all task-tag associations for a specific task
// @Tags tasks,tags
// @Accept json
// @Produce json
// @Param task_id query int true "Task ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/tags/by-task [delete]
func (h *TaskTagHandler) DeleteTaskTagsByTaskID(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Query("task_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	if err := h.taskTagUsecase.DeleteTaskTagsByTaskID(uint(taskID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
