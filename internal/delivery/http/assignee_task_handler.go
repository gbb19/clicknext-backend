package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssigneeTaskHandler struct {
	assigneeTaskUsecase *usecase.AssigneeTaskUsecase
}

func NewAssigneeTaskHandler(assigneeTaskUsecase *usecase.AssigneeTaskUsecase) *AssigneeTaskHandler {
	return &AssigneeTaskHandler{assigneeTaskUsecase: assigneeTaskUsecase}
}

// AssignTask godoc
// @Summary Assign a task to a user
// @Description Assign a task to a user
// @Tags assignments
// @Accept json
// @Produce json
// @Param assignment body dto.AssigneeTaskCreateRequest true "Assignment"
// @Success 201 {object} dto.AssigneeTaskResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /assignments [post]
func (h *AssigneeTaskHandler) AssignTask(c *fiber.Ctx) error {
	var assignmentRequest dto.AssigneeTaskCreateRequest
	if err := c.BodyParser(&assignmentRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	assignment := assignmentRequest.ToAssigneeTaskDomain()

	if err := h.assigneeTaskUsecase.AssignTask(assignment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not assign task"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.FromAssigneeTaskDomain(assignment))
}

// UnassignTask godoc
// @Summary Unassign a task from a user
// @Description Unassign a task from a user
// @Tags assignments
// @Accept json
// @Produce json
// @Param assignee_id query int true "Assignee ID"
// @Param task_id query int true "Task ID"
// @Success 204 {object} nil
// @Failure 400 {object} dto.ErrorResponse
// @Router /assignments [delete]
func (h *AssigneeTaskHandler) UnassignTask(c *fiber.Ctx) error {
	assigneeID, err := strconv.Atoi(c.Query("assignee_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid assignee id"})
	}

	taskID, err := strconv.Atoi(c.Query("task_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	if err := h.assigneeTaskUsecase.UnassignTask(uint(assigneeID), uint(taskID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not unassign task"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetAssigneesByTaskID godoc
// @Summary Get all assignees for a task
// @Description Retrieve all users assigned to a task
// @Tags assignments
// @Accept json
// @Produce json
// @Param task_id query int true "Task ID"
// @Success 200 {array} dto.AssigneeTaskResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /assignments/task [get]
func (h *AssigneeTaskHandler) GetAssigneesByTaskID(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Query("task_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	assignees, err := h.assigneeTaskUsecase.GetAssigneesByTaskID(uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no assignees found"})
	}

	var responses []dto.AssigneeTaskResponse
	for _, assignee := range assignees {
		responses = append(responses, *dto.FromAssigneeTaskDomain(assignee))
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// GetTasksByAssigneeID godoc
// @Summary Get all tasks assigned to a user
// @Description Retrieve all tasks assigned to a specific user
// @Tags assignments
// @Accept json
// @Produce json
// @Param assignee_id query int true "Assignee ID"
// @Success 200 {array} dto.AssigneeTaskResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /assignments/assignee [get]
func (h *AssigneeTaskHandler) GetTasksByAssigneeID(c *fiber.Ctx) error {
	assigneeID, err := strconv.Atoi(c.Query("assignee_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid assignee id"})
	}

	tasks, err := h.assigneeTaskUsecase.GetTasksByAssigneeID(uint(assigneeID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no tasks found"})
	}

	var responses []dto.AssigneeTaskResponse
	for _, task := range tasks {
		responses = append(responses, *dto.FromAssigneeTaskDomain(task))
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}
